#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service MnfService cli --call ListMnf  --file RequestListMnf.json UnaryRepeated
# OUTPUT
#{
#  "items": [
#    {
#      "id": "2fb7f969-2fdd-4f43-a501-4e9e3582ebd7",
#      "name": "Continental",
#      "createdAt": "2020-05-04T12:31:14.383516Z"
#    },
#    {
#      "id": "3b0ee2e5-7629-4256-aba3-526e797ca89a",
#      "name": "Yokohama",
#      "createdAt": "2020-05-04T12:31:14.383516Z"
#    },
#    {
#      "id": "4bf9a24f-dcb3-41ec-99b9-0e3c4e27bb93",
#      "name": "Michlen",
#      "createdAt": "2020-05-04T12:31:14.383516Z"
#    },
#    {
#      "id": "3824a6ee-2020-4df1-8548-22651ac9f714",
#      "name": "Pirelli",
#      "createdAt": "2020-05-04T12:31:14.383516Z"
#    },
#    {
#      "id": "eb9c98ea-4e29-4549-862d-837119082be5",
#      "name": "TOYO",
#      "createdAt": "2020-05-04T12:31:14.383516Z"
#    },
#    {
#      "id": "6926e659-564e-4c6d-adbb-07c9023fdbb1",
#      "name": "HANKOOK",
#      "createdAt": "2020-05-04T12:31:14.383516Z"
#    },
#    {
#      "id": "18a0add4-8c00-4e02-b567-a5fcb1c8c744",
#      "name": "GOODYEAR",
#      "createdAt": "2020-05-04T12:31:14.383516Z"
#    },
#    {
#      "id": "63c721af-1c59-4ddd-90a4-9744f02b6ee6",
#      "name": "Bridgestone",
#      "createdAt": "2020-05-04T12:31:14.383516Z"
#    },
#    {
#      "id": "9b92b546-74be-4628-8d99-b19e171beda1",
#      "name": "DUNLOP",
#      "createdAt": "2020-05-04T12:31:14.383516Z"
#    },
#    {
#      "id": "a0d0fe7e-6f44-4a79-9d4c-df951a0f8631",
#      "name": "KUMHO",
#      "createdAt": "2020-05-04T12:31:14.383516Z"
#    }
#  ]
#}

