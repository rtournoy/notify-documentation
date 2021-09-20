---
title: "Notification Patterns"
description:
date: 2018-12-20
layout: overview
---

*Notify* notifications are designed to be sent and received using the [W3C Linked Data Notifications (LDN)](https://www.w3.org/TR/2017/REC-ldn-20170502/) standard. Payloads have a predictable structure, based primarily
on [Activity Streams 2.0](https://www.w3.org/TR/activitystreams-core/), with some additional vocabularies included for particular properties.

## Core Payload
All *Notify* payloads define an *Activity Streams 2.0* `activity`, and include other properties from the *Notify* context. They may also, optionally, include other properties from other contexts. The following properties from Activity Streams 2.0 are used consistently in all the
notification patterns:

* **The `activity`** property must contain the following properties:
  * **`id`:** This must be a URI, and the use of URN:UUID is recommended. An HTTP URI may be used, but in such cases the URI should resolve to a resource which represents the activity.
  * **`type`:** This should include one of the [Activity Stream 2.0 Activity Types](https://www.w3.org/TR/activitystreams-vocabulary/). It may (depending on the activity) also include a type from the [Notify Activity Types vocabulary](/vocabularies/activity_types/)
  * **`origin`:** The originator of the activity, typically the service responsible for *sending* the notification.
  * **`target`:** The intended destination of the activity, typically the service which *consumes* the notification.
  * **`object`:** This should be the focus of the activity. Other object properties may appear in notifications, as properties of other properties.
* **The `activity`** property may (and often will) contain the following properties:
    * **`actor`:** This identifies the party or process that initiated the activity.

The following notification patterns are defined to be widely reusable. Their re-use is illustrated in the [example scenarios](/scenarios/).