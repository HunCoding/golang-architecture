package singleresponsibility

type FuncPJ struct{}
type FuncCLT struct{}
type FuncEstagiario struct{}

type Funcionario interface {
	CalcularSalario() float64
}

func main() {
	funcionario := &FuncPJ{}
	funcionario.CalcularSalario()

	stag := &FuncEstagiario{}
	stag.CalcularSalario()

	FolhaDePagamento(stag)
}

func (f *FuncPJ) CalcularSalario() float64 {
	//Calcular salario
	return 0.0
}

func (f *FuncCLT) CalcularSalario() float64 {
	//Calcular salario + beneficios
	return 0.0
}

func (f *FuncEstagiario) CalcularSalario() float64 {
	//Calcular bolsa auxilio
	return 0.0
}

func FolhaDePagamento(funcionario Funcionario) float64 {
	return funcionario.CalcularSalario()
}
