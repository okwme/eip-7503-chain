# 📛 EIP7503 (Contract-less Mixers)

## 📝 Description

In this repository, we are trying to implement the EIP7503 concept on the Polaris EVM.

EIP7503: My colleague and I found an innovative way for performing private proof of burn in August. We figured out we can add a new type of transaction to the EVM using this approach for privacy and scalability.


<img src="https://github.com/irnb/eip-7503-chain/assets/41897852/b9ccf391-6c11-4b79-bc60-1dfa33b81ef6" width="700" height="1000">



### 🔒 Privacy Part

Users can deposit their assets to the burn address and then withdraw them through the mint function. The mint function literally creates a new asset by the credit of the privately burned asset. The anonymity set in this approach is somewhat equal to all EVM transactions and addresses.

### 📈 Scalability Part

This is a part of future work. Today, centralized exchanges and other custody providers must handle millions of separate addresses for their user deposits. Afterward, they aggregate these amounts in their cold wallet. This approach consumes a lot of block space and also increases operational costs. 

These products can provide a burn address for their users. For aggregation, they can just send one mint transaction to aggregate assets from multiple addresses into a single transaction.

### 💪 Challenges

1. Modifying the EVM to support new types of transactions and verify our circuit's proof (Done)
2. Circuit verifier implementation
3. Implementing the burn address calculator circuit
4. Implementing the modified Merkle Patricia Trie verification circuit (binary substring finder (Done), MPT path commitment chaining (Done), RLP calculation: redefine the RLP rules in the new numeric system (Done) that can work with Circom and implement it)

## 🧰 Tech Stack

- Circom
- Golang
- Solidity
- TypeScript

## ⛓️ Blockchain

The blockchain component is a fork of [Berachain's Polaris blockchain](https://github.com/berachain/polaris) that uses the Cosmos-SDK with Geth as an EVM module. It has custom pre-compiles for interacting with the Cosmos SDK module, in our case we are targeting the bank module. It has been modified to add a new funtionality that allows a whitelisted smart contract to mint new ether upon proof of burn. The changes are to the following files:

 - `/cosmos/precompile/bank/bank.go`
 - `/contracts/bindings/cosmos/precompile/bank/i_bank_module.abigen.go`
 - `/contracts/eip-7503/Mint.sol`
 - `/contracts/eip-7503/VerifierI.sol`
 - `/contracts/package.json`
 - `/contracts/truffle-config.js`
 - `/contracts/scripts/deploy.js`
 - `/contracts/scripts/utils.js`


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

1. Install Foundry:

   ```sh
   curl -L https://foundry.paradigm.xyz | bash
   ```

2. Clone, Setup and Test:

   ```sh
   cd $HOME
   git clone https://github.com/berachain/polaris
   cd polaris
   git checkout main
   make test-unit
   ```

3. Start a local development network:

   ```sh
   make start
   ```

4. Deploy the Mint contract:

   ```sh
   cd contracts
   yarn deploy
   ```

## 🚧 WARNING: UNDER CONSTRUCTION 🚧

This project is work in progress and subject to frequent changes as we are still working on wiring up the final system.
It has not been audited for security purposes and should not be used in production yet.

The network will have an Ethereum JSON-RPC server running at `http://localhost:8545` and a Tendermint RPC server running at `http://localhost:26657`.
