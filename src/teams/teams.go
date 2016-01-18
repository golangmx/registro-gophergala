package teams

import "database"

// Team representa a un equipo
type Team struct {
	ID             int    `json:"id"`
	Nombre         string `json:"nombre"`
	NombreProyecto string `json:"nombre_proyecto"`
	//Members *[]members.Member
}

// AllTeams hace lo que viene en la etiqueta
func AllTeams() (*[]Team, error) {
	rows, err := database.DB().Query("SELECT * FROM teams")
	if err != nil {
		return nil, err
	}
	teams := make([]Team, 0)
	for rows.Next() {
		var t Team
		err = rows.Scan(&t.ID, &t.Nombre, &t.NombreProyecto)
		if err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}
	err = rows.Err()
	return &teams, err
}
