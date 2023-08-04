This is simple-login project
## STARTUP

First, run command `go` to start application as root project
(go mod is compatible for this project) 

find out config as `config` folder

```
APP:
  NAME: ""
  PORT: 9000
JWT:
  SECRET: ""
  ISSUER: ""
  DOMAIN: ""
  EXPIRATION_TIME: 0
  REFRESH_EXPIRATION_TIME: 0
DATABASE:
  DSN: "host=xxx user=xxx password=xxx dbname=xxx port=5432 sslmode=disable TimeZone=Asia/Bangkok"
```

```
go run bootstrap/main.go
```

if you don't have sql/database. a simple way use it docker cli 

## TESTING
Test api, you can import postman collection as `postman` folder

for add mock user data
actually, 2 easier ways

`1. INSERT INTO sql script to user table`

`2. seed data pkg/db/default_client.go`

for test service work or not?

```
curl -X GET http://localhost:<port>/healthcheck
```

you might get OK with 200 status code

## COMPILE

run `Makefile compile command` 