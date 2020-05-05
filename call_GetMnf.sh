#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service MnfService cli --call GetMnf  --file RequestGetMnf.json Unary
#OUTPUT
#{
#  "item": {
#    "id": "6926e659-564e-4c6d-adbb-07c9023fdbb1",
#    "name": "HANKOOK",
#    "createdAt": "2020-05-04T12:31:14.383516Z"
#  }
#}

