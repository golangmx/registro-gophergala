package teams

import (
	"database"
	"members"
)

// Team representa a un equipo
type Team struct {
	ID       int              `json:"id"`
	Nombre   string           `json:"nombre"`
	Proyecto string           `json:"proyecto"`
	Miembros []members.Member `json:"miembros"`
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
		t.Miembros, err = members.ForTeam(t.ID)
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
	var teamID int
	err := database.DB().QueryRow("INSERT INTO teams (nombre, proyecto) "+
		"VALUES ($1, $2) RETURNING id", t.Nombre, t.Proyecto).Scan(&teamID)
	if err != nil {
		return err
	}
	c := members.SaveMultiple(t.Miembros, int(teamID))
	for e := range c {
		if e != nil {
			return e
		}
	}
	return nil
}
