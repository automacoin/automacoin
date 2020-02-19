### Pool Manager

The following REST API methods are to be performed by the **client**

#### `GET /tm-set`

Obtains a new set of Turing Machines to compute. The client must submit its
assigned TM set before getting a new one.

````bash
## Request

## No payload

## Response

  "data": {
    "version": "v1.0",
    "set-id": "145",
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

#### `POST /tm-set`

Sends to the **pool manager** the computed set of TMs. The manager will
verify that all the assigned TMs have been computed.

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
