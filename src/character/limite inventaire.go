package character

import "fmt"

type Character struct {
	Name       string
	Inventaire []string
}

func AddInventaire(p *Character, item string) {
	if len(p.Inventaire) < 10 {
		p.Inventaire = append(p.Inventaire, item)
		fmt.Println(item, "ajouté.")
	} else {
		fmt.Println("Inventaire plein !")
	}
}
