### Why different services (APIs)?

* Provide a quick, easy to agree communications standard, that we can evolve
in further versions.

* Enable developers to come up with their own approaches to the solution. In
other words, as long as a **Client** complies with the API protocol to
communicate with the **Pool manager**, then anybody can implement their own
client, with the optimizations they see fit, and in the language they prefer.
We envision that eventually, given that the computations will be rewarded,
optimum ways to compute the TMs will emerge.

* Facilitate specialization from the contributors.

#### Evolution of the services into roles

A big motivation for this project is for us, explore the paths to become
a truly, decentralized system. In order to accomplish that goal, elements
of this system should evolve:

* Generally speaking all nodes have, while having the ability to change
  their role (computer, verificator, leader, archive), in principle, have
  the same potential as any other one.

* The **Database** System must evolve into a Storage functionality
  * Nodes will have the option of keeping a full copy of the database.
  * Nodes can have a partial copy of the database (namely, a _shard_).
  * Nodes can only keep reference of certain _roots_, and from there,
    be able to retrieve the database elements they need on demand.
  * Further discussions need to be applied for co-selector indexes and
    location address

* The **Pool Manager** must become the role of a _leader_.
  * In a number of _Proof of Stake_ consensus algorithms, there are roles
    akin to the one of a _leader_. That is, a node with the responsibility
    to accept transactions and mint blocks.
  * There are several challenges preceding this evolution goal:
    * We need to find the most convenient **Verification** schema, given
      the fact that the problem of computing a set of TMs cannot be verified
      in the same way as as _Proof of Work_ hash, we need to rely on
      statistical sampling.
    * We need to develop a protocol for **Challenging Results**, that is,
      at any moment a client can retrieve the set of TMs that another one
      produced, compute the results, and compare them with either a full
      storage of the results (rather expensive), or a merkle trie hash of
      the head of the results. In the event of finding discrepancies, there
      must exist some form or other of apply penalties to the faulty client.
    * We need to come up with the adecquate **Awarding** structure to prevent
      unwanted economic side effects, such as inflation. Earlier discussions
      pointed to a system of **Eras**, which is akin to putting the year a
      fiat coin (or note) is issued.

#### The roadmap to a ledger (DLT)

The coins will be practically useless if we don't allow for the
users to spend them. Hence, we need to work on a way to **Leverage our
system to enable for a ledger to be built on top of it**.

In this sense, if our system is similar to one with _leaders_ as opposed to
a _Proof of Work_ based one (i.e. the winner in the block race gets to mine),
we can have two clear layers:

* Layer 1: Where the clients compute sets of TMs, are validated and awarded
  (and can be challenged within a reasonable interval).

* Layer 2: Where the _Leaders_ in an iteration, similar to a _block_, alongisde
  registering computations and awarding, attach to the ledger the transactions
  made during that period.