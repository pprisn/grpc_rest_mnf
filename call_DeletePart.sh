#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service PartService cli --call DeletePart  --file RequestDeletePart.json Unary
#OUTPUT
#{}
#mnf=# select * from part where mnf_id='eb9c98ea-4e29-4549-862d-837119082be5';
#                  id                  |                mnf_id                | vendor_code |          created_at          |          deleted_at           
#--------------------------------------+--------------------------------------+-------------+------------------------------+-------------------------------
# 6934e2a3-bc5e-4e3f-9fdd-9b07ebd5e623 | eb9c98ea-4e29-4549-862d-837119082be5 | TS01345     | 2020-05-04 15:48:55.12077+03 | 2020-05-08 18:39:40.820661+03
# 7c78f55e-0e98-495a-b2a5-811fbe346a8a | eb9c98ea-4e29-4549-862d-837119082be5 | TS01375     | 2020-05-04 15:48:55.12077+03 | 
# 452196fc-0dcb-412d-bbcb-2f14b15d398a | eb9c98ea-4e29-4549-862d-837119082be5 | TS01387     | 2020-05-04 15:48:55.12077+03 | 
#(3 строки)

