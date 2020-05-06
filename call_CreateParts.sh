#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service PartService cli --call CreateParts --file RequestCreateParts.json ClientStreaming
#OUTPUT NEW []ID
#{
#  "ids": [
#    "3b5ce0eb-7544-4660-9b84-f83319b3e3c8",
#    "b582b36d-d00f-4e0d-bd84-7d3d894618b7",
#    "b1f353cf-7c08-4d2e-af35-f66fd4965c16"
#  ]
#}
#INFO[2020-05-06T22:59:02+03:00] Peer:127.0.0.1:57618 INSERT part (id,mnf_id, vendor_code) values (3b5ce0eb-7544-4660-9b84-f83319b3e3c8, 971a3522-47f3-4dd6-b075-5c60f3f60136, Ваз2115) 
#INFO[2020-05-06T22:59:02+03:00] Peer:127.0.0.1:57618 INSERT part (id,mnf_id, vendor_code) values (b582b36d-d00f-4e0d-bd84-7d3d894618b7, 971a3522-47f3-4dd6-b075-5c60f3f60136, VestaSport) 
#INFO[2020-05-06T22:59:02+03:00] Peer:127.0.0.1:57618 INSERT part (id,mnf_id, vendor_code) values (b1f353cf-7c08-4d2e-af35-f66fd4965c16, 390a13d0-ca93-4766-809b-a9994b964f38, GazelNEXT) 
#
#mnf=# select * from part where created_at >= '2020-05-06';
#                  id                  |                mnf_id                | vendor_code |          created_at           | deleted_at 
#--------------------------------------+--------------------------------------+-------------+-------------------------------+------------
# c029f11e-3532-4a0a-b165-dbdaad17ecde | 7da4bd51-67b6-4b82-ad48-fa2bfa37a503 | EMS         | 2020-05-06 22:15:05.22501+03  | 
# 3b5ce0eb-7544-4660-9b84-f83319b3e3c8 | 971a3522-47f3-4dd6-b075-5c60f3f60136 | Ваз2115     | 2020-05-06 22:59:02.651218+03 | 
# b582b36d-d00f-4e0d-bd84-7d3d894618b7 | 971a3522-47f3-4dd6-b075-5c60f3f60136 | VestaSport  | 2020-05-06 22:59:02.704834+03 | 
# b1f353cf-7c08-4d2e-af35-f66fd4965c16 | 390a13d0-ca93-4766-809b-a9994b964f38 | GazelNEXT   | 2020-05-06 22:59:02.714501+03 | 
#(4 строки)
#
