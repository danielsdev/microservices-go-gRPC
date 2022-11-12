package database

import (
	"log"
	"os"

	"github.com/danielsdev/microservices-go-gRPC/manager/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	HOST := os.Getenv("DB_HOST")
	PASSWORD := os.Getenv("DB_PASS")
	USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := "host=" + HOST +
		" user=" + USER +
		" password=" + PASSWORD +
		" dbname=" + DB_NAME +
		" port=" + DB_PORT +
		" sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	//Cria a tabela com base na struct de Aluno
	DB.AutoMigrate(&models.Aluno{})
}
