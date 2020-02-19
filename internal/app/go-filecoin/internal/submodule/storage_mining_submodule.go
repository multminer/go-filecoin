package submodule

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-sectorbuilder"
	"github.com/filecoin-project/go-storage-miner"
	"github.com/ipfs/go-datastore"

	"github.com/filecoin-project/go-filecoin/internal/app/go-filecoin/plumbing/msg"
	storageminerconnector "github.com/filecoin-project/go-filecoin/internal/app/go-filecoin/storage_miner_connector"
	"github.com/filecoin-project/go-filecoin/internal/pkg/piecemanager"
	"github.com/filecoin-project/go-filecoin/internal/pkg/poster"
	"github.com/filecoin-project/go-filecoin/internal/pkg/postgenerator"
	appstate "github.com/filecoin-project/go-filecoin/internal/pkg/state"
)

// StorageMiningSubmodule enhances the `Node` with storage mining capabilities.
type StorageMiningSubmodule struct {
	started bool

	// StorageMining is used by the miner to fill and seal sectors.
	PieceManager piecemanager.PieceManager

	// PoStGenerator generates election PoSts
	PoStGenerator postgenerator.PoStGenerator

	minerNode        *storageminerconnector.StorageMinerNodeConnector
	storageMiner     *storage.Miner
	heaviestTipSetCh chan interface{}
	poster           *poster.Poster
}

// NewStorageMiningSubmodule creates a new storage mining submodule.
func NewStorageMiningSubmodule(minerAddr address.Address, ds datastore.Batching, s sectorbuilder.Interface, c *ChainSubmodule, m *MessagingSubmodule, mw *msg.Waiter, w *WalletSubmodule, stateViewer *appstate.Viewer) (*StorageMiningSubmodule, error) {
	minerNode := storageminerconnector.NewStorageMinerNodeConnector(minerAddr, c.ChainReader, c.State, m.Outbox, mw, w.Wallet, stateViewer)
	storageMiner, err := storage.NewMiner(minerNode, ds, s, minerAddr)
	if err != nil {
		return nil, err
	}

	smbe := piecemanager.NewStorageMinerBackEnd(storageMiner, s)
	sbbe := postgenerator.NewSectorBuilderBackEnd(s)

	modu := &StorageMiningSubmodule{
		PieceManager:     smbe,
		PoStGenerator:    sbbe,
		minerNode:        minerNode,
		storageMiner:     storageMiner,
		heaviestTipSetCh: c.HeaviestTipSetCh,
		poster:           poster.NewPoster(minerAddr, m.Outbox, s, minerNode, storageMiner, c.HeaviestTipSetCh, c.State, stateViewer),
	}

	return modu, nil
}

// Start starts the StorageMiningSubmodule
func (s *StorageMiningSubmodule) Start(ctx context.Context) error {
	if s.started {
		return nil
	}

	s.minerNode.StartHeightListener(ctx, s.heaviestTipSetCh)
	err := s.storageMiner.Run(ctx)
	if err != nil {
		return err
	}

	s.poster.StartPoSting(ctx)

	s.started = true

	return nil
}

// Stop stops the StorageMiningSubmodule
func (s *StorageMiningSubmodule) Stop(ctx context.Context) error {
	if !s.started {
		return nil
	}

	s.minerNode.StopHeightListener()
	err := s.storageMiner.Stop(ctx)
	if err != nil {
		return err
	}

	s.poster.StopPoSting()

	s.started = false

	return nil
}
