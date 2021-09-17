---
title: Change Log
description:
date: 2021-05-07
---

#### 2021-09-17
This is a significant revision.

##### General changes:
These general changes have been made across all patterns and scenarios.
1. All notifications now contain an `object` (as required by Activity Streams 2.0). If the notification is for an activity which has caused a resource to be created or updated, then that resource will normally be referenced (with an HTTP URI) in the `object`. If the notification has not caused a resource to be created or updated, then the object will normally contain just a local ID (e.g. a UUID URI) and possibly some simple metadata.
2. All notifications now contain an `actor`. In some cases this is a person - e.g. a reviewer, however in most cases this is an organisation or a service. 
3. The nested `inbox` property has been removed from all instances of the `actor` property
4. `inReplyTo`, when used, no longer contains any references to resources - it just references the `id` of the `activity` which it is "replying" to.
5. Notifications of activities which pertain to a existing resource may reference that existing resource in a `context` property. For example, a notification of a new review of a preprint would reference the review in the `object` property and the preprint in a `context` property.
6. The vocabulary and related COAR @context for Notify has been significantly simplified.
7. Syntax has generally been simplified. Previously we were suggesting using arrays anytime there might be more than one value for a property. Now we use the simpler scalar form of property values where there is a single value in the example. This applies to the activity's `type`, `inReplyTo`, and the `type` in `origin`, `object`, `target` and `actor`.

##### Changes to patterns

###### Offer pattern
`actor` added

###### Acknowledgement patterns
* `inReplyTo` changed to just contain URI to `offer` activity
* `context` added
* `actor` added

###### Retraction pattern
* `context` added
* `actor` added

###### Announcement pattern
* `context` added

###### Announcement in reply to pattern
* `inReplyTo` changed to just contain URI to `offer` activity
* `context` added
* `actor` added

##### Changes to the vocabularies
The following terms from the COAR Notify vocabulary have changed:
* `coar-notify:ReviewSuccess` has become `coar-notify:ReviewAction`
* `coar-notify:EndorsementSuccess` has become `coar-notify:EndorsementAction`
* `coar-notify:IngestSuccess` has become `coar-notify:IngestAction`

The following terms from the COAR Notify vocabulary have been removed:
* `ReviewRequest`
* `ReviewFailure`
* `EndorsementRequest`
* `EndorsementFailure`
* `IngestRequest`
* `reviews`
* `endorses`

##### Changes to scenarios

###### Scenario 1
* `object` in steps 3 and 5 no longer references the resource that was reviewed
* `context` added in steps 3 and 5
* `actor` added

###### Scenario 2
* `inReplyTo` changed to just contain URI to `offer` activity in steps 4 and 6
* `object` in steps 4 and 6 no longer references the resource that was reviewed
* `context` added in steps 4 and 6
* `actor` added

###### Scenario 3
* `object` in steps 2 and 4 no longer references the resource that was reviewed
* `context` added in steps 2 and 4
* `actor` added

###### Scenario 4
* `object` in step 3 no longer references the resource that was reviewed
* `context` added in step 3
* `actor` added

###### Scenario 5
* `inReplyTo` changed to just contain URI to `offer` activity in steps 3 and 5
* `object` added in step 3
* `object` in step 5 no longer references the resource that was reviewed
* `context` added in steps 3 and 5
* `actor` added

###### Scenario 6
* `inReplyTo` changed to just contain URI to `offer` activity in step 4
* `object` in steps 7 and 9 no longer references the resource that was reviewed
* `context` added in steps 4, 7 and 9
* `actor` added

###### Scenario 7
* `object` in step 2 no longer references the resource that was reviewed
* `context` added in step 2
* `actor` added

###### Scenario 8
* `object` in step 2 no longer references the resource that was reviewed
* `context` added in step 2
* `actor` added

###### Scenario 9
* `inReplyTo` changed to just contain URI to `offer` activity in steps 5 and 9
* `object` in steps 5 and 9 no longer references the resource that was reviewed
* `context` added in steps 5 and 9
* `actor` added



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

