package routes

import (
	"os"
	"vote-system/cmd/middleware"
	handler "vote-system/internal/delivery"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(
	app fiber.Router,

) {
	app.Get("/", monitor.New())

	app.Get("/uploads/:filename", func(c *fiber.Ctx) error {
		// Extract the filename from the URL
		filename := c.Params("filename")
		filepath := "../uploads/" + filename

		// Check if the file exists
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).SendString("File not found")
		}

		// Serve the file
		return c.SendFile(filepath)
	})
}

func AdminRoutes(app fiber.Router, adminHandler handler.AdminHandler) {
	app.Post("/admin", adminHandler.RegisterAdmin)
	app.Post("/admin/login", adminHandler.LoginAdmin)
	app.Post("/admin/upload", middleware.AuthUser, adminHandler.UploadMovie)
	app.Post("/admin/insert", middleware.AuthUser, adminHandler.InsertMovie)
	app.Put("/admin/update", middleware.AuthUser, adminHandler.UpdateMovie)
}

func UserRoutes(app fiber.Router, userHandler handler.UserHandler) {
	app.Post("/user", userHandler.RegisterUser)
	app.Post("/user/login", userHandler.LoginUser)
}

func MovieRoutes(app fiber.Router, movieHandler handler.MovieHandler) {

}

func VoteRoutes(app fiber.Router, voteHandler handler.VotesHandler) {

}

func ViewershipRoutes(app fiber.Router, viewershipHandler handler.ViewershipHandler) {

}
