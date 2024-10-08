package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	seed, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	randomSource := rand.NewSource(int64(seed))
	random := rand.New(randomSource)

	quantidadeDeEstudantesC1, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	quantidadeDeEstudantesC2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}

	estudantesC1 := iniciarListaDeEstudantes(quantidadeDeEstudantesC1, C1)
	random.Shuffle(len(estudantesC1), func(i, j int) { estudantesC1[i], estudantesC1[j] = estudantesC1[j], estudantesC1[i] })
	estudantesC2 := iniciarListaDeEstudantes(quantidadeDeEstudantesC2, C2)
	random.Shuffle(len(estudantesC2), func(i, j int) { estudantesC2[i], estudantesC2[j] = estudantesC2[j], estudantesC2[i] })

	estudantesC1 = registrarIncompatibilidades(estudantesC1, estudantesC2, *random)
	estudantesC2 = registrarIncompatibilidades(estudantesC2, estudantesC1, *random)

	quantidadeDeDormitórios, err := strconv.Atoi(os.Args[4])
	if err != nil {
		panic(err)
	}

	pares, estudantesSemPar := obterPares(estudantesC1, estudantesC2, quantidadeDeDormitórios)
	fmt.Println("PARES")
	for i := 0; i < len(pares); i++ {
		imprimirPar(i, pares[i])
	}
	fmt.Println("ESTUDANTES SEM PAR")
	for i := 0; i < len(estudantesSemPar); i++ {
		fmt.Printf("%v\n", estudantesSemPar[i].str())
	}
}

func obterPares(estudantesC1, estudantesC2 []estudante, quantidadeDeDormitórios int) ([]par, []estudante) {
	var pares []par
	quantidadeMáximaDePares := quantidadeDeDormitórios / 2
	for i := 0; i < len(estudantesC1); i++ {
		e1 := estudantesC1[i]
		var e2 estudante
		for j := 0; j < len(estudantesC2); j++ {
			if estudanteJáFoiEscolhido(estudantesC2[j], pares) || sãoIncompatíveis(estudantesC1[i], estudantesC2[j]) {
				continue
			}
			e2 = estudantesC2[j]
			break
		}
		pares = append(pares, par{&e1, &e2})
		if len(pares) >= quantidadeMáximaDePares {
			break
		}
	}
	var estudantesSemPar []estudante
	if len(pares) != quantidadeMáximaDePares {
		for i := 0; i < len(estudantesC1); i++ {
			if !estudanteJáFoiEscolhido(estudantesC1[i], pares) {
				estudantesSemPar = append(estudantesSemPar, estudantesC1[i])
			}
		}
		for i := 0; i < len(estudantesC2); i++ {
			if !estudanteJáFoiEscolhido(estudantesC2[i], pares) {
				estudantesSemPar = append(estudantesSemPar, estudantesC2[i])
			}
		}
	}
	return pares, estudantesSemPar
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
		if par[i].estudante1.nome == e.nome || par[i].estudante2.nome == e.nome {
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

func iniciarListaDeEstudantes(quantidade int, curso Curso) []estudante {
	listaDeEstudantes := make([]estudante, 0)
	for i := 0; i < quantidade; i++ {
		listaDeEstudantes = append(listaDeEstudantes, estudante{nome: "Estudante" + fmt.Sprint(i), curso: curso})
	}
	return listaDeEstudantes
}

func registrarIncompatibilidades(estudantes1, estudantes2 []estudante, random rand.Rand) []estudante {
	for i := 0; i < len(estudantes1); i++ {
		quantidadeDeIncompatíveis := random.Intn(len(estudantes2) - 1)
		for j := 0; j < quantidadeDeIncompatíveis; j++ {
			var estudanteIncompatívelAleatório int
			for true {
				estudanteIncompatívelAleatório = random.Intn(len(estudantes2) - 1)
                var estudanteJáFoiRegistradoComoIncompatível bool
                for k := 0; k < len(estudantes1[i].estudantesIncompatíveis); k++ {
                    estudanteJáFoiRegistradoComoIncompatível = estudantes1[i].estudantesIncompatíveis[k].nome == estudantes2[estudanteIncompatívelAleatório].nome
                    if estudanteJáFoiRegistradoComoIncompatível {
                        break
                    }
                }
                if !estudanteJáFoiRegistradoComoIncompatível {
                    break
                }
			}
			estudantes1[i].estudantesIncompatíveis = append(
				estudantes1[i].estudantesIncompatíveis,
				&estudantes2[estudanteIncompatívelAleatório],
			)
		}
	}
	return estudantes1
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
