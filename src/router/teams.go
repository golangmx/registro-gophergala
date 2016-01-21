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
	err = enc.Encode(teams)
	if err != nil {
		logger.LogErr(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(""))
	}
}

func postTeam(res http.ResponseWriter, req *http.Request) {
	var t teams.Team
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&t)
	if err != nil {
		logger.LogErr(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(""))
		return
	}
	err = t.Save()
	if err != nil {
		logger.LogErr(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(""))
		return
	}
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(""))
}
