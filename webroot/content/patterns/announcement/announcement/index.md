---
title: "Announcement (standalone)"
date: "2021-03-08"
description: "This pattern is used to announce the outcome of an activity, sometimes (but not always) linking an original resource (referenced in `context`) to a new, related resource (referenced in `object`)."
layout: pattern_example
status: [review,draft]
weight: 1
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
---