# Get started

~~~bash
$ docker-compose up -d
$ docker exec -it go-server /bin/bash
$ go run main.go
# create peoples
$ curl -i -X POST http://localhost:8080/peoples -d '{ "name": "Takumi", "age": 20, "gender": "man"}' -H "Content-Type: application/json"
# get peoples
$ curl http://localhost:8080/peoples
~~~
