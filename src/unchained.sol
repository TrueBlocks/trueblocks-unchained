// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract UnchainedIndex {
    constructor() {
        owner = msg.sender;
        chainToHash[
            "mainnet"
        ] = "QmP4i6ihnVrj8Tx7cTFw4aY6ungpaPYxDJEZ7Vg1RSNSdm"; // empty file
        emit HashPublished("mainnet", chainToHash["mainnet"]);
        emit OwnerChanged(address(0), owner);
    }

    function publishHash(string memory chain, string memory hash) public {
        require(msg.sender == owner, "msg.sender must be owner");
        chainToHash[chain] = hash;
        emit HashPublished(chain, hash);
    }

    function changeOwner(address newOwner) public returns (address oldOwner) {
        require(msg.sender == owner, "msg.sender must be owner");
        oldOwner = owner;
        owner = newOwner;
        emit OwnerChanged(oldOwner, newOwner);
        return oldOwner;
    }

    function donate() public payable {
        require(owner != address(0), "owner is not set");
        emit DonationSent(owner, msg.value, block.timestamp);
        payable(owner).transfer(address(this).balance);
    }

    function readHash(string memory chain) public view returns (string memory) {
        return chainToHash[chain];
    }

    event HashPublished(string chain, string hash);
    event OwnerChanged(address oldOwner, address newOwner);
    event DonationSent(address from, uint256 amount, uint256 ts);

    mapping(string => string) public chainToHash;
    address public owner;
}
