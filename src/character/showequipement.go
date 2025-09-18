package character

type Equipement struct {
	Tête  string
	Torse string
	Pieds string
}

type character struct {
	Nom        string
	Pointdevie int
	Equipement Equipement
}
