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
