---
title: "Announcement (in response to offer)"
date: "2021-03-08"
description: "This pattern is used to announce the outcome of an activity, sometimes (but not always) linking an original resource (referenced in `context`) to a new, related resource (referenced in `object`), in response to an offer made in a previous notification."
layout: pattern_example
status: [review,draft]
weight: 2
payload:
    id: "urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f"
    type: "Announce"
    actor:
        lookup: "generic-service"
    origin:
        lookup: "generic-service"
    target:
        lookup: "generic-organisation"
    object:
        lookup: generic-object-service
    context:
        lookup: generic-object-repository
    in_reply_to:
        id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
---