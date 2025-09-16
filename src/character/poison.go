package character

import "fmt"

func poison(p *Character) {
	for i := 1; i <= 3; i++ {
		cible.PointDeVie = cible.PointDeVie - 10
		fmt.Println("PV actuels :", cible.PointDeVie + cible.PointDeVieMax)
		time.sleep(1 * time.seconde)
	}
}
                                       