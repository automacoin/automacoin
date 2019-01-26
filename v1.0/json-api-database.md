### Database

The following REST API methods can only be performed by the **pool manager**.
It is imperative the public key of the latter to be registered in the
**Database**.

#### `GET /uncomputed-tms/<max-tms>`

Obtains a list up to `max-tms` of Turing machines to be distributed to
requesting clients.

The **Database** should rely on a microservice able to prepare
this list of turing machines.

````bash
## Request

## No payload

## Response

  "data": {
    "version": "v1.0",
    "tms": [
      {
        "A0":"1BL",
        "A1":"1H",
        "B0":"0CR",
        "B1":"0CR",
        "C0":"0DR",
        "C1":"1ER",
        "D0":"0DL",
        "D1":"1AL",
        "E0":"1AR",
        "E1":"1AL"
      },
      {
        ...
      }
    ]
  },
}

````

#### `POST /non-halting`

Sending all the machines determined by the **Pool Manager** to not halt.
For example: TMs that do not contain the transition to the _HALT State_.

````bash
## Request

{
  "data": {
    "version": "v1.0",
    "set-id": "145",
    "tms": [
      {
        "id": "66154f2bab17cc6f0de81f0b121a4c9979ba993aa55fd1b9a372509084148a96"
      },
      {
        "id": "9cf77ea9393297fea3d48c3e6c1a454d73f812dc429bef923ac725c62484f396"
      },
      ...
    ]
  },
}

## Response

  "data": {
    "OK"
  },
}
````

#### `POST /assigned-tms/<user>`

The **Pool Manager** responds to the **Client** with a list of TMs to be computed.

* The **PM** sends a `POST /assigned-tms/<user>` to the **Database**. This
  step prevents the **Database** to send the TMs to a different request.

* A Time to Live (TTL) time should be considered for an assigned TM inside
  the **Database**.

````bash
## Request

{
  "data": {
    "version": "v1.0",
    "set-id": "145",
    "tms": [
      "66154f2bab17cc6f0de81f0b121a4c9979ba993aa55fd1b9a372509084148a96",
      ...
    ]
  },
}


## Response

  "data": {
    "OK"
  },
}
````

#### `POST /computed-tms/<user>`

Writes to the database the obtained computed TMs

````bash
## Request

{
  "data": {
    "version": "v1.0",
    "set-id": "145",
    "tms": [
      {
        "id": "66154f2bab17cc6f0de81f0b121a4c9979ba993aa55fd1b9a372509084148a96",
        "result": "10001"
      },
      {
        "id": "9cf77ea9393297fea3d48c3e6c1a454d73f812dc429bef923ac725c62484f396",
        "result": "H"
      },
      ...
    ]
  },
}

## Response

  "data": {
    "OK"
  },
}
````

#### `POST /award/<user>`

````bash
## Request

{
  "data": {
    "version": "v1.0",
    "award": "0.02",
    "era": "15",
    "set-id": "145",
    "timestamp": "1546414363",
    "nonce": "14"
  },
}

## Response

  "data": {
    "OK"
  },
}
````
