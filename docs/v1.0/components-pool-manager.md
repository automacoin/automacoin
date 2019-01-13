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