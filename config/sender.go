package config

import (
	"strings"
)

type IngressSenderParams struct {
	DestAddr    *string
	DestPort    *int
	LabelFilter *string
}

func (p *IngressSenderParams) GetDestAddr() string {
	return *p.DestAddr
}

func (p *IngressSenderParams) GetDestPort() int {
	return *p.DestPort
}

func (p *IngressSenderParams) GetLabelFilter() string {
	return *p.LabelFilter
}

func (p *IngressSenderParams) GetLabelFilterName() string {
	return strings.Split(p.GetLabelFilter(), "=")[0]
}

func (p *IngressSenderParams) GetLabelFilterValue() string {
	ret := strings.Split(p.GetLabelFilter(), "=")
	if len(ret) < 2 {
		return ""
	} else {
		return ret[1]
	}
}

// Check if current label is present in a given map
func (p *IngressSenderParams) In(labels map[string]string) bool {
	lName := p.GetLabelFilterName()
	lValue := p.GetLabelFilterValue()

	for key, val := range labels {
		if key == lName && val == lValue {
			return true
		}
	}

	return false
}
