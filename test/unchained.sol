// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "src/unchained.sol";

contract UnchainIndexTest is Test {
    UnchainedIndex_V2 unchained;

    function setUp() public {
        unchained = new UnchainedIndex_V2();
    }

    function testOwner() public view {
        console.log(unchained.owner());
    }

    function testRead() public {
        assertEq(
            unchained.manifestHashMap(unchained.owner(), "mainnet"),
            "QmP4i6ihnVrj8Tx7cTFw4aY6ungpaPYxDJEZ7Vg1RSNSdm"
        );
        unchained.publishHash("sepolia", "12");
        assertEq(unchained.manifestHashMap(unchained.owner(), "sepolia"), "12");
        assertEq(
            unchained.manifestHashMap(unchained.owner(), "mainnet"),
            "QmP4i6ihnVrj8Tx7cTFw4aY6ungpaPYxDJEZ7Vg1RSNSdm"
        );
    }

    function testWrite() public {
        unchained.publishHash("sepolia", "12");
        assertEq(unchained.manifestHashMap(unchained.owner(), "sepolia"), "12");
        assertEq(
            unchained.manifestHashMap(unchained.owner(), "mainnet"),
            "QmP4i6ihnVrj8Tx7cTFw4aY6ungpaPYxDJEZ7Vg1RSNSdm"
        );
    }
}
