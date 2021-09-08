---
title: Change Log
description:
date: 2021-05-07
---

#### 2021-09-08


#### 2021-07-13
Removed the `ldp:` prefix from the `inbox` property as it is already defined in Activity Streams 2

#### 2021-07-07
Removed the nested `object` (causing an unnecessary 'blank node') from the `coar-notify:reviews` and `coar-notify:endorses` patterns.

So, what was expressed, for example, as:
```json
"object": {
    "coar-notify:reviews": {
      "object": {
        "id": "https://repository.org/preprint/201203/421/",
        "ietf:cite-as": "https://doi.org/10.5555/12345680"
      }
    },
```

is now expressed as:
```json
"object": {
    "coar-notify:reviews": {
        "id": "https://repository.org/preprint/201203/421/",
        "ietf:cite-as": "https://doi.org/10.5555/12345680"
    },
```

#### 2021-07-01
Fixed typo in scenarios, replacing incorrect `nat:` namespace with `coar-notify:`.

Renamed the terms in the `coar-notify` vocabulary to use camelCase. So, for example:

`endorsement-success` has become `EndorsementSuccess`.

Removed some unused terms from the vocabulary.


#### 2021-06-23
Fixed typo in scenarios, replacing incorrect `nrr:` namespace with `coar-notify:`.

#### 2021-06-01
We have made some changes to the Notify documentation. The changes are to improve the notification payloads in terms of linked data (ensuring that they can be expressed as RDF) and so the changes should **not** have any breaking impact on current software development.

###### 1. Context
The structure of the notification payloads is essentially unchanged, except that we have created a first draft Notify JSON-LD context ([http://purl.org/coar/notify](http://purl.org/coar/notify)), and changed the example JSON-LD in the patterns and scenarios to reflect this. All patterns now have this as the context property:

```json
{
  "@context": [
    "https://www.w3.org/ns/activitystreams",
    "http://purl.org/coar/notify"
  ]
}
```

###### 2. Strictness of the descriptions
   The descriptions of the payloads were previously described in strict, 'IETF' language (e.g. MUST, SHOULD etc.). This language implied a strict 'schema' which is not really the case here, since we are working in a linked-data paradigm. It became clear that the imposition of a pseudo schema could have unforeseeable consequences in later re-use of the patterns. We expect these patterns to evolve with use, as the community becomes more skilled and ambitions in their application.

As such, these patterns are best thought of as *conventions*, rather than *standards*, relying on the underlying standards of W3C LDN, W3C Activity Streams 2.0 et. al.

For this reason, we have relaxed the language used to describe the patterns. However, the structures and patterns described have not changed (with the exception of the `@context` property as described above.

#### 2021-05-14
Changed incorrect value 'System' to 'Service' in all examples of `origin` and `target` properties used in patterns and scenarios.

#### 2021-05-13
* Changed URI used in `@context` property for Activity Streams 2.0 from 'https://www.w3.org/ns/activitystreams#' to 'https://www.w3.org/ns/activitystreams' (i.e. removed the '#' suffix)

#### 2021-05-07
* Changed docs to indicate that an array SHOULD be used for the `type` property.
* Updated all examples to use array values for the `type` property.

