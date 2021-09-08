---
title: Endorsement Announcement (with link)
description: The overlay journal notifies the aggregator that a resource has been endorsed, and provides a link to evidence of that endorsement.
date: "2021-03-08"
scope: notify
sender: actor_1
pattern: announcement/announcement
payload:
    "@context": ["sorg","ldp","ietf","nat","nrr"]
    id: "urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f"
    type: ["Announce","coar-notify:EndorsementSuccess"]
    origin:
        lookup: "overlay-journal"
    target:
        lookup: "aggregation-service"
    object:
        lookup: endorsement
    actor: ""
---

