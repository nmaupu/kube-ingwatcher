package srv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nmaupu/kube-ingwatcher/config"
	"github.com/nmaupu/kube-ingwatcher/controller"
	"github.com/nmaupu/kube-ingwatcher/event"
	"log"
	"net/http"
)

func StartSender(p *config.IngressSenderParams) {

	sender := event.Sender{
		Callback: func(action int, payload config.ClientPayload) {
			// Checking config's filter to know if we send or not this payload
			if !p.In(payload.Labels) {
				log.Printf("Label filtering prevents this ingress to be sent, aborting. %+v\n", payload)
				return
			}

			var err error
			var jsn []byte
			var req *http.Request
			var resp *http.Response

			// Try to marshal the payload object
			jsn, err = json.Marshal(payload)
			if err != nil {
				log.Printf("%v\n", err.Error())
				return
			}

			method := ""
			switch action {
			case config.ACTION_ADD:
				method = "PUT"
			case config.ACTION_DELETE:
				method = "DELETE"
			}

			addr := fmt.Sprintf("http://%s:%d", p.GetDestAddr(), p.GetDestPort())

			log.Printf("%s -> %s with payload: %+v", method, addr, payload)

			req, err = http.NewRequest(
				method,
				fmt.Sprintf("http://%s:%d", p.GetDestAddr(), p.GetDestPort()),
				bytes.NewBuffer(jsn),
			)
			if err != nil {
				log.Printf("%v\n", err.Error())
				return
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err = client.Do(req)
			if err != nil {
				log.Printf("%v\n", err.Error())
				return
			}
			defer resp.Body.Close()
		},
	}

	controller.InitIngressController(func() event.Event { return sender })
}
