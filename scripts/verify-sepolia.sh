#!/usr/bin/env fish

forge verify-contract --chain sepolia 0x6c1f2efbd79e6c952228308fe160cdca3238f0a5 src/unchained.sol:UnchainedIndex_V2 (cat .es_key)
