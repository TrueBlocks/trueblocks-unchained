# The Unchained Index

The unchained index smart contract

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