## Use case: Computing a Set of Turing Machines

This is the core activity of this system.

### Preliminaries

* The ECDSA _public key_ of the **Client** should be registered by the
  **Pool Manager**
  * Therefore, in **v1.0**, clients to be used by individuals or organizations
    must be coordinated in advance to use this system, as there is not an
    automatic (i.e. _not human_) process to add _public keys_ to the system.

* The **Pool Manager** (**PM**) must prepare in advance from the **Database**
  the Turing Machines to be processed.
  * The former can be accomplished by the **PM** by use of the
    `GET /uncomputed-tms/<max-tms>` API.
  * The **PM** will avoid sending TMs which are **known to not halt**.
    To that goal, the **PM** will contain _criteria_ to determine this fact.
    For example: TMs that do not contain the transition to the _HALT State_.
    Any machine responded by the **Database** via `POST /assigned-tms`, can
    be sent back using `POST /non-halting`, using the document payload.
  * It is important to mention over and over this document that the TMs must
    be enumerated using integers, and for a metodology to translate the
    integer id of a TM to the proper TM back and forth. As the number of TMs
    in a space can be of the number of millions, if not billions, even
    trillions. Therefore, having the TMs assigned in a non sucessive

 way can
    lead to an overuse of storage space.

### Use Case algorithm

* The **Client** requests a set of Turing Machines to the **Pool Manager**.
  * The API request is `GET /tm-set`.

* The **PM** responds to the **Client** with a list of TMs to be computed.
  * The **PM** sends a `POST /assigned-tms/<user>` to the **Database**. This
    step prevents the **Database** to send the TMs to a different request.
  * A Time to Live (TTL) time should be considered for an assigned TM inside
    the **Database**.

* The **Client** performs the computation of each of the received TMs of
  the set, recording the output strings of the TMs that halted.
  * The **Client** will store ....


... Pool manager verifies
   .... see discussions on verifications TBD
   .... compute the award for the client

... verification OK, sending to the database the following
  ... id of the client
  ... hash of the root of the results
  ... frequencies of the obtained strings
  ... award given to the client
  ... see discussion on challenging results

### Summary of results

... computation
  ... client id, machine set digest (id start, id end), hash of the results, frequencies, pool manager id, award

... frequencies obtained for a determined space

... client accounts vs automacoin balance

... Donde pongo metadata como tiempo y cantidad de TMs en cada computacion?
