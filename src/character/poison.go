package character

import (
	"fmt"
	"time"
)

func poison(p *Character) {
	for i := 1; i <= 3; i++ {
		p.PointDeVie -= 10
		if p.PointDeVie < 0 {
			p.PointDeVie = 0
		}
		fmt.Printf("PV actuels : %d/%d\n", p.PointDeVie, p.PointDeVieMax)
		time.Sleep(1 * time.Second)
	}
}
