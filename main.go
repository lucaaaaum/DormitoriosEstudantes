package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	estudantes := iniciarListaDeEstudantes()
	estudantes = registrarIncompatibilidades(estudantes)

	estudantesC1 := filtrarEstudantesPorTipo(estudantes, C1)
	rand.Shuffle(len(estudantesC1), func(i, j int) { estudantesC1[i], estudantesC1[j] = estudantesC1[j], estudantesC1[i] })
	imprimirEstudantes(estudantesC1)
	estudantesC2 := filtrarEstudantesPorTipo(estudantes, C2)
	rand.Shuffle(len(estudantesC2), func(i, j int) { estudantesC2[i], estudantesC2[j] = estudantesC2[j], estudantesC2[i] })
	imprimirEstudantes(estudantesC2)

	pares, err := obterPares(estudantesC1, estudantesC2, 100)
	for i := 0; i < len(pares); i++ {
		imprimirPar(i, pares[i])
	}
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func obterPares(estudantesC1 []estudante, estudantesC2 []estudante, quantidadeDeParesEsperada int) ([]par, error) {
	var pares []par
	for i := 0; i < len(estudantesC1); i++ {
		for j := 0; j < len(estudantesC2); j++ {
			if estudanteJáFoiEscolhido(estudantesC2[j], pares) || sãoIncompatíveis(estudantesC1[i], estudantesC2[j]) {
				continue
			}
			pares = append(pares, par{&estudantesC1[i], &estudantesC2[j]})
		}
	}
	if len(pares) != quantidadeDeParesEsperada {
		return nil, errors.New("Quantidade de pares inesperada. Esperava [" + fmt.Sprint(quantidadeDeParesEsperada) + "], mas conseguiu formar [" + fmt.Sprint(len(pares)) + "]")
	}
	return pares, nil
}

type par struct {
	estudante1, estudante2 *estudante
}

func imprimirPar(n int, p par) {
	texto1 := p.estudante1.str()
	texto2 := p.estudante2.str()

	fmt.Println("[" + fmt.Sprint(n) + "] Par: ")
	fmt.Println("         " + texto1)
	fmt.Println("         " + texto2)
}

func estudanteJáFoiEscolhido(e estudante, par []par) bool {
	for i := 0; i < len(par); i++ {
		if par[i].estudante2.nome == e.nome {
			return true
		}
	}
	return false
}

func sãoIncompatíveis(e1, e2 estudante) bool {
	for i := 0; i < len(e1.estudantesIncompatíveis); i++ {
		if e1.estudantesIncompatíveis[i].nome == e2.nome {
			return true
		}
	}

	for i := 0; i < len(e2.estudantesIncompatíveis); i++ {
		if e2.estudantesIncompatíveis[i].nome == e1.nome {
			return true
		}
	}

	return false
}

// Funções para gerar a lista de estudantes e imprimir as informações

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
	s := "Nome: " + e.nome + " "
	var curso string
	if e.curso == C1 {
		curso = "C1"
	} else {
		curso = "C2"
	}
	s += "Curso: " + curso + " "
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
