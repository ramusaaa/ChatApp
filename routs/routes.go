package routs

import (
	"ChatApp/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	//app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/api/aboutme", controllers.AllAboutme)
	app.Post("/api/aboutme", controllers.CreateAboutme)
	app.Get("/api/aboutme/:id", controllers.GetAboutme)
	app.Put("/api/aboutme/:id", controllers.UpdateAboutme)
	app.Delete("/api/aboutme/:id", controllers.DeleteAboutme)

	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/uploads", "./uploads")

}
