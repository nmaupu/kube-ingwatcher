package config

type IngressSenderParams struct {
	DestAddr *string
	DestPort *int
}

func (p *IngressSenderParams) GetDestAddr() string {
	return *p.DestAddr
}

func (p *IngressSenderParams) GetDestPort() int {
	return *p.DestPort
}
