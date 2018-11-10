# gin-lnp-api - Local Number Portability

## Goals

Build or loy out the template that ready to implement REST API & gRPC services.

* API-readiness: must support HTTP/HTTPs protocol & JSON content type.
* Testability:
* Deployability:
* Adaptability:

## Libraries

* Routing framework: [gin gonic](https://github.com/gin-gonic/gin)
* Database: [go-cql](https://github.com/gocql/gocql)
* Data validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* JWT token: [jwt.io](https://github.com/gbrlsnchs/jwt)
* Logging: [logrus](https://github.com/sirupsen/logrus)
* Configuration: [viper](https://github.com/spf13/viper)
* Wire: [wire](https://github.com/google/go-cloud/tree/master/wire)
* Dependency management: [dep](https://github.com/golang/dep)
* Testing: [testify](https://github.com/stretchr/testify)


## API Endpoints:

| Path                               | Method | Description                        | Response          |Done|
|------------------------------------|--------|------------------------------------|-------------------|----|
| /v1/lnp/dids/:did                  | GET    | get phone by phone #               | LRN json object   | Y  |
| /v1/lnp/dids?did_list=#,#,#        | GET    | get phones by list of phone #      | LRN list          | Y  |
| /v1/lnp/ocns/:ocn                  | GET    | get ocn info by ocn code           | OCN json object   | Y  |
| /v1/lnp/lergs/:npa/:nxx/:block_id  | GET    | get lerg info from lerg database   | LERG json object  | Y  |



## [gRPC](https://grpc.io/)

Install gRPC
Use the following command to install gRPC.
go get -u google.golang.org/grpc

* Generate gPRC code

protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworl

