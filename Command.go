package main

import "fmt"

type Command interface {
	Execute() int
}

type AddCommand struct {
	a, b int
}

func (c *AddCommand) Execute() int {
	return c.a + c.b
}

type SubtractCommand struct {
	a, b int
}

func (c *SubtractCommand) Execute() int {
	return c.a - c.b
}

type Calculator struct {
	command Command
}

func (calc *Calculator) SetCommand(command Command) {
	calc.command = command
}

func (calc *Calculator) Compute() int {
	return calc.command.Execute()
}

func main() {
	calculator := &Calculator{}

	addCommand := &AddCommand{a: 5, b: 3}
	subtractCommand := &SubtractCommand{a: 10, b: 4}

	calculator.SetCommand(addCommand)
	result := calculator.Compute()
	fmt.Printf("Результат сложения: %d\n", result)

	calculator.SetCommand(subtractCommand)
	result = calculator.Compute()
	fmt.Printf("Результат вычитания: %d\n", result)
}
