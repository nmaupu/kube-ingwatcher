package srv

import (
	"encoding/json"
	"fmt"
	"github.com/nmaupu/kube-ingwatcher/config"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	params *config.IngressReceiverParams
)

func handleError(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Printf("Error encountered: %v\n", err.Error())
		http.Error(w, err.Error(), 400)
		return true
	}

	return false
}

func doExecCmd(w http.ResponseWriter, execCmd string, payload *config.ClientPayload) {
	if execCmd != "" {
		log.Printf("Executing command %s\n", execCmd)
		err := exec.Command(execCmd, strings.Join(payload.Hosts, ",")).Run()
		if handleError(w, err) {
			return
		}
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var payload config.ClientPayload

	if r.Body == nil {
		http.Error(w, "Please send a payload", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if handleError(w, err) {
		return
	}

	destFile := ""
	if params.GetTemplate() != "" && params.GetDestination() != "" {
		destFile = fmt.Sprintf("%s/%s%s%s", params.GetDestination(), params.GetPrefix(), payload.Hosts[0], params.GetSuffix())
	}

	switch r.Method {
	case "PUT", "POST":
		if destFile != "" {
			log.Printf("Generating %s file from template %s\n", destFile, params.GetTemplate())
			err = payload.GenerateTemplate(params.GetTemplate(), destFile)
			if handleError(w, err) {
				return
			}

			defer doExecCmd(w, params.GetExecCmdAdd(), &payload)
		}

	case "DELETE":
		if destFile != "" {
			log.Printf("Deleting %s file\n", destFile)
			err = os.Remove(destFile)
			if handleError(w, err) {
				return
			}

			defer doExecCmd(w, params.GetExecCmdDelete(), &payload)
		}
	}
}

func StartReceiver(p *config.IngressReceiverParams) {
	log.Printf("Starting http server on %s:%d\n", p.GetBindAddr(), p.GetPort())
	params = p

	http.HandleFunc("/", handleRequest)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", p.GetBindAddr(), p.GetPort()), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
