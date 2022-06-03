package event

import (
	"github.com/nmaupu/kube-ingwatcher/config"
	networking "k8s.io/api/networking/v1"
)

var (
	_ Event = Sender{}
)

type Sender struct {
	Callback func(action int, payload config.ClientPayload)
}

func (e Sender) action(action int, obj interface{}) {
	ing, _ := obj.(*networking.Ingress)

	// Getting all hosts from ingress
	var hosts []string
	for _, rule := range ing.Spec.Rules {
		hosts = append(hosts, rule.Host)
	}

	e.Callback(action, config.ClientPayload{
		Namespace: ing.Namespace,
		Name:      ing.Name,
		Hosts:     hosts,
		Labels:    ing.Labels,
	})
}

func (e Sender) Add(obj interface{}) error {
	e.action(config.ActionAdd, obj)
	return nil
}

func (e Sender) Delete(obj interface{}) error {
	e.action(config.ActionDelete, obj)
	return nil
}

func (e Sender) Update(oldObj, newObj interface{}) error {
	e.Delete(oldObj)
	e.Add(newObj)
	return nil
}
