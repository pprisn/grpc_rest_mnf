#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service PartService cli --call CreatePart --file RequestCreatePart.json Unary
#OUTPUT NEW ID
#{
#  "id": "c029f11e-3532-4a0a-b165-dbdaad17ecde"
#}
#
#mnf=# select * from part where vendor_code='EMS';
#                  id                  |                mnf_id                | vendor_code |          created_at          | deleted_at
#--------------------------------------+--------------------------------------+-------------+------------------------------+------------
# c029f11e-3532-4a0a-b165-dbdaad17ecde | 7da4bd51-67b6-4b82-ad48-fa2bfa37a503 | EMS         | 2020-05-06 22:15:05.22501+03 |
#(1 строка)
#
