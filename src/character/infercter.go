package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Infecter struct {
	Nom           string
	PointDeVieMax int
	PointDeVie    int
	PointAttaque  int
}

func NewInfecter(nom string, pvMax int, attaque int) Infecter {
	return Infecter{
		Nom:           nom,
		PointDeVieMax: pvMax,
		PointDeVie:    pvMax,
		PointAttaque:  attaque,
	}
}

func CombatEntrainement(j *Character, m *Infecter) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Un %s apparaît !\n", m.Nom)
	time.Sleep(500 * time.Millisecond)
	tour := 1
	for j.PointDeVie > 0 && m.PointDeVie > 0 {
		ClearScreen()
		fmt.Println("=== Nouveau tour ===")
		fmt.Printf("PV Joueur : %d/%d | PV %s : %d/%d\n\n", j.PointDeVie, j.PointDeVieMax, m.Nom, m.PointDeVie, m.PointDeVieMax)
		fmt.Println("Votre tour :")
		fmt.Println("1) Attaquer")
		fmt.Println("2) Se soigner (+10 PV)")
		fmt.Print("Choix : ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			m.PointDeVie -= 10
			if m.PointDeVie < 0 {
				m.PointDeVie = 0
			}
			fmt.Printf("Vous attaquez %s !\n", m.Nom)
		case "2":
			j.PointDeVie += 10
			if j.PointDeVie > j.PointDeVieMax {
				j.PointDeVie = j.PointDeVieMax
			}
			fmt.Println("Vous vous soignez !")
		default:
			fmt.Println("Choix invalide ! Vous perdez votre tour.")
		}
		if m.PointDeVie == 0 {
			fmt.Println("Vous avez vaincu le monstre !")
			break
		}
		j.PointDeVie -= m.PointAttaque
		if j.PointDeVie < 0 {
			j.PointDeVie = 0
		}
		fmt.Printf("%s inflige à %s %d de dégâts\n", m.Nom, j.Nom, m.PointAttaque)
		fmt.Printf("%s a maintenant %d/%d PV\n", j.Nom, j.PointDeVie, j.PointDeVieMax)
		fmt.Println("-----------------------")

		if j.PointDeVie == 0 {
			fmt.Println("Vous avez été vaincu...")
			break
		}
		tour++
		time.Sleep(500 * time.Millisecond)
	}
}

func InitColosse() Infecter {
	return NewInfecter("Colosse", 100, 15)
}

func CombatColosse(j *Character, c *Infecter) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Un %s apparaît !\n", c.Nom)
	time.Sleep(500 * time.Millisecond)
	tour := 1
	for j.PointDeVie > 0 && c.PointDeVie > 0 {
		ClearScreen()
		fmt.Println("=== Nouveau tour ===")
		fmt.Printf("PV Joueur : %d/%d | PV Colosse : %d/%d\n\n", j.PointDeVie, j.PointDeVieMax, c.PointDeVie, c.PointDeVieMax)
		fmt.Println("Votre tour :")
		fmt.Println("1) Attaquer")
		fmt.Println("2) Se soigner (+10 PV)")
		fmt.Print("Choix : ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			c.PointDeVie -= 10
			if c.PointDeVie < 0 {
				c.PointDeVie = 0
			}
			fmt.Printf("Vous attaquez %s !\n", c.Nom)
		case "2":
			j.PointDeVie += 10
			if j.PointDeVie > j.PointDeVieMax {
				j.PointDeVie = j.PointDeVieMax
			}
			fmt.Println("Vous vous soignez !")
		default:
			fmt.Println("Choix invalide ! Vous perdez votre tour.")
		}
		if c.PointDeVie == 0 {
			fmt.Println("Vous avez vaincu le Colosse !")
			break
		}
		var degats int
		if tour%3 == 0 {
			degats = c.PointAttaque * 2
		} else {
			degats = c.PointAttaque
		}

		j.PointDeVie -= degats
		if j.PointDeVie < 0 {
			j.PointDeVie = 0
		}
		fmt.Printf("%s inflige à %s %d de dégâts\n", c.Nom, j.Nom, degats)
		fmt.Printf("%s a maintenant %d/%d PV\n", j.Nom, j.PointDeVie, j.PointDeVieMax)
		fmt.Println("-----------------------")
		if j.PointDeVie == 0 {
			fmt.Println("Vous avez été vaincu par le Colosse !")
			break
		}
		tour++
		time.Sleep(500 * time.Millisecond)
	}
}

func EntrainementMenu(j *Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("==============================")
		fmt.Println("        Menu Entraînement")
		fmt.Println("==============================")
		fmt.Println("1) Combattre un Rodeur")
		fmt.Println("2) Combattre un Colosse")
		fmt.Println("0) Retour au menu principal")
		fmt.Print("Votre choix : ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			rodeur := NewInfecter("Rodeur", 40, 5)
			CombatEntrainement(j, &rodeur)
		case "2":
			colosse := InitColosse()
			CombatColosse(j, &colosse)
		case "0":
			fmt.Println("Retour au menu principal...")
			time.Sleep(500 * time.Millisecond)
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
		fmt.Println()
	}
}
