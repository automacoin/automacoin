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

* Design Rationale
  * [Distributed System](design-rationale-distributed-system.md)
  * [Challenges to become a Decentralized System](design-rationale-challenges-decentralized-system.md)
  * [Storage: Now and Vision](design-rationale-storage-now-and-vision.md)
  * [Incentivization](design-rationale-incentivization.md)
* System Components
  * [Database](components-database.md)
  * [Pool Manager](components-pool-manager.md)
  * [Website](components-website.md)
  * [Client](components-client.md)
* Use Cases
  * [Computing a Set of Turing Machines](use-case-computing-a-set-of-tms.md)
  * [Challenging the Result of a Set of Computations](use-case-challenging-results.md)
  * [Querying String Frequencies](use-case-querying-string-frequencies.md)
  * [Querying AutomaCoin Awards](use-case-querying-awards.md)
* JSON API Specifications
  * [Authentication](json-api-authenticating-requests.md)
  * [Database](json-api-database.md)
  * [Pool Manager](json-api-pool-manager.md)
  * [Website](json-api-website.md)
