package controllers

import (
	"auth/models"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"

	_ "log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitializeUserEndpoints(app *iris.Application, db *sqlx.DB, verifyMiddleWare context.Handler) {

	app.Get("/user_details", verifyMiddleWare, func(ctx iris.Context) {
		info := []models.UserInfo{}

		err := db.Select(&info, "SELECT * FROM user_details")

		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(info)
	})
}
