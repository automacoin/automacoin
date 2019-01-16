## AutomaCoin
## System Architecture v1.0

### Abstract

The present document describes the first version of the _AutomaCoin_ system,
from now, **v1.0**.

The purpose of this system is to explore the Turing Machine space (of different
types, differing in number of states, symbols and, dimensions). This exploration
will help us to know which TMs halt, and their outputs. These outputs will help
us building increasingly accurate estimations of the universal distribution.

The scope of this project on its first version is constrained to understand the
challenges of distributing these tasks over a number of clients, enabling them
to be awarded in a currency we are calling _AutomaCoin_.

This is a multi service architecture, each service communicating to each other via
JSON API. There is a huge size **Database** of Turing Machines of certain features,
containing user accounts and obtained string frequencies by TM space, alongside
indexes sorting and grouping them. It will be the task of a **Pool manager**
to determine which set of combinations of TMs will be processed upon the request
of a **Client**. Finished the computations of the **Client**, the **Pool
manager** will receive the set of TMs, store the results into the **Database**
and assign an _AutomaCoin_ reward to the **Client**. The latter, as well as any
other user, can verify the space of TMs solved, the awards given and other data
to a **Web site** feeding from the **Database**.

Bigger challenges in this iteration:

* Arrive to a satisfactory resolution and **Protocol** for the following items:
  * **Verification** of solved computation sets.
  * Ability to **Challenge** solved computation sets.
  * **Awarding** of _AutomaCoin_ tokens.
  * **Storage Categories** for the clients and subsequent
    **Distribution of the database**.
* Have discussions on how to build a ledger that leverages over this system.
* Face total decentralization of the system in the future, and roadmap accordingly.

### Table of Contents

* Design Rationale
  * [Distributed System](design-rationale-distributed-system.md)
  * [Challenges to become a Decentralized System](design-rationale-challenges-decentralized-system.md)
  * [Storage: Now and Vision](design-rationale-storage-now-and-vision.md)
  * [Incentivization](design-rationale-incentivization.md)
* Use Cases
  * [Computing a Set of Turing Machines](use-case-computing-a-set-of-tms.md)
  * [Challenging the Result of a Set of Computations](use-case-challenging-results.md)
  * [Querying the Database](use-case-querying-the-database.md)
  * [Storage Discussion](use-case-storage-discussion.md)
* System Components
  * [Database](components-database.md)
  * [Pool Manager](components-pool-manager.md)
  * [Website](components-website.md)
  * [Client](components-client.md)
* JSON API Specifications
  * [Authentication](json-api-authenticating-requests.md)
  * [Database](json-api-database.md)
  * [Pool Manager](json-api-pool-manager.md)
  * [Website](json-api-website.md)
