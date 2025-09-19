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
	if _, exists := p.Inventaire[item]; exists || len(p.Inventaire) < 10 {
		p.Inventaire[item]++
		fmt.Println(item, Green, "ajouté à l'inventaire.")
	} else {
		fmt.Println(Red, "Votre inventaire est plein ! Impossible d'ajouter", item)
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

func UpgradeInventorySlot(p *Character) {
	if p.Upgrade >= 3 {
		fmt.Printf("%sVous ne pouvez plus augmenter votre inventaire.%s\n", Red, Reset)
		return
	}

	p.CapaciteMax += 10
	p.Upgrade++
	fmt.Printf("%sVotre inventaire a été augmenté de 10 slots ! Capacité actuelle : %d%s\n", Green, p.CapaciteMax, Reset)
}
