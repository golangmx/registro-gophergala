package teams

import "database"

// Team representa a un equipo
type Team struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ProjectName string `json:"project_name"`
	//Members *[]members.Member
}

// AllTeams hace lo que viene en la etiqueta
func AllTeams() (*[]Team, error) {
	rows, err := database.DB().Query("SELECT * FROM teams")
	if err != nil {
		return nil, err
	}
	var teams []Team
	for rows.Next() {
		var t Team
		err = rows.Scan(&t.ID, &t.Name, &t.ProjectName)
		if err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}
	err = rows.Err()
	return &teams, err
}
