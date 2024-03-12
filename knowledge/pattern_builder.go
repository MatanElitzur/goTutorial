package knowledge

import "fmt"

type Pizza struct {
	Size      string
	Crust     string
	Cheese    bool
	Pepperoni bool
	Mushrooms bool
	// other toppings..
}

type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder
	SetCrust(crust string) PizzaBuilder
	AddCheese() PizzaBuilder
	AddPepperoni() PizzaBuilder
	AddMushrooms() PizzaBuilder
	Build() Pizza
}

type ConcurrentPizzaBuilder struct {
	pizza Pizza
}

func (b *ConcurrentPizzaBuilder) SetSize(size string) PizzaBuilder {
	b.pizza.Size = size
	return b
}

func (b *ConcurrentPizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.pizza.Crust = crust
	return b
}

func (b *ConcurrentPizzaBuilder) AddCheese() PizzaBuilder {
	b.pizza.Cheese = true
	return b
}

func (b *ConcurrentPizzaBuilder) AddPepperoni() PizzaBuilder {
	b.pizza.Pepperoni = true
	return b
}

func (b *ConcurrentPizzaBuilder) AddMushrooms() PizzaBuilder {
	b.pizza.Mushrooms = true
	return b
}

func (b *ConcurrentPizzaBuilder) Build() Pizza {
	return b.pizza
}

type PizzaDirector struct{}

func (d *PizzaDirector) CreateMargherita(builder PizzaBuilder) Pizza {
	return builder.SetSize("Medium").SetCrust("Thin").AddCheese().Build()
}

//Other predefind pizzas can be added...

func BuilderPattern() {
	builder := &ConcurrentPizzaBuilder{}
	director := PizzaDirector{}
	//Predefind Pizza
	margherita := director.CreateMargherita(builder)
	fmt.Println("Margherita:", margherita)
	//Custom Pizza
	customPizza := builder.SetSize("Large").AddPepperoni().AddMushrooms().Build()
	fmt.Println("Custom Pizza:", customPizza)
}
