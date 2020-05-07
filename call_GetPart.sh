#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service PartService cli --call GetPart  --file RequestGetPart.json Unaryi
#OUTPUT
#{
#  "item": {
#    "id": "b1f353cf-7c08-4d2e-af35-f66fd4965c16",
#    "mnfId": "390a13d0-ca93-4766-809b-a9994b964f38",
#    "vendorCode": "GazelNEXT",
#    "createdAt": "2020-05-06T19:59:02.714501Z"
#  }
#}
