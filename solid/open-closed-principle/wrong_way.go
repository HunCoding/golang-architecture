package singleresponsibility

import "reflect"

type FuncionarioCLT struct{}
type FuncionarioPJ struct{}
type Estagiario struct{}

func FolhaDePag[TipoFuncionario FuncionarioCLT | Estagiario | FuncionarioPJ](funcionario TipoFuncionario) float64 {

	if reflect.TypeOf(FuncionarioCLT{}) == reflect.TypeOf(funcionario) {

		//Calcular salario + beneficios
		return 0.0

	} else if reflect.TypeOf(FuncionarioPJ{}) == reflect.TypeOf(funcionario) {

		// Calcular salario do PJ
		return 0.0

	} else {

		//Calcular bolsa auxilio do estagiario
		return 0.0

	}
}
