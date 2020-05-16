grpc_rest_mnf
=============

Микросервис для хранения и обработки запасных частей
----------------------------------------------------

### Технологический стек реализации микросервиса 
 go 1.13+
 grpc-rest-gateway
 protobuf 3
 postgres 11.7+
 docker

### Структура таблицы part
id, mnf_id, vendor_code, created_at, deleted_at

### Структура таблицы part_manufacturer
 id, name, created_at

### Реализованные методы сервиса PartService
mnf.v1.PartService@127.0.0.1:2338 show rpc
CreatePart
CreateParts
DeletePart
GetPart
ListPart
UpdatePart
UpdateParts

### Реализованные методы сервиса MnfService
mn f.v1.MnfService@127.0.0.1:2338 show rpc
CreateMnf
CreateMnfs
DeleteMnf
GetMnf
ListMnf
UpdateMnf
UpdateMnfs

### GRPC REST Gateway
см. описание в файле api.swagger.json


### Тестовые проверочные запросы к микросервису gRPC выполнены с применением клиента evans
https://evans.syfm.me/about

```
call_CreatePart.sh
```
RequestCreatePart.json

```
call_CreateParts.sh
```
RequestCreateParts.json

```
call_ListPart.sh
```
RequestListPart.json

```
call_GetPart.sh
```
RequestGetPart.json

```
call_CreateMnf.sh
```
RequestCreateMnf.json

```
call_CreateMnfs.sh
```
RequestCreateMnfs.json

```
call_GetMnf.sh
```
RequestGetMnf.json

```
call_ListMnf.sh
```
RequestListMnf.json

```
call_DeleteMnf.sh
```
RequestDeleteMnf.json

```
call_UpdateMnf.sh
```
RequestUpdateMnf.json

```
call_UpdateMnfs.sh
```
RequestUpdateMnfs.json

   
### В проекте использованы знания сообщества
Как  перейти на gRPC, сохранив REST
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

