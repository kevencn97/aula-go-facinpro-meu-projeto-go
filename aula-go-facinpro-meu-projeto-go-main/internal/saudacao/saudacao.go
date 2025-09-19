package saudacao

import "fmt"

func GerarSaudacao(nome string) string {
	return fmt.Sprintf("Olá, %s! Func presente no modulo saudação.", nome)
}
