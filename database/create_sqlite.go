package database

import (

	"fmt"
	"os"
)

const DATABASE_NAME = "banco.db"



func CreateDatabase() {
	if FileNotExists(DATABASE_NAME) {
		err := CreateFile(DATABASE_NAME)

		if(err != nil){
			fmt.Println(err.Error())
			//Panic usado quando ocorre um erro que não deveria e compromete a aplicação, ele encera a aplicação caso dê erro
			panic(err)
		}
		

	}

}

//Cria um arquivo S.O 
func CreateFile(file string) error {
	newFile, err := os.Create(file)

	if err != nil {
		return err
	}
	newFile.Close()
	return nil
}

func FileNotExists(file string) bool {
	return !FileExists(file)
}

// verifiica se o arquivo existe
func FileExists(file string) bool {
	//os.Stat retorna o status do arquivo, ou erro caso ele não exista
	//colocamos _ sempre que não querermos utilizar um retorno da função
	_, err := os.Stat(file)

	// Se o erro que retornar for ARQUIVO NÃO EXISTE
	if (err != nil) && (os.IsNotExist(err)) {
		return false
	}
	return true
}
