module github.com/alactic/ministore/userservice

go 1.13

require (
	github.com/alactic/ministore v0.0.0-20191203074322-50417243e2bb
	github.com/alactic/mygoproject v0.0.0-20191123211049-2df91565b1d2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.3 // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	gopkg.in/couchbase/gocb.v1 v1.6.4
	gopkg.in/couchbase/gocbcore.v7 v7.1.15 // indirect
	gopkg.in/couchbaselabs/gocbconnstr.v1 v1.0.4 // indirect
	gopkg.in/couchbaselabs/gojcbmock.v1 v1.0.3 // indirect
	gopkg.in/couchbaselabs/jsonx.v1 v1.0.0 // indirect
)

replace github.com/satori/go.uuid v1.2.0 => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
