# Filecoin (go-filecoin)

[![codecov](https://codecov.io/gh/filecoin-project/go-filecoin/branch/master/graph/badge.svg?token=J5QWYWkgHT)](https://codecov.io/gh/filecoin-project/go-filecoin)
[![CircleCI](https://circleci.com/gh/filecoin-project/go-filecoin.svg?style=svg&circle-token=5a9d1cb48788b41d98bdfbc8b15298816ec71fea)](https://circleci.com/gh/filecoin-project/go-filecoin)

> Filecoin implementation in Go

Filecoin is a decentralized storage network that turns cloud storage into an algorithmic market. The
market runs on a blockchain with a native protocol token (also called "filecoin" or FIL), which miners earn
by providing storage to clients.

## Table of Contents

- [Development](#development)
  - [Prerequisites: Installing Go and Rust](#install-go)
  - [Clone](#clone)
  - [Install Dependencies](#install-dependencies)
  - [Managing Submodules](#managing-submodules)
  - [Testing](#testing)
  - [Supported Commands](#supported-commands)
- [Contribute](#contribute)

## Development

### Prerequisites: Installing Go and Rust

The build process for go-filecoin requires at least Go version 1.10. You'll also need Rust to build the `rust-proofs` submodule, which you can download [here][5].

We've found the Go install instructions can lead to some confusing behavior if you're not already familiar with Go's internals. If you're setting up Go for the first time specifically to experiment with Filecoin, we recommend the following process:

[Here we need install instructions for Mac and Linux, using tarballs instead of the package installer. Introduce the concept of $GOPATH by telling the user to set it, even if it's already defaulted to the right place. Windows install instructions needed as well, TBD]

### Clone

```sh
> mkdir -p ${GOPATH}/src/github.com/filecoin-project
> git clone git@github.com:filecoin-project/go-filecoin.git ${GOPATH}/src/github.com/filecoin-project/go-filecoin
```

### Install Dependencies

go-filecoin's dependencies are managed by [gx][2]; this project is not "go gettable." To install gx, gometalinter, and
other build and test dependencies, run:

```sh
> cd ${GOPATH}/src/github.com/filecoin-project/go-filecoin
> go run ./build/*.go deps
```

### Managing Submodules

Filecoin uses Git Submodules to consume `go-proofs`. To initialize the submodule, either run `deps` (as per above), or
initialize the submodule manually:

```sh
> cd ${GOPATH}/src/github.com/filecoin-project/go-filecoin
> git submodule update --init
```

Later, when the head of the `go-proofs` `master` branch changes, you may want to update `go-filecoin` to use these changes:

```sh
> git submodule update --remote
```

Note that updating the `go-proofs` submodule in this way will require a commit to `go-filecoin` (changing the submodule hash).

### Testing

The filecoin binary must be built prior to testing changes made during development. To do so, run:

```sh
> go run ./build/*.go build
```

Then, run the tests:

```sh
> go run ./build/*.go test
```

Note: Build and test can be combined:

```sh
> go run ./build/*.go best
```

### Supported Commands

```sh
# Build
> go run ./build/*.go build

# Install
> go run ./build/*.go install

# Test
> go run ./build/*.go test

# Build & Test
> go run ./build/*.go best

# Coverage
> go run ./build/*.go test -cover

# Lint
> go run ./build/*.go lint

# Race
> go run ./build/*.go test -race

# Deps, Lint, Build, Test (with args passed to Test)
> go run ./build/*.go all
```

Note: Any flag passed to `go run ./build/*.go test` (e.g. `-cover`) will be passed on to `go test`.

## Contribute

See [the contribute file](CONTRIBUTING.md).

If editing the readme, please conform to the [standard-readme][3] specification.

[1]: https://golang.org/dl/
[2]: https://github.com/whyrusleeping/gx
[3]: https://github.com/RichardLitt/standard-readme
[4]: https://golang.org/doc/install
[5]: https://www.rust-lang.org/en-US/install.html
