grpc_rest_mnf
=============

Микросервис для хранения и обработки запасных частей
----------------------------------------------------

### Технологический стек реализации микросервиса 
> go 1.13+
> grpc-rest-gateway
> protobuf 3
> postgres 9.6+

### Структура таблицы part
id, mnf_id, vendor_code, created_at, deleted_at

### Структура таблицы part_manufacturer
id, name, created_at

### Реализованные методы сервиса PartService
mnf.v1.PartService@127.0.0.1:2338> show rpc
+-------------+--------------------------------+---------------------------+----------------------------+
|    NAME     |      FULLY-QUALIFIED NAME      |       REQUEST TYPE        |       RESPONSE TYPE        |
+-------------+--------------------------------+---------------------------+----------------------------+
| CreatePart  | mnf.v1.PartService.CreatePart  | mnf.v1.CreatePartRequest  | mnf.v1.CreatePartResponse  |
| CreatePart  | mnf.v1.PartService.CreatePart  | mnf.v1.CreatePartRequest  | mnf.v1.CreatePartResponse  |
| CreateParts | mnf.v1.PartService.CreateParts | mnf.v1.CreatePartsRequest | mnf.v1.CreatePartsResponse |
| CreateParts | mnf.v1.PartService.CreateParts | mnf.v1.CreatePartsRequest | mnf.v1.CreatePartsResponse |
| DeletePart  | mnf.v1.PartService.DeletePart  | mnf.v1.DeletePartRequest  | mnf.v1.DeletePartResponse  |
| DeletePart  | mnf.v1.PartService.DeletePart  | mnf.v1.DeletePartRequest  | mnf.v1.DeletePartResponse  |
| GetPart     | mnf.v1.PartService.GetPart     | mnf.v1.GetPartRequest     | mnf.v1.GetPartResponse     |
| GetPart     | mnf.v1.PartService.GetPart     | mnf.v1.GetPartRequest     | mnf.v1.GetPartResponse     |
| ListPart    | mnf.v1.PartService.ListPart    | mnf.v1.ListPartRequest    | mnf.v1.ListPartResponse    |
| ListPart    | mnf.v1.PartService.ListPart    | mnf.v1.ListPartRequest    | mnf.v1.ListPartResponse    |
| UpdatePart  | mnf.v1.PartService.UpdatePart  | mnf.v1.UpdatePartRequest  | mnf.v1.UpdatePartResponse  |
| UpdatePart  | mnf.v1.PartService.UpdatePart  | mnf.v1.UpdatePartRequest  | mnf.v1.UpdatePartResponse  |
| UpdateParts | mnf.v1.PartService.UpdateParts | mnf.v1.UpdatePartsRequest | mnf.v1.UpdatePartsResponse |
| UpdateParts | mnf.v1.PartService.UpdateParts | mnf.v1.UpdatePartsRequest | mnf.v1.UpdatePartsResponse |
+-------------+--------------------------------+---------------------------+----------------------------+

###Реализованные методы сервиса MnfService
mnf.v1.MnfService@127.0.0.1:2338> show rpc
+------------+------------------------------+--------------------------+---------------------------+
|    NAME    |     FULLY-QUALIFIED NAME     |       REQUEST TYPE       |       RESPONSE TYPE       |
+------------+------------------------------+--------------------------+---------------------------+
| CreateMnf  | mnf.v1.MnfService.CreateMnf  | mnf.v1.CreateMnfRequest  | mnf.v1.CreateMnfResponse  |
| CreateMnf  | mnf.v1.MnfService.CreateMnf  | mnf.v1.CreateMnfRequest  | mnf.v1.CreateMnfResponse  |
| CreateMnfs | mnf.v1.MnfService.CreateMnfs | mnf.v1.CreateMnfsRequest | mnf.v1.CreateMnfsResponse |
| CreateMnfs | mnf.v1.MnfService.CreateMnfs | mnf.v1.CreateMnfsRequest | mnf.v1.CreateMnfsResponse |
| DeleteMnf  | mnf.v1.MnfService.DeleteMnf  | mnf.v1.DeleteMnfRequest  | mnf.v1.DeleteMnfResponse  |
| DeleteMnf  | mnf.v1.MnfService.DeleteMnf  | mnf.v1.DeleteMnfRequest  | mnf.v1.DeleteMnfResponse  |
| GetMnf     | mnf.v1.MnfService.GetMnf     | mnf.v1.GetMnfRequest     | mnf.v1.GetMnfResponse     |
| GetMnf     | mnf.v1.MnfService.GetMnf     | mnf.v1.GetMnfRequest     | mnf.v1.GetMnfResponse     |
| ListMnf    | mnf.v1.MnfService.ListMnf    | mnf.v1.ListMnfRequest    | mnf.v1.ListMnfResponse    |
| ListMnf    | mnf.v1.MnfService.ListMnf    | mnf.v1.ListMnfRequest    | mnf.v1.ListMnfResponse    |
| UpdateMnf  | mnf.v1.MnfService.UpdateMnf  | mnf.v1.UpdateMnfRequest  | mnf.v1.UpdateMnfResponse  |
| UpdateMnf  | mnf.v1.MnfService.UpdateMnf  | mnf.v1.UpdateMnfRequest  | mnf.v1.UpdateMnfResponse  |
| UpdateMnfs | mnf.v1.MnfService.UpdateMnfs | mnf.v1.UpdateMnfsRequest | mnf.v1.UpdateMnfsResponse |
| UpdateMnfs | mnf.v1.MnfService.UpdateMnfs | mnf.v1.UpdateMnfsRequest | mnf.v1.UpdateMnfsResponse |
+------------+------------------------------+--------------------------+---------------------------+

###GRPC REST Gateway
см. описание в файле api.swagger.json

###Тестовые проверочные запрсы к микросервису gRPC выполнены с применением клиента evans
https://evans.syfm.me/about

call_CreatePart.sh
RequestCreatePart.json

call_CreateParts.sh
RequestCreateParts.json

call_ListPart.sh
RequestListPart.json

call_GetPart.sh
RequestGetPart.json

call_CreateMnf.sh
RequestCreateMnf.json

call_CreateMnfs.sh
RequestCreateMnfs.json

call_GetMnf.sh
RequestGetMnf.json

call_ListMnf.sh
RequestListMnf.json

call_DeleteMnf.sh
RequestDeleteMnf.json

call_UpdateMnf.sh
RequestUpdateMnf.json

call_UpdateMnfs.sh
RequestUpdateMnfs.json

 
###В проекте использованы знания сообщества
Как перейти на gRPC, сохранив REST
https://habr.com/ru/post/337716/

Отличный проект для изучения материала
https://github.com/abronan/todo-grpc

Как отлаживать сервер gRPC
https://medium.com/@EdgePress/how-to-interact-with-and-debug-a-grpc-server-c4bc30ddeb0b

Отличный пример для шаблона http-rest-api
https://github.com/gopherschool/http-rest-api/

Инструментарий мониторинга gRPC сервиса
Prometheus Go client library
https://github.com/prometheus/client_golang

Библиотека инструментов, которая реализует OpenTracing Go Tracer для Jaeger (https://jaegertracing.io)
https://github.com/uber/jaeger-client-go

