package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Onealife/Nutchapholshop/internal/adapters/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort, config.DBSSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	log.Println("Database conected successfully")

	if shouldRunMigration() {
		RunMigration(db)

		if err := SeedAdminUser(db, config); err != nil {
			log.Panicf("Admin seeding failed: %v", err)
		}
	} else {
		autoMigrate := os.Getenv("AUTO_MIGRATE")
		appEnv := os.Getenv("APP_ENV")

		if autoMigrate == "false" {
			log.Printf("Skipping database migration (AUTO_MIGRATE)=false")
		} else if appEnv == "production" && autoMigrate != "true" {
			log.Printf("Skipping database migration (production enviroment, set AUTO_MIGRATE=true to enable")
		} else {
			log.Printf("Skipping database migration (set AUTO_MIGRATE=true to enable)")
		}

		if err := SeedAdminUser(db, config); err != nil {
			log.Panicf("Admin seeding failed: %v", err)
		}
	}

	return db
}

func shouldRunMigration() bool {
	if os.Getenv("AUTO_MIGRATE") == "false" {
		return false
	}
	if os.Getenv("AUTO_MIGRATE") == "true" {
		return true
	}
	if os.Getenv("APP_ENV") == "development" {
		return true
	}
	return false
}

func RunMigration(db *gorm.DB) {

	fmt.Println("Starting database migration....")

	err := db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Error to migration database:", err)
	}

	log.Println("Database migration migration successfuly")
}

func RunMigrationManual(config *Config) error {
	db := SetupDatabase(config)

	log.Println("Running manual migration....")

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("migration failed %v", err)
	}

	log.Println("Manual migration completed successfully")

	return nil
}
