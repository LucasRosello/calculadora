package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func (Calc) Operate(entrada string, operador string) int {

	numeros := strings.Split(entrada, operador)
	numero0 := parsear(numeros[0])
	numero1 := parsear(numeros[1])

	switch operador {
	case "+":
		fmt.Println(numero0 + numero1)
	case "-":
		fmt.Println(numero0 - numero1)
	case "*":
		fmt.Println(numero0 * numero1)
	case "/":
		fmt.Println(numero0 / numero1)
	default:
		fmt.Println("che, ¿Que carajo pusiste?")
	}

	return numero0
}

func parsear(entrada string) int {
	numero, _ := strconv.Atoi(entrada)
	return numero
}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
