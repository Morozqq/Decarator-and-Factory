package main

import "fmt"

type IGuitar interface {
	setName(name string)
	setPrice(price int)
	getName() string
	getPrice() int
}

type Guitar struct {
	name  string
	price int
}

func (g *Guitar) setName(name string) {
	g.name = name
}

func (g *Guitar) getName() string {
	return g.name
}

func (g *Guitar) setPrice(price int) {
	g.price = price
}

func (g *Guitar) getPrice() int {
	return g.price
}

type AcousticGuitar struct {
	Guitar
}

func newAcousticGuitar() IGuitar {
	return &AcousticGuitar{
		Guitar: Guitar{
			name:  "Acoustic Guitar",
			price: 100,
		},
	}
}

type ElectricGuitar struct {
	Guitar
}

func newElectricGuitar() IGuitar {
	return &ElectricGuitar{
		Guitar: Guitar{
			name:  "Electric Guitar",
			price: 150,
		},
	}
}

func getGuitar(guitarType string) (IGuitar, error) {
	if guitarType == "Acoustic" {
		return newAcousticGuitar(), nil
	}
	if guitarType == "Electric" {
		return newElectricGuitar(), nil
	}
	return nil, fmt.Errorf("Wrong Guitar type passed")
}

func printDetails(g IGuitar) {
	fmt.Printf("Guitar: %s", g.getName())
	fmt.Println()
	fmt.Printf("Price: %d", g.getPrice())
	fmt.Println()
}
func main() {
	Acoustic, _ := getGuitar("Acoustic")
	Electric, _ := getGuitar("Electric")

	printDetails(Acoustic)
	printDetails(Electric)

}
