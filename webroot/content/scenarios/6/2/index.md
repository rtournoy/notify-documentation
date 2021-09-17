---
title: Request for Ingest
description: "The overlay journal requests that a repository ingest an author's submitted manuscript"
date: "2021-03-08"
scope: "notify"
sender: "actor_1"
pattern: "offer/offer"
payload:
    "@context": ["sorg","ldp","ietf","nat"]
    id: "urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd"
    type: ["Offer","coar-notify:IngestAction"]
    actor:
        lookup: overlay-journal
    origin:
        lookup: "overlay-journal"
    target:
        lookup: "repository"
    object:
        lookup: submission
---

