#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service PartService cli --call UpdateParts  --file RequestUpdateParts.json Unary
#OUTPUT
#{}
#mnf=# select * from part where mnf_id='971a3522-47f3-4dd6-b075-5c60f3f60136';
#                  id                  |                mnf_id                |  vendor_code   |          created_at           | deleted_at 
#--------------------------------------+--------------------------------------+----------------+-------------------------------+------------
# b582b36d-d00f-4e0d-bd84-7d3d894618b7 | 971a3522-47f3-4dd6-b075-5c60f3f60136 | VESTA_update   | 2020-05-06 22:59:02.704834+03 | 
# 3b5ce0eb-7544-4660-9b84-f83319b3e3c8 | 971a3522-47f3-4dd6-b075-5c60f3f60136 | VAZ2115_update | 2020-05-06 22:59:02.651218+03 | 

