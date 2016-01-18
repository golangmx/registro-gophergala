package router

import (
	"encoding/json"
	"fmt"
	"logger"
	"net/http"
	"teams"
)

func getTeams(res http.ResponseWriter) {
	teams, err := teams.AllTeams()
	if err != nil {
		logger.LogErr(fmt.Sprintf("error de db: %s", err.Error()))
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(""))
		return
	}
	res.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(res)
	enc.Encode(teams)
}
