package main

import (
	"aula4/api"
	"aula4/database"
)

func main() {
	bootStrap()

}

func bootStrap() {
	configuraBanco()
	database.Migration() //Cria as tabelas
	api.InicializeApi()
}

func configuraBanco() {
	config := database.CarregaConfigDb{}
	config.CarregarConfig()
	database.DbConfig = config.Configuracao
	if database.DbConfig.SGDB == "SQLITE" {
		database.CreateDatabase() //Cria o banco
	}
}
