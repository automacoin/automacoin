
### Database

The Storage System is the nucleus of the system. It contains the result of our
exploration of Turing Machines, which, pressumably we would want to use in more
than a single application; As well as the user accounts, with their balances,
and list of solved machines. A subsystem of indexes as well as microservices is
key for an optimum performances

It is conformed by:

* Core Database
* Indexing Database
* Maintainance Microservices
* API Server

#### Core Database

A key/value store with storing capacity in the PB order. This
figure considering that the full space of `(5,2) TMs` is of `6.23 TB`. The core
database is defined in a _Patricia Merkle Trie_, in a way that a set of
transactions made by the **pool manager** ends up with the computation of the
new _storage root_,.

The core database contains only two sets of data:

- Computed Turing Machines

The **key** of the Computed Turing Machines being the _SHA-256_ (size 32 bytes)
of the _JSON representation_ of the transformation.

Example: The following `(5,2) TM` mapping using
[this online application](https://www.freeformatter.com/sha256-generator.html).

```json
{
  "A0":"1BL",
  "A1":"1H",
  "B0":"0CR",
  "B1":"0CR",
  "C0":"0DR",
  "C1":"1ER",
  "D0":"0DL",
  "D1":"1AL",
  "E0":"1AR",
  "E1":"1AL"
}

```

renders to

```
66154f2bab17cc6f0de81f0b121a4c9979ba993aa55fd1b9a372509084148a96
```

While a hash function is comfortable in terms of making the distribution of the
keys of a huge k/v store uniform, as well as having a constant length of key;
There is the problem of irreversibility of the hashes. Alternative indexes of
preimages may be needed to be maintained for keys of interest.

The **value** of the Computed Turing Machine is an _RLP encoded representation_
of the output of the machine, (for example, the string `0111100`)
if the machine halted, or a special flag if not.

Given that the output of the TMs may be a string of `0`s and `1`s, or a special
flag for non-halting machines. A handy conversion algorithm need to be
determined. For example, a byte could be used to represent 4 of such strings:

````
00: represents `0`
01: represents `1`
10: represents a terminator: Ignore everything after this.
11: represents that the machine didn't halt.
````

- User accounts

This is the second component of the Core Database, relevant to the assignment
of computation rewards.

The **key** of a user account is derived from a **user's private key**:

1. A private key is a random 256-bit (32 bytes) number.
2. The ECDSA public key is computed using the `secp256k1` parameters. Its
uncompressed output is 65 bytes.
3. Finally the user key is the SHA256 hash of the latter public key, size
32 bytes.

In this starting iteration, the only **value** such key will hold is the
account balance. Some discussion needs to be done whether we include (in
further implementations) the _user_ that computed the TM in the TM key, or,
include the TM inside the user data, or, both.

The above commentary becomes of relevance as we develop a system to validate
the computed TMs, to be able to audit which user computed a wrong machine, if
found.

#### Indexing Database:

Will store non-critical (yet performance relevant) information:

* Set of TMs (by id) -> user -> TM hashes
* Next TMs to be packaged for computation
* Computed TMs
* Other statistics

#### Maintainance Microservices

In charge of updating indexes, for further consumption.

* Prepare the next TM packages

The criteria proposed in `(Zenil 2012)` to avoid
computing machines where symmetry is observed, or non-halting is evident.

* Computed TMs

To be updated on each transaction performed by *pool managers*.

* Statistics Services

Loops to be developed ad-hoc.

#### API Server

The communication gate with both the *pool manager* as well as the *web site*.
While the API specification is to be detailed below, it is worth mentioning
that all writing messages from the *pool manager* require to carry both
signatures from the *pool manager*, as well as the *client* producing the
computed data.

It is expected in this version to have low throughput.