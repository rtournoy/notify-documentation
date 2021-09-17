---
title: Announce Review
description: 'The repository notifies the overlay journal of a new review '
date: "2021-03-08"
scope: notify
sender: actor_2
pattern: announcement/announcement
payload:
    "@context": ["sorg","ldp","ietf","nat","nrr"]
    id: "urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f"
    type: ["Announce","coar-notify:ReviewAction"]
    actor:
        lookup: research-organisation
    origin:
        lookup: "repository"
    target:
        lookup: "overlay-journal"
    object:
        lookup: "repository-review"
    context:
        lookup: preprint
---

