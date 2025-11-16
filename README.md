# Projeto: Folha de Pagamento com Polimorfismo em Go

Este é um projeto de console simples para praticar os conceitos de **Interfaces** e **Polimorfismo** na linguagem Go.

## Desafio

O objetivo era criar um sistema que pudesse processar a folha de pagamento de diferentes tipos de funcionários. A função principal (`processarFolha`) deveria ser capaz de calcular o custo total da folha sem saber os detalhes de cálculo de cada tipo de funcionário.

## Solução (Polimorfismo)

Para resolver isso, usei os seguintes conceitos do Go:

1.  **Interface (`Funcionario`):** Defini um "contrato" que exige um único método: `CalcularSalario() float64`.
2.  **Structs Concretos (`Efetivo`, `Terceirizado`):** Defini os tipos de dados que guardam as propriedades de cada funcionário.
3.  **Métodos:** Implementei o método `CalcularSalario()` para cada _struct_, cumprindo o contrato da interface. Cada um tem sua própria lógica de cálculo.

A implementação do polimorfismo está na função `processarFolha`, que aceita um `[]Funcionario` (um slice da interface).

### Objetivo do Design: Código Desacoplado

A maior vantagem disso é a **extensibilidade**. A `processarFolha` não sabe (e não se importa) se está lidando com um `Efetivo` ou `Terceirizado`. Ela só opera no contrato da interface.

Se amanhã eu quiser adicionar um `type Estagiario`, basta criar o _struct_ e implementar o método `CalcularSalario()` para ele. A função `processarFolha` continuará funcionando **sem que eu precise alterar uma única linha** nela.
