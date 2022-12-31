# Go Task ✅

![go-task logo](/docs/task-icon.png)


Service to manage task completion

-------

## Docker

### Running Kafka and Mysql
```bash
docker-compose -f docker-compose-dev.yaml up -d --build
```

### How to find ips locally
```bash
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql.go-task
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' kafka.go-task
```

## API

### Create new task
[POST] /task
```bash
curl --location --request POST 'http://127.0.0.1:8080/task' \
--header 'userId: 7ee7698f-7467-4c3c-932b-1a1574ae8f7b' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Title": "Title Task",
    "Description": "desc task",
    "ManagerUserId": "484fb924-92e4-4d30-9897-22dcf7ce9fec"
}'
```

### List tasks
[GET] /task
```bash
curl --location --request GET 'http://127.0.0.1:8080/task' \
--header 'userId: 7ee7698f-7467-4c3c-932b-1a1574ae8f7b'
```

### Complete task
[PATCH] /task/conclude/93278ef1-04b1-4133-aa23-57b6e1bf7d2a
```bash
curl --location --request PATCH 'http://127.0.0.1:8080/task/conclude/93278ef1-04b1-4133-aa23-57b6e1bf7d2a' \
--header 'userId: 7ee7698f-7467-4c3c-932b-1a1574ae8f7b'
```