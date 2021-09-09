---
title: Announce Review
description: 'The overlay journal notifies the repository of a new review '
date: "2021-03-08"
scope: notify
sender: actor_1
pattern: announcement/announcement
payload:
    "@context": ["sorg","ldp","ietf","nat","nrr"]
    id: "urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f"
    type: ["Announce","coar-notify:ReviewSuccess"]
    origin:
        lookup: "overlay-journal"
    target:
        lookup: "repository"
    object:
        lookup: "journal-review"
    context:
        lookup: generic-object-repository
    actor:
        lookup: reviewer
---

