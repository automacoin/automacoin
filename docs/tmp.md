

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