#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service PartService cli --call UpdatePart  --file RequestUpdatePart.json Unary
#OUTPUT
#{}
#mnf=# select * from part where mnf_id='390a13d0-ca93-4766-809b-a9994b964f38';
#--------------------------------------+--------------------------------------+------------------+-------------------------------+------------
#                  id                  |                mnf_id                |   vendor_code    |          created_at           | deleted_at 
#--------------------------------------+--------------------------------------+------------------+-------------------------------+------------
# b1f353cf-7c08-4d2e-af35-f66fd4965c16 | 390a13d0-ca93-4766-809b-a9994b964f38 | GazelNEXT_Update | 2020-05-06 22:59:02.714501+03 | 
#(1 строка)
