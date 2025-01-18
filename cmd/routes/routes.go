package routes

import (
	handler "vote-system/internal/delivery"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(
	app fiber.Router,

) {
	app.Get("/", monitor.New())
}

func AdminRoutes(app fiber.Router, adminHandler handler.AdminHandler) {
	app.Post("/admin", adminHandler.RegisterAdmin)
	app.Post("/admin/login", adminHandler.LoginAdmin)
}

func UserRoutes(app fiber.Router, userHandler handler.UserHandler) {

}

func MovieRoutes(app fiber.Router, movieHandler handler.MovieHandler) {

}

func VoteRoutes(app fiber.Router, voteHandler handler.VotesHandler) {

}

func ViewershipRoutes(app fiber.Router, viewershipHandler handler.ViewershipHandler) {

}
