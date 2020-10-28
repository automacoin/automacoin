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