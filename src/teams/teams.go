package teams

import (
	"database"
	"members"
)

// Team representa a un equipo
type Team struct {
	ID       int               `json:"id"`
	Nombre   string            `json:"nombre"`
	Proyecto string            `json:"proyecto"`
	Members  []*members.Member `json:"members"`
}

// AllTeams hace lo que viene en la etiqueta
func AllTeams() ([]Team, error) {
	rows, err := database.DB().Query("SELECT * FROM teams")
	if err != nil {
		return nil, err
	}
	teams := make([]Team, 0)
	for rows.Next() {
		var t Team
		err = rows.Scan(&t.ID, &t.Nombre, &t.Proyecto)
		if err != nil {
			return nil, err
		}
		t.Members, err = members.ForTeam(t.ID)
		if err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}
	err = rows.Err()
	return teams, err
}

// Save guarda un equipo en la base de datos
func (t *Team) Save() error {
	_, err := database.DB().Exec("INSERT INTO teams (nombre, proyecto) "+
		"VALUES ($1, $2)", t.Nombre, t.Proyecto)
	if err != nil {
		return err
	}
	c := members.SaveMultiple(t.Members)
	for e := range c {
		if e != nil {
			return e
		}
	}
	return nil
}
