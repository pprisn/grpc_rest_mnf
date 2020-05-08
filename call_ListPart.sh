#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service PartService cli --call ListPart  --file RequestListPart.json UnaryRepeated
# OUTPUT
#{
#  "items": [
#    {
#      "id": "6934e2a3-bc5e-4e3f-9fdd-9b07ebd5e623",
#      "mnfId": "eb9c98ea-4e29-4549-862d-837119082be5",
#      "vendorCode": "TS01345",
#      "createdAt": "2020-05-04T12:48:55.120770Z"
#    },
#    {
#      "id": "7c78f55e-0e98-495a-b2a5-811fbe346a8a",
#      "mnfId": "eb9c98ea-4e29-4549-862d-837119082be5",
#      "vendorCode": "TS01375",
#      "createdAt": "2020-05-04T12:48:55.120770Z"
#    },
#    {
#      "id": "452196fc-0dcb-412d-bbcb-2f14b15d398a",
#      "mnfId": "eb9c98ea-4e29-4549-862d-837119082be5",
#      "vendorCode": "TS01387",
#      "createdAt": "2020-05-04T12:48:55.120770Z"
#    }
#  ]
#}

