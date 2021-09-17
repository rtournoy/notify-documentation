---
title: Announce Review
description: 'The review service notifies the repository of a new review '
date: "2021-03-08"
scope: notify
sender: actor_2
pattern: announcement/announcement_in_reply_to
payload:
    "@context": ["sorg","ldp","ietf","nat","nrr"]
    id: "urn:uuid:d4ed8e1d-d6ce-4160-9f84-2546a72376a1"
    type: ["Announce","coar-notify:ReviewAction"]
    origin:
        lookup: "review-service"
    target:
        lookup: repository
    object:
        lookup: review
    context:
        lookup: preprint
    actor:
        lookup: reviewer-2
    in_reply_to:
        id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
---

