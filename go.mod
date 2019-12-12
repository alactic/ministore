module github.com/alactic/ministore

go 1.13

require (
	github.com/alactic/ministore/sharedservice v0.0.0-20191212114200-4d44f404dd7f // indirect
	github.com/alactic/ministore/userservice v0.0.0-20191203090227-fc11f6c33b76
	github.com/alactic/ministore/userservices v0.0.0-20191211215435-bbb20428b6f2
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/satori/go.uuid v1.2.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.3
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
	google.golang.org/grpc v1.25.1
	gopkg.in/couchbase/gocb.v1 v1.6.4
)

replace github.com/alactic/ministore/proto/userdetail v0.0.0 => ./userdetail

replace github.com/satori/go.uuid v1.2.0 => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
