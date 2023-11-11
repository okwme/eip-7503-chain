// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;
import "./VerifierI.sol";
import "../cosmos/precompile/Bank.sol";
import "../cosmos/precompile/Bank.sol";

contract Mint {
    address public owner;
    address public verifier;
    IBankModule public bank =
        IBankModule(0x4381dC2aB14285160c808659aEe005D51255adD7);

    constructor() {
        owner = msg.sender;
    }

    function setVerifier(address verifier_) public {
        require(msg.sender == owner);
        verifier = verifier_;
    }

    function setOwner(address owner_) public {
        require(msg.sender == owner);
        owner = owner_;
    }

    function mint(
        address payable recipient,
        uint256 amount,
        uint[2] memory a,
        uint[2][2] memory b,
        uint[2] memory c,
        uint[30] memory input
    ) public {
        require(verifyProof(a, b, c, input), "Invalid proof");
        // TODO: better to get denom from staking params
        Cosmos.Coin memory coin = Cosmos.Coin({amount: amount, denom: "abera"});
        Cosmos.Coin[] memory coins = new Cosmos.Coin[](1);
        coins[0] = coin;
        require(bank.send(recipient, coins), "Send failed");
    }

    function verifyProof(
        uint[2] memory a,
        uint[2][2] memory b,
        uint[2] memory c,
        uint[30] memory input
    ) public view returns (bool) {
        return VerifierI(verifier).verifyProof(a, b, c, input);
    }
}
