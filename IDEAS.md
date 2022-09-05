## Ideas for Future Versions

### Version 2 (Required):

- needs to accept chain (string) for any publication
- needs to store the hashes per chain
- our manifest needs to change to reference two additional peices of data. The timestamp file and the names database (in csv format).

### Would be nice

- A value called 'upgradedTo' that accepts an address and stays empty until the owner sets it. 
  - If non-empty, all functions are disabled (through a require) and the value points to the upgraded contract
- Marketing material surrounding the idea that this hash is the cheapest possible way to publish access to the entire chain's index
- Allow others to submit names files for possible inclusion in the names database
- Expansion to accept a 'reason' string which can take on any value
  - index-manifest
  - timestamps
  - names
  - abis

   ### Usage

Anyone may publish. The map records hashes per chain per address. This means anyone can be their own publisher.

Our software will only read our manifest. Other people's software can choose to read our manifest if they wish, but they are free to publish their own manifest. As long as we coordinate on the contents of the manifest and the internal contents of the pointed to resources, it will work.

Analysis code can view the number of times other people publish other manifest files. We don't care. Our software will always read our entries in the database. By allowing anyone to publish, we can collect information on how many times "our" hashes are published. Counters. Voting.

We could also have staking -- without slashing -- but who cares -- we're not serving anyone but ourselves.