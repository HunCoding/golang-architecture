package interfacesegregrationprinciple

type Trabalho interface {
	Entrar()

	ComecarATrabalhar()

	ContinuarATrabalhar()

	Sair()
}

type HumanoTasks interface {
	Trabalho

	PausaProCafe()

	Almocar()
}

type RoboTasks interface {
	Trabalho

	VerificarOleo()

	CarregarBateria()
}

type Robo struct{}

type Humano struct{}
