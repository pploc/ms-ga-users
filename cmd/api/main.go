package main

import (
	"log"

	"go.uber.org/zap"

	"ms-ga-user/internal/api/handler"
	"ms-ga-user/internal/api/router"
	"ms-ga-user/internal/infrastructure/messaging"
	"ms-ga-user/internal/infrastructure/persistence/gorm/repository"
	"ms-ga-user/internal/service"
	"ms-ga-user/pkg/config"
	"ms-ga-user/pkg/database"
	"ms-ga-user/pkg/utils"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	utils.InitLogger(cfg.AppEnv)
	defer utils.SyncLogger()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		utils.Log.Fatal("Database connection failed", zap.Error(err))
	}

	kafkaProducer, err := messaging.NewKafkaProducer(cfg)
	if err != nil {
		utils.Log.Fatal("Kafka producer initialization failed", zap.Error(err))
	}
	defer kafkaProducer.Close()

	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	userService := service.NewUserService(userRepo, kafkaProducer)
	profileService := service.NewProfileService(profileRepo)

	userHandler := handler.NewUserHandler(userService)
	profileHandler := handler.NewProfileHandler(profileService)

	r := router.SetupRouter(userHandler, profileHandler)

	utils.Log.Info("Starting ms-ga-user on port " + cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		utils.Log.Fatal("Server forced to shutdown", zap.Error(err))
	}
}
