package main

import (
	"fmt"
)

type waste struct {
	Name     string
	Type     string
	Quantity int
}

type Dustbin struct {
	Name     string
	Type     string
	contents []waste
}

var counstwaste = 0 int

func (d *Dustbin) DisplayInfo() {
	fmt.Println("Poubelle", d.Name, "-", d.Type)
	if len(d.contents) == 0 {
		fmt.Println("Vide")
	} else {
		for _, w := range d.contents {
			fmt.Println("-", w.Name, "x", w.Quantity)
		}
	}
}

func Addwaste(d *Dustbin, w waste) {
	if w.Type == d.Type {
		for _, waste := range d.contents{
		fmt.Println("ajouté dans", d.Name, ":", w.Name)
	} else {
		fmt.Println("Ajout Impossible")
		}

	}
}

func (d *Dustbin) Empty() {
	if len(d.contents) == 0 {
		fmt.Println(d.Name, "deja vide")
	}
	return

}

count := 0
for _, w:= range