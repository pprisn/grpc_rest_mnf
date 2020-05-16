# Развертывание контейнера-прицепа topz для мониторинга работы приложения
 
Определим хэш сумму для контейнера приложения командой
```
     docker ps 
```
в моем случае 
```
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                                                                              NAMES
c80d17f5f638        grpcrestmnf_app     "/bin/sh -c ./${APP_…"   37 minutes ago      Up 37 minutes       0.0.0.0:2338->2338/tcp, 0.0.0.0:5775->5775/tcp, 0.0.0.0:8080-8081->8080-8081/tcp   grpc_mnf
ad0d6cb13f46        postgres:11.7       "docker-entrypoint.s…"   37 minutes ago      Up 37 minutes       0.0.0.0:5436->5432/tcp                                                             grpcrestmnf_db_1
```
это ID = c80d17f5f638

Запишем команду для запуска контейнера прицепа в файл run_topz.sh
```
#!/bin/sh                                                                                                                                                                                                                      
  2 APP_ID='c80d17f5f638'                                                                                                   
  3 docker run --pid=container:${APP_ID} \                                                                                  
  4         -p 8083:8083 brendanburns/topz:db0fa58 /server -addr 0.0.0.0:8083 
```
Установим атрибуты разрешения на запуск.
chmod 755 run_topz.sh

После запуска контейнера-прицепа командой
./run_topz.sh
 
 Получить срез информации о процессах, работающих в  контейнере приложения и используемых ими ресурсах
 можно будет, обратившись по  адресу http://localhost:8083/topz

В моем случае отображается следующая информация
```
20 9.905868592174695 0.03887358  /server -addr 0.0.0.0:8083
1  0                 0.000591483 /bin/sh -c ./${APP_NAME}
6  0                 0.12378426  ./grpc_mnf
```

 Из исходного кода topz можно определить назначение полей вывода 
 это :
 > proc.PID, 
 > proc.CPUPercent, 
 > proc.MemoryPercent, 
 > proc.Command

Исходный код  приложения topz посмотреть можно здесь
```
https://github.com/brendandburns/topz/blob/master/pkg/topz/topz.go
```

