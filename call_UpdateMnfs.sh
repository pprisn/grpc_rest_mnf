#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service MnfService cli --call UpdateMnfs --file RequestUpdateMnfs.json ClientStreaming
#OUTPUT  []ID

#mnf=# select * from part_manufacturer;
#                  id                  |        name         |          created_at           
#--------------------------------------+---------------------+-------------------------------
# 2fb7f969-2fdd-4f43-a501-4e9e3582ebd7 | Continental         | 2020-05-04 15:31:14.383516+03
# 3b0ee2e5-7629-4256-aba3-526e797ca89a | Yokohama            | 2020-05-04 15:31:14.383516+03
# 4bf9a24f-dcb3-41ec-99b9-0e3c4e27bb93 | Michlen             | 2020-05-04 15:31:14.383516+03
# 3824a6ee-2020-4df1-8548-22651ac9f714 | Pirelli             | 2020-05-04 15:31:14.383516+03
# eb9c98ea-4e29-4549-862d-837119082be5 | TOYO                | 2020-05-04 15:31:14.383516+03
# 6926e659-564e-4c6d-adbb-07c9023fdbb1 | HANKOOK             | 2020-05-04 15:31:14.383516+03
# 18a0add4-8c00-4e02-b567-a5fcb1c8c744 | GOODYEAR            | 2020-05-04 15:31:14.383516+03
# 63c721af-1c59-4ddd-90a4-9744f02b6ee6 | Bridgestone         | 2020-05-04 15:31:14.383516+03
# 9b92b546-74be-4628-8d99-b19e171beda1 | DUNLOP              | 2020-05-04 15:31:14.383516+03
# a0d0fe7e-6f44-4a79-9d4c-df951a0f8631 | KUMHO               | 2020-05-04 15:31:14.383516+03
# 96f3f504-bdda-4ac1-a4c1-f0b8e195f8af | Rosinca             | 2020-05-05 13:02:04.54421+03
# 390a13d0-ca93-4766-809b-a9994b964f38 | GAZ                 | 2020-05-05 16:04:11.295085+03
# 7da4bd51-67b6-4b82-ad48-fa2bfa37a503 | Russian Post        | 2020-05-05 15:24:17.056683+03
# 971a3522-47f3-4dd6-b075-5c60f3f60136 | АвтоВАЗ update test | 2020-05-05 16:04:11.133433+03
# 35ed13f5-aca4-4b33-b9e1-aad4023d6d77 | УАЗ update test     | 2020-05-05 16:04:11.251617+03
#(15 строк)
#
# - в результате обновлены две записи в составе "update test"
