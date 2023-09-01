package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type configDb struct {
	SGDB     string
	Endereco string
	Usuario  string
	Senha    string
	Porta    string
	Banco    string
}

type CarregaConfigDb struct {
	arquivo      string
	Configuracao *configDb
	scanner      *bufio.Scanner
}

func (cf *CarregaConfigDb) CarregarConfig() {
	cf.arquivo = "config.json"
	cf.Configuracao = &configDb{}

	if FileExists(cf.arquivo) {
		cf.lerConfig()
		return
	} else {
		CreateFile(cf.arquivo)

		cf.scanner = bufio.NewScanner(os.Stdin)
		cf.Configuracao.SGDB = cf.lerValor("Informe o banco de dados 1. PostgreSQL \n 2. SQLITE")
		if cf.Configuracao.SGDB == "2" {
			cf.Configuracao.SGDB = "SQLITE"

		} else if cf.Configuracao.SGDB == "1" {
			cf.Configuracao.SGDB = "PostgreSQL"
			cf.Configuracao.Endereco = cf.lerValor("Informe o Endereço do servidor: ")
			cf.Configuracao.Porta = cf.lerValor("Informe a porta: ")
			cf.Configuracao.Usuario = cf.lerValor("Informe o usuário: ")
			cf.Configuracao.Senha = cf.lerValor("Informe a senha: ")
			cf.Configuracao.Banco = cf.lerValor("Informe o Banco de dados: ")
		} else {
			//valor Invalido
			panic("Banco de dados inválido")
		}

	}

	//Gravar configuração
	cf.salvarConfig()

	return
}

//Decoder transforma texto em struct
//Encoder transforma struct em texto

func (cf *CarregaConfigDb) salvarConfig() {
	file, _ := os.Create(cf.arquivo)
	defer file.Close() // Garante que o arquivo vai ser fechado quando o método terminar

	//Transforma a estrutura do objeto em texto json salvando um file
	json.NewEncoder(file).Encode(cf.Configuracao)
}

func (cf *CarregaConfigDb) lerConfig() {
	file, _ := os.Open(cf.arquivo)
	defer file.Close()

	//Lê o conteudo do arquivo e carrega no objeto configuracao
	json.NewDecoder(file).Decode(cf.Configuracao)
}

func (cf *CarregaConfigDb) lerValor(titulo string) string {
	fmt.Println(titulo)      // Imprime a informação que o user deve digitar;
	cf.scanner.Scan()        // Abre para digitação
	return cf.scanner.Text() // Retorna o texto digitado
}
