# coding test

## prerequisites

install golang

```
$ sudo apt-get update
$ sudo apt-get install golang-go
```

install dep
```
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

run dep ensure
```
$ dep ensure
```

to run project:
```
$ go run main.go
```

Server will start on
```
http://localhost:8080
```

## Available queries
* GET http://localhost:8080/plans
* POST http://localhost:8080/subscribe

### plans

#### get list of plans
```
GET http://localhost:8080/plans
```

```
curl -X GET http://localhost:8080/plans
```

example response:
```json
[
    {
        "id": 1,
        "name": "1 - month",
        "price": 9.95
    },
    {
        "id": 2,
        "name": "3 - months",
        "price": 19.95
    },
    {
        "id": 3,
        "name": "1 - year",
        "price": 49.95
    }
]
```

####  get list of plans for particular currency

```
GET http://localhost:8080/plans?currency=USD
```

```curl
curl -X GET 'http://localhost:8080/plans?currency=USD'
```

example response:
```json
[
    {
        "id": 1,
        "name": "1 - month",
        "price": 11.26719
    },
    {
        "id": 2,
        "name": "3 - months",
        "price": 22.591002
    },
    {
        "id": 3,
        "name": "1 - year",
        "price": 56.56243
    }
]
```

### subscribe

```
POST http://localhost:8080/subscribe
```

#### valid query

```
curl -X POST http://localhost:8080/subscribe  -d '{"planId":1,"name":"Bobby","email":"example@example.com","phone":"+ 49 30 738788"}'
```

example response:
```json
{
    "planId": 1,
    "name": "Bobby",
    "email": "example@example.com",
    "phone": "+ 49 30 738788"
}
```

#### invalid query

```
curl -X POST http://localhost:8080/subscribe  -d '{}'
```

example response:
```json
{
    "email": "is empty",
    "name": "is empty",
    "phone": "is empty",
    "planID": "is empty"
}
```

#### query with invalid phone
```
curl -X POST http://localhost:8080/subscribe -d '{"planId":1,"name":"Bobby","email":"example@example.com","phone":"+380500000000"}'
```

example response:
```
{
    "phone": "is invalid"
}
```
