<h1 align="center"> EIP-7503 Chain </h1>

This is the blockchain component of ZK-Hack's EIP-7503 implementation by [Hamid Bateni](https://github.com/irnb) and [Billy Rennekamp](https://github.com/okwme). It's a fork of Berachain's Polaris blockchain. It has a custom pre-compile for interacting with the Cosmos SDK's bank module. It has been modified to add a new funtionality that allows a whitelisted smart contract to mint new ether upon proof of burn. The changes are to the following files:

 - `/cosmos/precompile/bank/bank.go`
 - `/contracts/bindings/cosmos/precompile/bank/i_bank_module.abigen.go`
 - `/contracts/eip-7503/Mint.sol`
 - `/contracts/eip-7503/VerifierI.sol`
 - `/contracts/package.json`
 - `/contracts/truffle-config.js`
 - `/contracts/scripts/deploy.js`
 - `/contracts/scripts/utils.js`

## Build & Test

[Golang 1.20+](https://go.dev/doc/install) and [Foundry](https://book.getfoundry.sh/getting-started/installation) are required for Polaris.

1. Install [go 1.21+ from the official site](https://go.dev/dl/) or the method of your choice. Ensure that your `GOPATH` and `GOBIN` environment variables are properly set up by using the following commands:

   For Ubuntu:

   ```sh
   cd $HOME
   sudo apt-get install golang jq -y
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

   For Mac:

   ```sh
   cd $HOME
   brew install go jq
   export PATH=$PATH:/opt/homebrew/bin/go
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

2. Install Foundry:

   ```sh
   curl -L https://foundry.paradigm.xyz | bash
   ```

3. Clone, Setup and Test:

   ```sh
   cd $HOME
   git clone https://github.com/berachain/polaris
   cd polaris
   git checkout main
   make test-unit
   ```

4. Start a local development network:

   ```sh
   make start
   ```

5. Deploy the Mint contract:

   ```sh
   cd contracts
   yarn deploy
   ```

## 🚧 WARNING: UNDER CONSTRUCTION 🚧

This project is work in progress and subject to frequent changes as we are still working on wiring up the final system.
It has not been audited for security purposes and should not be used in production yet.

The network will have an Ethereum JSON-RPC server running at `http://localhost:8545` and a Tendermint RPC server running at `http://localhost:26657`.
