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

package main

import "fmt"

func main() {
	fmt.Println(threeTimes("Thank You"))
}

func threeTimes(msg string) (tMsg string) {
	tMsg = msg + ", " + msg + ", " + msg
	return
}
*/

/*
 Chamar funções

package main

import (
	"fmt"
)

func minhaMensagem() {
	fmt.Println("Vai executar!")
}

func main() {
	minhaMensagem() // chama a func
}
*/

/*
	Array

package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
*/

/*
	Map com Make()

package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// a is no longer empty
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}
*/

/*
	Acessando elementos do map

package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Printf(a["brand"])
}
*/

/*
	Acessar membros da estrutura
*/
package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)
}
