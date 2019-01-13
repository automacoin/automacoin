https://staking.cardano.org/

incentivization viene de la validacion

inicialmente el pool manager con un sampling de los resultados de un cliente
dice "esto esta correcto con este grado de certeza", y envia
los resultados al storage system, premiando al cliente

en el futuro, el leader o firmador tiene ese trabajo

en que varia?

En casi nada:
  - recibe el paquete de un cliente
  - lo verifica con cierto sampleo
  - asigna la recompensa
  - almacena los resultados de frecuencias, y la recompensa

La diferencia es que no seria un API call al storage
system, sino que esto ya es un bloque firmado, que se propaga
en la red.

De ahi en el documento "vision para automacoin como un blockchain",
ponemos la evolucion

- pool manager -> se transforma en leader, con un sistema similar a cardano
- storage system -> se distribuye, donde los clientes administran (segun configuracion)
  shards de la base de datos.
