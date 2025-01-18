package main

import (
	"log"
	"vote-system/cmd/routes"
	"vote-system/database"
	adminRepo "vote-system/internal/app/repository/admin"
	movieRepo "vote-system/internal/app/repository/movie"
	userRepo "vote-system/internal/app/repository/user"
	viewershipRepo "vote-system/internal/app/repository/viewership"
	voteRepo "vote-system/internal/app/repository/votes"
	"vote-system/internal/app/usecase/admin"
	movieSvc "vote-system/internal/app/usecase/movie"
	userSvc "vote-system/internal/app/usecase/user"
	viewershipSvc "vote-system/internal/app/usecase/viewership"
	votesSvc "vote-system/internal/app/usecase/votes"
	handler "vote-system/internal/delivery"
	"vote-system/internal/entity"
	"vote-system/pkg/config"
	"vote-system/pkg/identifier"
	"vote-system/pkg/validator"

	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.SetupEnvFile()

	db := database.InitPostgres(envConfig)

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())

	adminRepository := adminRepo.NewAdminRepository(db)
	userRepository := userRepo.NewUserRepository(db)
	movieRepository := movieRepo.NewMovieRepository(db)
	votesRepository := voteRepo.NewVotesRepository(db)
	viewershipRepository := viewershipRepo.NewViewershipRepository(db)

	adminService := admin.NewService(adminRepository, validator, identifier)
	userService := userSvc.NewService(userRepository, validator, identifier)
	movieService := movieSvc.NewService(movieRepository, validator, identifier)
	votesService := votesSvc.NewService(votesRepository, validator, identifier)
	viewershipService := viewershipSvc.NewService(viewershipRepository, validator, identifier)

	adminHandler := handler.NewAdminHandler(adminService, movieService)
	userHandler := handler.NewUserHandler(userService)
	movieHandler := handler.NewMovieHandler(movieService)
	votesHandler := handler.NewVotesHandler(votesService)
	viewershipHandler := handler.NewViewershipHandler(viewershipService)

	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024, // 50 MB
	})
	app.Static("/uploads", "../uploads")

	routes.SetupRoutes(app)
	routes.AdminRoutes(app, adminHandler)
	routes.UserRoutes(app, userHandler)
	routes.MovieRoutes(app, movieHandler)
	routes.VoteRoutes(app, votesHandler)
	routes.ViewershipRoutes(app, viewershipHandler)

	log.Print("Database Migration is:", config.SetupEnvFile().DatabaseMigration)
	if config.SetupEnvFile().DatabaseMigration {
		if err := db.DB.AutoMigrate(
			&entity.Admin{},
			&entity.User{},
			&entity.Movie{},
			&entity.Vote{},
			&entity.Viewership{},
		); err != nil {
			log.Fatalf("AutoMigrate failed: %v", err)
		}

	}

	if err := app.Listen(":5004"); err != nil {
		log.Fatalf("listen: %s", err)
	}

}
