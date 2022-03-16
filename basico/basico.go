/*
Tipos de variáveis ​​Go

package main

import "fmt"

func main() {
	a := 1    // var a int
	b := 3.14 // var b float
	c := "hi" // var c string
	d := true // var d bool
	fmt.Println(a, b, c, d)

	e := []int{1, 2, 3} // slice
	e = append(e, 4)
	fmt.Println(e, len(e), e[0], e[1:3], e[1:], e[:2])

	f := make(map[string]int) // map
	f["one"] = 1
	f["two"] = 2
	fmt.Println(f, len(f), f["one"], f["three"])
}
*/

/*
Declaração de variável sem valor inicial

package main

import (
	"fmt"
)

func main() {
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
*/

/*
Atribuição de valor após declaração

package main
import ("fmt")

func main() {
  var student1 string
  student1 = "John"
  fmt.Println(student1)
}
*/

/*
Data Types

package main

import (
	"fmt"
)

var (
	nome string
	n1   int
	n2   int
)

func main() {
	nome = "Gustavo"
	fmt.Println("Hello", nome, "!")
}
*/

/*
 Condicional

package main

import "fmt"

func main() {
	// if, else
	a := 5
	if a > 3 {
		a = a - 3
	} else if a == 3 {
		a = 0
	} else {
		a = a + 3
	}
	fmt.Println(a)

	// switch, case
	b := "NO"
	switch b {
	case "YES":
		b = "Y"
	case "NO":
		b = "N"
	default:
		b = "X"
	}
	fmt.Println(b)
}
*/

/*
 Funções
*/
package main

import "fmt"

func main() {
	fmt.Println(threeTimes("Thank You"))
}

func threeTimes(msg string) (tMsg string) {
	tMsg = msg + ", " + msg + ", " + msg
	return
}
