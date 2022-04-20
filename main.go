package main

import "fmt"

func main() {
	// Uma função pode retornar mais de um valor.
	// var ok bool
	// var err error
	//a construção := cria a variavel caso ela não exista, já defindo seu tipo.
	ok, err := say("Olá Jaque")
	if err != nil {
		panic(err.Error())
	}
	switch ok {
	case true:
		fmt.Println("Bom dia")
	default:
		fmt.Println("Falha na comunicação!")
	}
}

//Função devem declarar o tipo de cada variável que recebe ou que retorna.
func say(what string) (bool, error) {
	if what == "" {
		return false, fmt.Errorf("Empty string")
	}
	fmt.Println(what)
	return true, nil
}
