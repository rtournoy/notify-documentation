---
title: Announce Review
description: 'The review service notifies the repository of a new review '
date: "2021-03-08"
scope: notify
sender: actor_2
pattern: announcement/announcement_in_reply_to
payload:
    "@context": ["sorg","ldp","ietf","nat","nrr"]
    id: "urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f"
    type: ["Announce","coar-notify:ReviewAction"]
    actor:
        lookup: overlay-journal
    origin:
        lookup: "overlay-journal"
    target:
        lookup: repository
    object:
        lookup: "journal-review"
    context:
        lookup: preprint
    in_reply_to:
        id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
        # type: ["Offer","coar-notify:ReviewAction"]
---

