### Why a Patricia Merkle Trie for the storage?

This data structure,
as defined in [Ethereum](https://github.com/ethereum/wiki/wiki/Patricia-Tree),
provides a number of convenient aspects:

* It is ultimately a key value store.
* Provides `O(log(n))` for inserts and lookups.
* Validates into a unique hash called the _root_.
* It is _shardeable_, which is convenient for replication and backups.
* Also, it will make the transition of this system to a blockchain simpler,
migration-wise.