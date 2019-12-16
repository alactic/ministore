package user

import (
	// "github.com/alactic/ministore/shareservice/models/userm"
	"github.com/alactic/ministore/shareservice/utils/connection"
	"gopkg.in/couchbase/gocb.v1"
)

var bucket *gocb.Bucket = connection.Connection()

// Getting userdetails
func UserDetails(email string) map[string]string {
	// var users []userm.UpdateUser
	// query := gocb.NewN1qlQuery("SELECT " + bucket.Name() + ".* FROM " + bucket.Name() + " WHERE `email` = $1")
	// rows, err := bucket.ExecuteN1qlQuery(query, []interface{}{email})

	// if err != nil {
	// 	// ctx.JSON(http.StatusBadRequest, err.Error())
	// 	//	return
	// }

	// var row userm.UpdateUser
	// for rows.Next(&row) {
	// 	users = append(users, row)
	// }
	// details := make(map[string]string)
	// details["firstname"] = users[0].Firstname
	// details["lastname"] = users[0].Lastname
	// details["email"] = users[0].Email
	// details["id"] = users[0].ID
	// details["type"] = users[0].Type


	details := make(map[string]string)
	details["firstname"] = "elvis"
	details["lastname"] = "echezona"
	details["email"] = email
	details["id"] = "23404"
	details["type"] = "admin"


	return details
}
