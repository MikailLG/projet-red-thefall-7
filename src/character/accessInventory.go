package character

import "fmt"

func AccessInventaire(p Character) {
	fmt.Println("==============================")
	fmt.Println("        Inventaire")
	fmt.Println("==============================")

	if len(p.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide.")
	} else {
		for nom, quantite := range p.Inventaire {
			fmt.Printf("  - %s x%d\n", nom, quantite)
		}
	}
}

func AddInventaire(p *Character, item string) {
	if p.Inventaire == nil {
		p.Inventaire = make(map[string]int)
	}
	p.Inventaire[item]++
	if p.item >= 10  {
		return fmt.Println("Votre inventaire est plein !")
	}
}

func RemoveItem(p *Character, item string) bool {
	if quantite, exists := p.Inventaire[item]; exists {
		if quantite > 1 {
			p.Inventaire[item]--
		} else {
			delete(p.Inventaire, item)
		}
		return true
	}
	return false
}
