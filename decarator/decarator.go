package main

import "fmt"

type Guitar interface {
	GetPrice() int
}

type OriginalGuitar struct{}

func (og OriginalGuitar) GetPrice() int {
	return 100
}

type TunedModification struct {
	Guitar Guitar
}

func (tm TunedModification) GetPrice() int {
	return tm.Guitar.GetPrice() + 10
}

type CedarModification struct {
	Guitar Guitar
}

func (cm CedarModification) GetPrice() int {
	return cm.Guitar.GetPrice() + 30
}

func main() {

	originalGuitar := OriginalGuitar{}

	tunedGuitar := TunedModification{Guitar: originalGuitar}
	fmt.Printf("Price of Tuned Guitar: $%d\n", tunedGuitar.GetPrice())

	cedarGuitar := CedarModification{Guitar: originalGuitar}
	fmt.Printf("Price of Cedar Guitar: $%d\n", cedarGuitar.GetPrice())
}
