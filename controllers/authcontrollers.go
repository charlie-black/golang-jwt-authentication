package controllers

import (
	"auth/models"

	appfunctions "auth/functions"

	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"

	_ "github.com/lib/pq"
)

func InitializeAuthEndpoints(signer *jwt.Signer, app *iris.Application, db *sqlx.DB){
	
	app.Post("/sign_up", func(ctx iris.Context){
		var signup_params models.SignUpParams
		err := ctx.ReadJSON(&signup_params)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error":err.Error()})
			return
		}
		hashedPassword := appfunctions.HashPassword(signup_params.Password)

	claims := &models.UserClaims {
		Username: signup_params.Email,}
		token, err := signer.Sign(claims)

		_, err = db.NamedExec(`INSERT INTO users (email, password, phone, name)
        VALUES (:email, :password, :phone, :name)`, map[string]interface{}{
			"email": signup_params.Email, "password": hashedPassword, "phone":signup_params.Phone, "name":signup_params.Name,
		})

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return

		}
		_, err = db.NamedExec(`INSERT INTO user_details (email, phone, name)
        VALUES (:email, :phone, :name)`, map[string]interface{}{
			"email": signup_params.Email, "phone":signup_params.Phone, "name":signup_params.Name,
		})

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return

		}
		ctx.JSON(iris.Map{"message": "User Created Successfully", "token": string(token)})
	})

	app.Post("/login", func(ctx iris.Context) {
		var login_params models.LoginParams
		err := ctx.ReadJSON(&login_params)

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}

		hashedPassword := appfunctions.HashPassword(login_params.Password)

		var result models.LoginParams
		userCheck := db.Get(&result, "SELECT email,password FROM users WHERE email=$1", login_params.Email)

		if userCheck != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"message": "user does not exist"})
			return
		}

		if result.Password != hashedPassword {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"message": "invalid password"})
			return
		}

		claims := &models.UserClaims{
			Username: login_params.Email,
			
		}
		token, err := signer.Sign(claims)

		ctx.JSON(iris.Map{"message": "Logged in Successfully", "token": string(token)})

	})

	

	

	}
	

	

	


