package main

import (
	"fmt"
	"math/rand"
)

func main() {
	estudantes := iniciarListaDeEstudantes()
	estudantes = registrarIncompatibilidades(estudantes)

	estudantesC1 := filtrarEstudantesPorTipo(estudantes, C1)

	imprimirEstudantes(estudantesC1)
	estudantesC2 := filtrarEstudantesPorTipo(estudantes, C2)
	imprimirEstudantes(estudantesC2)
}

func iniciarListaDeEstudantes() []estudante {
	listaDeEstudantes := make([]estudante, 0)
	for i := 0; i < 100; i++ {
		var curso Curso
		if i%2 == 0 {
			curso = C1
		} else {
			curso = C2
		}
		listaDeEstudantes = append(listaDeEstudantes, estudante{nome: "Estudante" + fmt.Sprint(i), curso: curso})
	}
	return listaDeEstudantes
}

func registrarIncompatibilidades(estudantes []estudante) []estudante {
	randomSource := rand.NewSource(123312) // número pré-definido para facilitar os testes
	random := rand.New(randomSource)
	for i := 0; i < len(estudantes); i++ {
		quantidadeDeIncompatíveis := random.Intn(5)
		for j := 0; j < quantidadeDeIncompatíveis; j++ {
			var estudanteIncompatívelAleatório int

			// tem que ser um loop porque não é permitido um aluno ser incompatível com ele mesmo
			for true {
				estudanteIncompatívelAleatório = random.Intn(len(estudantes) - 1)
				if estudanteIncompatívelAleatório != i {
					break
				}
			}

			estudantes[i].estudantesIncompatíveis = append(
				estudantes[i].estudantesIncompatíveis,
				&estudantes[estudanteIncompatívelAleatório],
			)
		}
	}
	return estudantes
}

func filtrarEstudantesPorTipo(estudantes []estudante, curso Curso) []estudante {
	estudantesFiltrados := make([]estudante, 0)
	for i := 0; i < len(estudantes); i++ {
		if estudantes[i].curso == curso {
			estudantesFiltrados = append(estudantesFiltrados, estudantes[i])
		}
	}
	return estudantesFiltrados
}

type Curso int

const (
	C1 = iota
	C2
)

type estudante struct {
	nome                    string
	curso                   Curso
	estudantesIncompatíveis []*estudante
}

func (e *estudante) str() string {
	s := "Nome: " + e.nome + "\n"
	var curso string
	if e.curso == C1 {
		curso = "C1"
	} else {
		curso = "C2"
	}
	s += "Curso: " + curso + "\n"
	s += "Incompatíveis: "
	if len(e.estudantesIncompatíveis) > 0 {
		for i := 0; i < len(e.estudantesIncompatíveis)-1; i++ {
			s += e.estudantesIncompatíveis[i].nome + ", "
		}
		s += e.estudantesIncompatíveis[len(e.estudantesIncompatíveis)-1].nome
	}

	return s
}

func imprimirEstudantes(estudantes []estudante) {
	for i := 0; i < len(estudantes); i++ {
		fmt.Println(estudantes[i].str())
	}
}
