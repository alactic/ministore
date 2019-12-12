module github.com/alactic/ministore/sharedservice

go 1.13

require (
	github.com/alactic/ministore v0.0.0-20191211215435-bbb20428b6f2 // indirect
	github.com/alactic/ministore/userservice v0.0.0-20191203090227-fc11f6c33b76
	github.com/alactic/ministore/userservices v0.0.0-20191211215435-bbb20428b6f2
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gin-gonic/gin v1.5.0
	github.com/satori/go.uuid v1.2.0
	github.com/swaggo/swag v1.6.3
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect
	gopkg.in/couchbase/gocb.v1 v1.6.4
)

replace github.com/alactic/ministore v0.0.0 => github.com/alactic/ministore v0.0.0-20191211215435-bbb20428b6f2
