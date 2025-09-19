package character

import "fmt"

type Recette struct {
	Name        string
	Ingredients map[string]int
	ObjetCraft  string
}

var Recipes = []Recette{
	{
		Name: "Casque renforcé",
		Ingredients: map[string]int{
			"Kevlar": 3,
			"Tissu":  2,
		},
		ObjetCraft: "Casque renforcé",
	},
	{
		Name: "Gilet pare-balles",
		Ingredients: map[string]int{
			"Kevlar":       3,
			"Plaque de fer": 2,
			"Tissu":        1,
		},
		ObjetCraft: "Gilet pare-balles",
	},
	{
		Name: "Bottes renforcées",
		Ingredients: map[string]int{
			"Cuir":          2,
			"Plaque de fer": 1,
		},
		ObjetCraft: "Bottes renforcées",
	},
}

func CheckIngredient(p *Character, recipe Recette) bool {
	for item, quantite := range recipe.Ingredients {
		if p.Inventaire[item] < quantite {
			return false
		}
	}
	return true
}

func Craft(p *Character, recipe Recette) {
	if !CheckIngredient(p, recipe) {
		fmt.Printf("%sVous n'avez pas les ressources nécessaires pour crafter : %s%s\n", Red, recipe.Name, Reset)
		return
	}

	for item, qty := range recipe.Ingredients {
		p.Inventaire[item] -= qty
		if p.Inventaire[item] <= 0 {
			delete(p.Inventaire, item)
		}
	}

	AddInventaire(p, recipe.ObjetCraft)
	fmt.Printf("%sVous avez crafté : %s%s\n", Green, recipe.ObjetCraft, Reset)
}

func CraftMenu(p *Character) {
	PlayMusic("craft.mp3")
	defer StopMusic() 
	for {
		fmt.Printf("%s==============================%s\n", Blue, Reset)
		fmt.Printf("%s          Atelier Craft%s\n", Blue, Reset)
		fmt.Printf("%s==============================%s\n", Blue, Reset)

		for i, recipe := range Recipes {
			fmt.Printf("%d) %s (", i+1, recipe.Name)
			for item, quantite := range recipe.Ingredients {
				fmt.Printf("%s x%d ", item, quantite)
			}
			fmt.Println(")")
		}

		fmt.Println("0) Retour")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		if choix == 0 {
			return
		}
		if choix < 1 || choix > len(Recipes) {
			fmt.Println("Choix invalide.")
			continue
		}

		Craft(p, Recipes[choix-1])
	}
}
