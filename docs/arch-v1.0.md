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
from now, *v1.0*.

In one single paragraph, this is a multi-service architecture, each service
communicating to each other via JSON API. There is a huge size (Terabyte, maybe
Petabyte) *database* of Turing Machines of certain features, user accounts, and
indexes sorting and grouping them. It will be the task of a *pool manager* to
determine which set of combinations of TMs will be processed upon the request
of a *client*. Finished the computations of the *client*, the *pool manager*
will receive the set of TMs, store the results into the *database* and assign
an _AutomaCoin_ reward to the *client*. The latter, as well as any other user,
can verify the space of TMs solved, the awards given and other data to a *web
site* feeding from the *database*.

Fine grained discussions, such as the verification of the computations, their
challenge by other *clients*; as well as the chance of a *client* to choose
their own set of TM will be left for subsequent updates of the system.

## Table of Contents

* Design Rationale
* System Components
* User Stories
* JSON API Specifications

## Design Rationale

Answers to high level questions regarding the design of this system:

* Why this system, and not a blockchain like Bitcoin or Ethereum?
* Why different services (APIs)?
* Why a Patricia Merkle Trie for the storage?
* Can we evolve this distributed system into a truly decentralized one?

### Why this system, and not a blockchain like Bitcoin or Ethrereum?

TODO 02

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
