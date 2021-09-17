---
title: Request Review
description: 'The repository offers the resource for review by the review service. '
date: "2021-03-08"
scope: notify
sender: actor_1
pattern: offer/offer
payload:
    "@context": ["sorg","ldp","ietf","nat"]
    id: "urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd"
    type: ["Offer","coar-notify:ReviewAction"]
    actor:
        lookup: author
    origin:
        lookup: repository
    target:
        lookup: "review-service"
    object:
        lookup: preprint
    actor:
        lookup: author
---


