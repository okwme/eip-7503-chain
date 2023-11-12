pragma circom 2.1.6;


template AccountRLP() {
    signal nonce;
    signal nonceLength;
    signal nonceOffset;
    

    signal balance;
    signal balanceLength;
    signal balanceOffset;

    signal storageRoot;
    signal storageRootLength;
    signal storageRootOffset;

    signal codeHash;
    signal codeHashLength;
    signal codeHashOffset;

    signal rlpLength;
    signal rlpOffset;

    signal output rlpEncodedAccount;

    signal nonceRlp;

    // calculate nonce rlp

    signal nonceRlpPrefix;

    
}