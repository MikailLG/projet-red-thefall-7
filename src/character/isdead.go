package character

import (
	"fmt"
)

func isDead(personnage *Character) bool {
	if personnage.HP <= 0 {
		fmt.Printf("%s est mort(e)...\n", &personnage.Name)
		personnage.HP = personnage.MaxHP / 2
		fmt.Printf("%s ressuscite avec %d PV.\n", &personnage.Name, personnage.HP)
		return true
	}
	return false
}
