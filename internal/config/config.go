package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	_ "strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv         string
	AppPort        string
	AppUrl         string
	DBHost         string
	DBPort         string
	DBName         string
	DBUser         string
	DBPass         string
	DBSSL          string
	JWTSecret      string
	JWTExpiresIn   string
	AdminEmail     string
	AdminPassword  string
	AdminFirstName string
	AdminLastName  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loding .end file: %v", err)
	}
	config := &Config{
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "3000"),
		AppUrl:       getEnv("APP_URL", "http://localhost:3000"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBSSL:        getEnv("DB_SSL", "disable"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),

		DBName:         getEnv("DB_NAME", ""),
		DBPass:         getEnv("DB_PASS", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		AdminEmail:     getEnv("ADMIN_EMAIL", ""),
		AdminPassword:  getEnv("ADMIN_PASS", ""),
		AdminFirstName: getEnv("ADMIN_FIRST_NAME", ""),
		AdminLastName:  getEnv("ADMIN_LAST_NAME", ""),
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil

}

func getEnv(key, defualtValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defualtValue
}

func validateConfig(config *Config) error {
	if config.AppEnv == "production" {
		if config.DBPass == "" {
			return fmt.Errorf("DB_PASS is required for production environment")
		}
		if config.JWTSecret == "" {
			return fmt.Errorf("JWTSecret is required for production environment")
		}
		if len(config.JWTSecret) < 32 {
			return fmt.Errorf("JWTSecret must be at least 32 characters long for production")
		}
		if config.DBSSL == "disable" {
			log.Println("Warning: SSL is disabled for database connection in production")
		}
		if config.AdminEmail == "" {
			return fmt.Errorf("ADMIN_EMAIL is required for production environment")
		}
		if config.AdminPassword == "" {
			return fmt.Errorf("ADMIN_PASS is required for production environment")
		}
		if config.AdminFirstName == "" {
			return fmt.Errorf("ADMIN_FRIST_NAME is required for production environment")
		}
		if config.AdminLastName == "" {
			return fmt.Errorf("ADMIN_LAST_NAME is required for production environment")
		}
	}

	if config.AdminEmail != "" && !isValidEmail(config.AdminEmail) {
		return errors.New("ADMIN_EMAIL must be a valid email address")
	}

	if config.DBName == "" {
		return fmt.Errorf("DB_Name is required")
	}

	return nil
}

func isValidEmail(email string) bool {
	if email == "" {
		return false
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	return emailRegex.MatchString(email)
}
