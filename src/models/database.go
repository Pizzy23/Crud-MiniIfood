package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	databaseURL := os.Getenv("DB")
	if databaseURL == "" {
		log.Fatal("Variável de ambiente DB não configurada")
	}

	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info), // Adicionando um logger para exibir logs de consulta
	})

	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: ", err)
	}

	Migrate(db)

	DB = db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&Food{})
	db.AutoMigrate(&Login{})
	db.AutoMigrate(&User{})
}
