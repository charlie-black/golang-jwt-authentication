package main

import (
	authcontroller "auth/controllers"
	usercontroller "auth/controllers"
	"auth/models"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"

	_ "github.com/lib/pq"
)

func main(){
	var SECRET = []byte("@#$%$#%$@#&*$")
	signer := jwt.NewSigner(jwt.HS256, SECRET, 90*time.Minute)
	verifier := jwt.NewVerifier(jwt.HS256, SECRET)

	verifyMiddleWare := verifier.Verify(func() interface{} {
		return new(models.UserClaims)
	})
	//connect to the database

	db, err := sqlx.Connect("postgres", "user=piccasso dbname=authDB sslmode=disable")

	if err != nil {
		log.Fatalln(err)

	}
	println("Connected to the database")

}