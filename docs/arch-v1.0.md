# The Beaver Project
# System Architecture v1.0

## Abstract

The purpose of this system is to explore the Turing Machine space (of different
types, differing in number of states, symbols and, dimensions). This exploration
will help us to know which TMs halt, and hence give us an estimate of busy
beaver numbers. The output of these machines will also help us build
increasingly accurate estimations of the universal distribution.

The scope of this project on its first version is constrained to understand the
challenges of distributing these tasks over a number of clients, enabling them
to be awarded in a currency we are calling _AutomaCoin_.

The present document describes the first version of the _AutomaCoin_ system,
from now, **v1.0**.

In one single paragraph, this is a multi-service architecture, each service
communicating to each other via JSON API. There is a huge size (Terabyte, maybe
Petabyte) **database** of Turing Machines of certain features, user accounts,
and indexes sorting and grouping them. It will be the task of a **pool manager**
to determine which set of combinations of TMs will be processed upon the request
of a **client**. Finished the computations of the **client**, the **pool
manager** will receive the set of TMs, store the results into the **database**
and assign an _AutomaCoin_ reward to the **client**. The latter, as well as any
other user, can verify the space of TMs solved, the awards given and other data
to a **web site** feeding from the **database**.

Fine grained discussions, such as the verification of the computations, their
challenge by other **clients**; as well as the chance of a **client** to choose
their own set of TM will be left for subsequent updates of the system.

## Table of Contents

* [Design Rationale](#design-rationale)
* [System Components](#components)
* [User Stories](#user-stories)
* [JSON API Specifications](#json-api-specifications)

## Design Rationale

Where we answer to high level questions regarding the design of this system:

* Why this system, and not a blockchain like Bitcoin or Ethereum?
* Why different services (APIs)?
* Why a Patricia Merkle Trie for the storage?
* Can we evolve this distributed system into a truly decentralized one?

### Why this system, and not a blockchain like Bitcoin or Ethrereum?

The spirit of _AutomaCoin_ is to perform _wasteless computation_. Today, a
number of blockchain systems base their minting of blocks with a process called
_Proof of Work_. Such task is using computing power to find a _hash_ given some
_difficulty_. The latter guarantees the existence of a truly decentralized
system, that is, there is no central authority, as cryptography is the ultimate
guarantor of true.

In our first iteration of _AutomaCoin_, we find a number of challenges that
prevent us to take off with a system without any central authority at all:

* Validation: The success of _Proof of Work_ is that, while it is hard to
obtain a value satisfying the challenge, verifying such value is computationally
simple. Currently, it is necessary to find a way to validate the results
obtained for a set of TMs, other than executing the very TMs and comparing.

* Block minting: The difficulty of a certain block is revisited in its chain
depending on the time it takes to produce them. That is, we can infer the
_hashrate_ able to produce these blocks, and adjust the _difficulty_ such that
blocks are produced with a known interval. This calculation is practical due
to the uniform distribution of hashes and the proportional probability of
finding one obeying the required features given by the _difficulty_ value (ex:
_by varying your nonce, find a hash from this set of values that starts with
five zeroes_). In the case of the space of TMs, given a set of machines, even if
we set a maximum number of steps to consider that a machine won't halt, we
cannot forecast the average number of steps a set of machines has. One of the
goals of this **Beaver Project** is explore whether we can find by practical
means measures to help us on our exploration.

* Awarding: Both in Bitcoin, as well as Ethereum, the rewards for computing a
block are very clear and hardcoded into their protocol. A change into the
awarding rate implies what is called as a _hard fork_, taking extensive
coordination, as the cryptocurrency mined at block 1 is the same unit as the one
mined at block 3MM and so on. We desire to have an _era_ system to be more
malleable to changes in our design until we find an optimate combination that
balances discovery of TM output against rewarding.

### Why different services (APIs)?

* Provides a quick, easy to agree communications standard, that we can evolve
in further versions.

* Enable developers to come up with their own approaches to the solution. In
other words, as long as a **client** complies with the API protocol to
communicate with the **pool manager**, then anybody can implement their own
client, with the optimizations they see fit in the language they prefer. We
envision that eventually, given that the computations will be rewarded,
optimum ways to compute the TMs will emerge.

* Facilitate specialization from the contributors.

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

### Can we evolve this distributed system into a truly decentralized one?

Three practical problems need to be solved, or more knowledge to be obtained,
in order to use the calculation of set of TMs as a substitute of the _Proof of
Work_:

* To come up with a satisfactory validation mechanism for the computed TMs.
This procedure does not to be a novel one. For example, *clents* can commit
their computed sets, receive an award from the system, and be punished in future
iterations if it is found, after some random validation, they've cheated.

* Find in one of the iterations of this project a practical measure for sets
of TMs able to infer a distribution for the time it would take to solve these
machines (i.e. find their outputs at halting, or whether they can halt at all).
Today, for blockchains, given a certain _difficulty_, there is an expected
value for the time it would take to arrive at a block, given a _hashrate_.

* Discover the most stable and efficient way to award the computation of TMs.

## Components

* Storage System
* Pool Manager
* Web Site
* Clients

### Storage System

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

### Pool Manager

The pool manager is the gateway of the clients with _AutomaCoin_. Clients will
connect with the pool manager, and via authenticated messages they are able to
receive a new set of TMs to compute. In this first iteration is this manager
the one that modifies the _balance_ of an account after receiving (and lightly
verify) a set of solved TMs.

In *v1.0* there will be a single **Pool Manager** running in a hosted server.

````
pool.automacoin.com
````

### Web Site

Responsible to render an `HTML/CSS/JS` webpage, showing statistics of
_AutomaCoin_ to the public.

### Clients

The most interesting part of the System of *v1.0*: with the API released, and
the system in motion, different teams can work optimizing their algorithms to
compute more efficiently their assigned TMs in order to increase their reward.

At a high level, the **client** connects with the **pool manager**, via
authenticated JSON RPC messages, asking for a
new set of TMs to compute. If available in the **storage system**, the **pool
manager** will oblige with a set, and disconnect.

The **client** will perform the computation of TMs, finding their output values,
or determining that they won't halt (given a threshold given by the *pool
manager, to be determined in each subspace of TMs*).

Once computed the whole set, the **client** send these values back to the **pool
manager**, which will only acknowledge the reception. To keep the channel
asynchronous, it is task of the client to further inquire for the awarding on
solving the sent set.

In later versions, the clients will have the chance to perform verifier roles.
In such tasks, they will receive already computed TMs, trying to find one of
those, with a different output to the one given, in order to get a bounty, and
the system to deal with the original sender.

## JSON API Specifications

### Authentication of the requests

The REST API uses the `Authorization` header to pass authentication information,
with the following form

```
Authorization: AutomaCoin ECDSAPublicId:Signature
```

Where the receiver must have the requester's `ECDSAPublicId` registered.

The process to obtain the Signature is as follows:

```
StringToSign = HTTP-Verb + "\n" +
               sha256(Payload) + "\n" +
               resource;

Signature = sha256(StringToSign);
```

If the message does not contain a payload (expected in a `GET` request, for
example), then `sha256(Payload)` must be left blank.

### Storage System

The following REST API methods can only be performed by the **pool manager**.
It is imperative the public key of the latter to be registered in the
**storage system**.

#### `GET /uncomputed-tms/<max-tms>`

Obtains a list up to `max-tms` of Turing machines to be distributed to
requesting clients.

````bash
## Request

## No payload

## Response

  "data": {
    "version": "v1.0",
    "tms": [
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
      },
      {
        ...
      }
    ]
  },
}

````

#### `POST /assigned-tms/<user>`

````bash
## Request

{
  "data": {
    "version": "v1.0",
    "set-id": "145",
    "tms": [
      "66154f2bab17cc6f0de81f0b121a4c9979ba993aa55fd1b9a372509084148a96",
      ...
    ]
  },
}


## Response

  "data": {
    "OK"
  },
}
````

#### `POST /computed-tms/<user>`

Writes to the database the obtained computed TMs

````bash
## Request

{
  "data": {
    "version": "v1.0",
    "set-id": "145",
    "tms": [
      {
        "id": "66154f2bab17cc6f0de81f0b121a4c9979ba993aa55fd1b9a372509084148a96",
        "result": "10001"
      },
      {
        "id": "9cf77ea9393297fea3d48c3e6c1a454d73f812dc429bef923ac725c62484f396",
        "result": "H"
      },
      ...
    ]
  },
}

## Response

  "data": {
    "OK"
  },
}
````

#### `POST /award/<user>`

````bash
## Request

{
  "data": {
    "version": "v1.0",
    "award": "0.02",
    "era": "15",
    "set-id": "145",
    "timestamp": "1546414363",
    "nonce": "14"
  },
}

## Response

  "data": {
    "OK"
  },
}
````

### Pool Manager

The following REST API methods are to be performed by the **client**

#### `GET /tm-set`

Obtains a new set of Turing Machines to compute. The client must submit its
assigned TM set before getting a new one.

````bash
## Request

## No payload

## Response

  "data": {
    "version": "v1.0",
    "set-id": "145",
    "tms": [
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
      },
      {
        ...
      }
    ]
  },
}

````

#### `POST /computed-tm-set`

Sends to the **pool manager** the computed set of TMs. The manager will
verify that all the assigned TMs have been computed.

````bash
## Request

{
  "data": {
    "version": "v1.0",
    "set-id": "145",
    "tms": [
      {
        "id": "66154f2bab17cc6f0de81f0b121a4c9979ba993aa55fd1b9a372509084148a96",
        "result": "10001"
      },
      {
        "id": "9cf77ea9393297fea3d48c3e6c1a454d73f812dc429bef923ac725c62484f396",
        "result": "H"
      },
      ...
    ]
  },
}

## Response

  "data": {
    "OK"
  },
}
````

#### `GET /award/<set-id>`

Queries the **pool manager** for the award obtained for a certain set of TMs

````bash
## Request

{
  "data": {
    "version": "v1.0",
    "award": "0.02",
    "era": "15",
    "set-id": "145",
    "timestamp": "1546414363",
    "nonce": "14"
  },
}

## Response

  "data": {
    "OK"
  },
}
````