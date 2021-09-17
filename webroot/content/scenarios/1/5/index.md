---
title: Announce Endorsement
description: 'The overlay journal notifies the repository of its endorsement'
date: "2021-03-08"
scope: notify
sender: actor_1
pattern: announcement/announcement
payload:
    "@context": ["sorg","ldp","ietf","nat","nrr"]
    id: "urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f"
    type: ["Announce","coar-notify:EndorsementAction"]
    origin:
        lookup: "overlay-journal"
    target:
        lookup: repository
    object:
        lookup: endorsement
    actor:
        lookup: overlay-journal
    context:
        lookup: preprint
---

