package members

import "database"

// Member representa a un miembro del equipo
type Member struct {
	ID        int    `json:"id"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	TipoID    int    `json:"tipo_id"`
	NumeroID  string `json:"numero_id"`
}

// ForTeam toma una id y busca los miembros pertenecientes
// a ese equipo
func ForTeam(id int) ([]Member, error) {
	rows, err := database.DB().Query("SELECT id, nombres, apellidos, tipo_id,"+
		"numero_id FROM members WHERE team=$1", id)
	if err != nil {
		return nil, err
	}
	members := make([]Member, 0)
	for rows.Next() {
		var m Member
		err = rows.Scan(&m.ID, &m.Nombres, &m.Apellidos,
			&m.TipoID, &m.NumeroID)
		if err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	err = rows.Err()
	return members, err
}

// Save guarda un miembro en la base de datos
func (m *Member) Save(teamID int) error {
	_, err := database.DB().Exec("INSERT INTO members (nombres, apellidos, "+
		"tipo_id, numero_id, team) VALUES "+
		"($1, $2, $3, $4, $5)",
		m.Nombres, m.Apellidos, m.TipoID, m.NumeroID, teamID)
	return err
}

// SaveMultiple guarda varios miembros de manera concurrente
func SaveMultiple(members []Member, teamID int) <-chan error {
	c := make(chan error)
	for i := range members {
		go func(m Member) {
			err := m.Save(teamID)
			c <- err
		}(members[i])
	}
	return c
}
