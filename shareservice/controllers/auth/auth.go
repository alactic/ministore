package auth

import (
	jwtFile "github.com/alactic/ministore/shareservice/utils/jwt"
)


func Authorization(token string) int64 {
	if len(token) == 0 {
		return 400
	}

	e := jwtFile.DecodeJWT(token)
    if e != nil {
		if e != nil {
			return 400
		}
	}
	return 200
}

