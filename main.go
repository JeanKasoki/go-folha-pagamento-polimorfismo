package main

import (
	"fmt"
)

type funcionario interface{
	CalcularSalario() float64
}

type efetivo struct{
	nome string
	salarioBase float64
	bonus float64
}

type terceirizado struct{
	nome string
	horasTrabalhadas float64
	valorHora float64
}

func (e efetivo) CalcularSalario() float64{
		return e.salarioBase + e.bonus	
}

func (t terceirizado) CalcularSalario() float64{
	return t.horasTrabalhadas * t.valorHora
}

func processarFolha(f []funcionario){
	total := 0.0
	for _, v := range f{
		salario := v.CalcularSalario()

	fmt.Printf("Salário calculado: %.2f\n", salario)
		total += salario
	}
	fmt.Printf("Custo total da folha: %.2f\n", total)
}

func main(){
	anderson := efetivo{"anderson", 3600.0, 1200.0}
	maria := terceirizado{"maria", 6, 15}

	x := []funcionario{anderson, maria}

	processarFolha(x)
}
