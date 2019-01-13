### Authentication of the requests

The REST API uses the `Authorization` header to pass authentication information,
with the following form

```
Authorization: AutomaCoin ECDSAPublicId:Signature
```

Where the receiver must have the requester's `ECDSAPublicId` registered.

The process to obtain the Signature is as follows:

```
StringToSign = HTTP-Verb + "\n" +
               sha256(Payload) + "\n" +
               resource;

Signature = sha256(StringToSign);
```

If the message does not contain a payload (expected in a `GET` request, for
example), then `sha256(Payload)` must be left blank.