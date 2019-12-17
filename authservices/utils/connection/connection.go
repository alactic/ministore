package connection

import (
	"fmt"
	"os"

	"gopkg.in/couchbase/gocb.v1"
)

func Connection() *gocb.Bucket {
	var bucket *gocb.Bucket
	var bucketName string

	cluster, _ := gocb.Connect(os.Getenv("COUCHBASE_HOST"))
    bucketName = os.Getenv("COUCHBASE_BUCKET")
 


	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: os.Getenv("COUCHBASE_USERNAME"),
		Password: os.Getenv("COUCHBASE_PASSWORD"),
	})
	// bucket, error = cluster.OpenBucket("default", "")
	bucket, _ = cluster.OpenBucket(bucketName, "")
	fmt.Println("bucketName :: ", bucketName)
	fmt.Println("host bucket:: ", bucket)


	return bucket
}
