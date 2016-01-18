package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// LogOut escribe un mensaje a Stdout
func LogOut(message string) {
	log.SetOutput(os.Stdout)
	log.Println(message)
}

// LogErr escribe un mensaje a Stderr
func LogErr(message string) {
	log.SetOutput(os.Stderr)
	log.Println(message)
}

// LogRequest usa LogOut para escribir la información general del request
func LogRequest(req *http.Request) {
	LogOut(fmt.Sprintf("%s '%s' - %s", req.Method, req.URL, req.RemoteAddr))
}
