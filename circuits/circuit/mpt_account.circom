pragma circom 2.1.6;

include "./utils/keccak/keccak.circom";
include "./utils/substring_finder.circom";
include "./utils/hasher.circom";
include "./hashbits.circom";

template account_storage (maxBlocks) {
    signal input numLeafLayerBlocks;
    signal input leafLayer[maxBlocks * 136 * 8];

    signal input salt;

    signal input address[136 * 8];

    signal input nonce[256];
    signal input balance[256];
    signal input codeHash[256];
    signal input storageRoot[256];

    signal input numRlpAccountBlocks;
    signal input rlpAccount[136 * 8];

    signal output commitLeafLayer;

    // using the substring finder to find the account data in the rlpAccount

    component checker_1 = substringCheck(maxBlocks, 136 * 8, 32 * 8);

    checker_1.subInput <== nonce;
    checker_1.numBlocks <== numRlpAccountBlocks;
    checker_1.mainInput <== rlpAccount;
    checker_1.out === 1;

    component checker_2 = substringCheck(maxBlocks, 136 * 8, 32 * 8);

    checker_2.subInput <== balance;
    checker_2.numBlocks <== numRlpAccountBlocks;
    checker_2.mainInput <== rlpAccount;
    checker_2.out === 1;

    component checker_3 = substringCheck(maxBlocks, 136 * 8, 32 * 8);

    checker_3.subInput <== codeHash;
    checker_3.numBlocks <== numRlpAccountBlocks;
    checker_3.mainInput <== rlpAccount;
    checker_3.out === 1;

    component checker_4 = substringCheck(maxBlocks, 136 * 8, 32 * 8);

    checker_4.subInput <== storageRoot;
    checker_4.numBlocks <== numRlpAccountBlocks;
    checker_4.mainInput <== rlpAccount;
    checker_4.out === 1;


    // check the rlpAccount in the leafLayer

    component checker_5 = substringCheck(maxBlocks, 136 * 8, 136 * 8);

    checker_5.subInput <== rlpAccount;
    checker_5.numBlocks <== numLeafLayerBlocks;
    checker_5.mainInput <== leafLayer;
    checker_5.out === 1;

    // calculate the trie key

    signal keccakAddress[32 * 8];
    component keccak = Keccak(1);

    keccak.in <== address;
    keccak.blocks <== 1;
    keccak.out ==> keccakAddress;

    signal trieKey[28 * 8];

    for (var i = 0; i < 28 * 8; i++) {
        trieKey[i] <== keccakAddress[i + 4 * 8];
    }

    // check the trie key in the leafLayer

    component checker_6 = substringCheck(maxBlocks, 136 * 8, 28 * 8);

    checker_6.subInput <== trieKey;
    checker_6.numBlocks <== numLeafLayerBlocks;
    checker_6.mainInput <== leafLayer;
    checker_6.out === 1;

    // commit the leaf layer

    component hasherLeaf = HashBits(maxBlocks * 136 * 8, 250);
    hasherLeaf.inp <== leafLayer;

    component hasherCommitToBlocks = Hasher();
    hasherCommitToBlocks.left <== hasherLeaf.out;
    hasherCommitToBlocks.right <== numLeafLayerBlocks;

    component hasherCommitToSalt = Hasher();
    hasherCommitToSalt.left <== hasherCommitToBlocks.hash;   
    hasherCommitToSalt.right <== salt;

    commitLeafLayer <== hasherCommitToSalt.hash;


    // address ownership with using esdsa checking
}

 component main public [ a ] = account_storage(1);