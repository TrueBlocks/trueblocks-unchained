# The Unchained Index

The UnchainedIndex is a smart contract that accomplishes an extremely simple task: it allows anyone to publish an IPFS hash pointing to any data.

That's it. It does nothing else.

While the previous sentence sounds sort of silly, behind it there's an astonishing wallop.

In the same way a pointer to memory in any programming language can point to any data, an IPFS hash can also do so. A pointer in a
programming language can point to something simple, like an integer, or it can point to something very complex such as an entire database.
In either case, no matter what the pointer points to, the pointer itself is the same size. In the case of computer memory, the pointer
is a 64-bit unsigned integer. In the case of IPFS, the pointer is a 32-byte CID.

The UnchainedIndex, by providing a place to store pointers to arbitrary data, may be seen as a pointer into the global memory space
that is IPFS. In other words, into a globally-accessible database. But, unlike a computer's memory space or a database, the memory pointed
to by the UnchainedIndex cannot change. Nor will the fact that this pointer (IPFS hash) has been recorded on an immutable smart contract
ever change. In this sense, the history of the state of the database can be recorded forever for anyone to see.

This desire for immutability is why the UnchainedIndex smart contract is not upgradable. We are purposefully creating an immutable history
to the data that's stored in the location where this pointer points.

We've chosen to make it permissionless, in the sense that anyone may read from it, but more importantly, anyone may publish to it. Each time
someone called `publishHash` we record the sender's address in the database. We call this address the `publisher`. End users may then later
query by `publisher`, there is a Preferred publisher which is the deployer of the contract. In this way, if a user wishes to retrieve data
published by TrueBlocks (the first deployer of the first version of the UnchainedIndex), she may do so by querying the pointer provided by
our address. However, anyone else may also publish, and users may choose to query against that publisher. It's a form of 'oracle by reputation
by choice' on the end-user's part. If the "US Academy of Accounting Professionals" (a made-up organization) one day decides to publish these
hashes, anyone is free to use that organization's data. Not only do we not want to stop them, we can't stop them--by design.

## Pre-requisites

You must have `foundry` installed. [Follow these instructions](https://book.getfoundry.sh/getting-started/installation.html).

## Building

```[bash]
git clone https://github.com/TrueBlocks/trueblocks-unchained
cd trueblocks-unchained
forge build
```

## Testing

```[bash]
cd trueblocks-unchained
forge test
```

## Getting the manifest

```[bash]
chifra state --call "0xcfd7f3b24f3551741f922fd8c4381aa4e00fc8fd|manifestHash()(string)"
```

returns

```[bash]
blockNumber	address	signature	encoding	bytes	compressedResult
14926334	0xcfd7f3b24f3551741f922fd8c4381aa4e00fc8fd	manifestHash()	0x337f3f32		QmXQ6hGXun7QaLwpYtocAm8uG9LkKHnn3sr9GVCjkRc4Dj
```

The last column is the IPFS hash of the manifest, so in one window:

```[bash]
ipfs daemon
```

In another:

```[bash]
ipfs get QmXQ6hGXun7QaLwpYtocAm8uG9LkKHnn3sr9GVCjkRc4Dj
```

or, if you don't want to wait:

```[bash]
curl "https://gateway.pinata.cloud/ipfs/QmXQ6hGXun7QaLwpYtocAm8uG9LkKHnn3sr9GVCjkRc4Dj" -o manifest.tsv
cat manifest.tsv
```