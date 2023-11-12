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
