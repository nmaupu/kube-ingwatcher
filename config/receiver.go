package config

type IngressReceiverParams struct {
	BindAddr       *string
	Port           *int
	Template       *string
	Destination    *string
	Prefix, Suffix *string
	ExecCmdAdd     *string
	ExecCmdDelete  *string
}

func (p *IngressReceiverParams) GetBindAddr() string {
	return *p.BindAddr
}

func (p *IngressReceiverParams) GetPort() int {
	return *p.Port
}

func (p *IngressReceiverParams) GetTemplate() string {
	return *p.Template
}

func (p *IngressReceiverParams) GetDestination() string {
	return *p.Destination
}

func (p *IngressReceiverParams) GetPrefix() string {
	return *p.Prefix
}

func (p *IngressReceiverParams) GetSuffix() string {
	return *p.Suffix
}

func (p *IngressReceiverParams) GetExecCmdAdd() string {
	return *p.ExecCmdAdd
}

func (p *IngressReceiverParams) GetExecCmdDelete() string {
	return *p.ExecCmdDelete
}
