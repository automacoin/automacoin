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
_difficulty_. The latter guarantees the existance of a truly decentralized
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
goals of this **Beaver Project** is explore whether we can have find by
practical means measures to help us on our exploration.

* Awarding: Both in Bitcoin, as well as Ethereum, the rewards for computing a
block are very clear and hardcoded into their protocol. A change into the
awarding rate implies what is called as a _hard fork_, taking extensive
coordination, as the cryptocurrency mined at block 1 is the same unit as the one
mined at block 3MM and so on. We desire to have an _era_ system to be more
malleable to changes in our design until we find an optimate combination that
balances discovery of TM output against rewarding.

### Why different services (APIs)?

TODO 03
	- Enable developers to come up with their own approaches to the solution

### Why a Patricia Merkle Trie for the storage?

TODO 04

### Can we evolve this distributed system into a truly decentralized one?

TODO 05

## Components

* Storage System
* Pool Manager
* Web Site
* Clients

### Storage System

TODO 06

### Pool Manager

TODO 07

### Web Site

TODO 08

### Clients

TODO 09

## User Stories

### Client: Computation of a new set of TMs

TODO 10

### Client: Querying the result of a set of TMs

TODO 11

### Website: ?

TODO 12

## JSON API Specifications

### Pool Manager

TODO 13

### Storage System

TODO 14

### Web Site

TODO 15

### Clients

TODO 16
