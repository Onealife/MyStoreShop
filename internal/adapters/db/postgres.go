package db

import (
	"fmt"
	"log"

	"github.com/Onealife/MyStoreShop/config"
	"github.com/Onealife/MyStoreShop/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(config *config.Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort, config.DBSSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("ไม่สามารถเชื่อมต่อฐานข้อมูลได้: %v", err)
	}
	log.Println("เชื่อมต่อฐานข้อมูลสำเร็จ")

	return db
}

func migrateDatabase(db *gorm.DB) {
	//
	err := db.AutoMigrate(
		&domain.Role{},
		&domain.Permission{},
		&domain.User{},
		&domain.Category{},
		&domain.Product{},
		&domain.ProductImage{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Transaction{},
	)
	if err != nil {
		log.Fatalf("ไม่สามารถสร้างฐานข้อมูลได้: %v", err)
	}
	log.Println("สร้างฐานข้อมูลสำเร็จ")
}
