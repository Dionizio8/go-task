# Go Task

![go-task logo](/docs/task-icon.png)

Service to manage tasks

-------

## Initial settings local
The project has the following dependencies :

* Docker ***(recommended version >= 20)***
* Docker Compose
* Go 1.19


*In the terminal, access the project folder and execute the following commands:*

### Running Kafka and Mysql local
```bash
docker-compose -f docker-compose-dev.yaml up -d --build
```

### How to find ips locally [IP_MYSQL, IP_KAFKA]
```bash
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql.go-task
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' kafka.go-task
```

### Copying configuration file
```bash
cp example.env .env
```

### Installing project dependencies
```bash
go mod download
```

### Start project
#### API
```bash
go run ./cmd/api/main.go
```
#### Worker
```bash
go run ./cmd/worker/main.go
```

---- 

## API

[Postman Collection](./docs/postman_collection.json) ***[file to import the requests in the Postman tool]***

### Create new task
**[POST] /task**
#### Parameters
* **userId** : UUID registered in the database. ***[Need to replace as it is created dynamically]***
* **managerUserId** : UUID registered in the database. ***[Need to replace as it is created dynamically]**
* **title**: Task title
* **description**: Task description
```bash
curl --location --request POST 'http://127.0.0.1:8080/task' \
--header 'userId: 7ee7698f-7467-4c3c-932b-1a1574ae8f7b' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Title Task",
    "description": "desc task",
    "managerUserId": "484fb924-92e4-4d30-9897-22dcf7ce9fec"
}'
```

### List tasks
**[GET] /task**
#### Parameters
* **userId** : UUID registered in the database. ***[Need to replace as it is created dynamically]***
```bash
curl --location --request GET 'http://127.0.0.1:8080/task' \
--header 'userId: 7ee7698f-7467-4c3c-932b-1a1574ae8f7b'
```

### Complete task
**[PATCH] /task/conclude/<ID_TASK>**
#### Parameters
* **userId** : UUID registered in the database. ***[Need to replace as it is created dynamically]***
* **<ID_TASK>** : UUID registered in the database. ***[Need to replace as it is created dynamically]***
```bash
curl --location --request PATCH 'http://127.0.0.1:8080/task/conclude/93278ef1-04b1-4133-aa23-57b6e1bf7d2a' \
--header 'userId: 7ee7698f-7467-4c3c-932b-1a1574ae8f7b'
```

----
