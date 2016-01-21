package router

import (
	"logger"
	"net/http"
)

// Route se encarga de manejar las rutas
func Route(res http.ResponseWriter, req *http.Request) {
	logger.LogRequest(req)
	switch req.Method {
	case "POST":
		postTeam(res, req)
	case "GET":
		getTeams(res)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(""))
	}
}
