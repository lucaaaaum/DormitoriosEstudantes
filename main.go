package main

import "fmt"

func main() {
	listaDeEstudantes := make([]*estudante, 0)
	for i := 0; i < 100; i++ {
		var curso Curso
		if i%2 == 0 {
			curso = C1
		} else {
			curso = C2
		}
		listaDeEstudantes = append(listaDeEstudantes, &estudante{nome: "Estudante" + fmt.Sprint(i), curso: curso})
        fmt.Print(*listaDeEstudantes[i])
	}
}

type Curso int

const (
	C1 = iota
	C2
)

type estudante struct {
	nome  string
	curso Curso
}
