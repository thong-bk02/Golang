package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("Không tìm thấy file .env.local, sẽ dùng biến môi trường hệ thống")
	}
	// Lấy URL từ file .env.local bạn đã tạo
	dsn := os.Getenv("YUGABYTE_URL")
	schemaName := os.Getenv("YUGABYTE_SCHEMA") // "core"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: schemaName + ".", // Tự động thêm "core." trước tên bảng
		},
	})

	if err != nil {
		log.Fatal("Kết nối YugabyteDB thất bại:", err)
	}
	return db
}
