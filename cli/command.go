package cli

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/nmaupu/kube-ingwatcher/config"
	"github.com/nmaupu/kube-ingwatcher/srv"
	"os"
)

const (
	DEFAULT_PORT = 6565
)

func Process(appName, appDesc, appVersion string) {
	app := cli.App(appName, appDesc)
	app.Version("v version", fmt.Sprintf("%s version %s", appName, appVersion))

	app.Command("ingressSender is", "Ingress sender mode", ingressSender)
	app.Command("ingressReceiver ir", "Ingress receiver mode", ingressReceiver)

	app.Run(os.Args)
}

func ingressSender(cmd *cli.Cmd) {
	params := config.IngressSenderParams{
		DestAddr: cmd.StringOpt("a addr", "", "Destination address to send ingress to"),
		DestPort: cmd.IntOpt("p port", DEFAULT_PORT, "Default port to connect to"),
	}

	cmd.Action = func() {
		srv.StartSender(&params)
	}
}

func ingressReceiver(cmd *cli.Cmd) {
	params := config.IngressReceiverParams{
		BindAddr:      cmd.StringOpt("a addr", "0.0.0.0", "Interface address to bind"),
		Port:          cmd.IntOpt("p port", DEFAULT_PORT, "Port to bind"),
		Template:      cmd.StringOpt("t template", "", "Template to render when an ingress is received"),
		Destination:   cmd.StringOpt("d destination", "", "Destination directory where to generate templates"),
		Prefix:        cmd.StringOpt("p prefix", "", "Prefix to use for filename when generating a template"),
		Suffix:        cmd.StringOpt("s suffix", "", "Suffix to use for filenme when generating a template"),
		ExecCmdAdd:    cmd.StringOpt("A addCmd", "", "Command to execute after rendering a file from the template"),
		ExecCmdDelete: cmd.StringOpt("D deleteCmd", "", "Command to execute after deleting a file previously created with the template"),
	}

	cmd.Action = func() {
		srv.StartReceiver(&params)
	}
}
