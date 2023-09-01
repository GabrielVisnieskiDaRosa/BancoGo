package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConfig *configDb

func GetDatabaseConnection() (*gorm.DB, error) {
	if DbConfig.SGDB == "SQLITE" {
		return gorm.Open(sqlite.Open(DATABASE_NAME))
	} else if DbConfig.SGDB == "PostgreSQL" {
		dbUrl := "postgres://" + DbConfig.Usuario + ":" + DbConfig.Senha + "@" + DbConfig.Endereco + ":" + DbConfig.Porta + "/" + DbConfig.Banco // acessar o postgres

		con, err := gorm.Open(postgres.Open(dbUrl))

		if err != nil {
			panic(err)
		}
		return con, nil

	} else {
		panic("SGBD não reconhecido")
	}
}
