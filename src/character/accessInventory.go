package character

import "fmt"

func AccessInventaire(p Character) {
	if len(p.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide.")
	} else {
		fmt.Println("Inventaire :")
		for _, objet := range p.Inventaire {
			fmt.Println("- " + objet)
		}
	}
}

func AddInventaire(p *Character, item string) {
	p.Inventaire = append(p.Inventaire, item)
}

func RemoveItem(p *Character, item string) bool {
	for i, it := range p.Inventaire {
		if it == item {
			p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
			return true
		}
	}
	return false
}
