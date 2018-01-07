package event

import (
	"github.com/nmaupu/kube-ingwatcher/config"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

var (
	_ Event = Sender{}
)

type Sender struct {
	Callback func(action int, payload config.ClientPayload)
}

func (e Sender) action(action int, obj interface{}) {
	ing, _ := obj.(*v1beta1.Ingress)

	// Getting all hosts from ingress
	var hosts []string
	for _, rule := range ing.Spec.Rules {
		hosts = append(hosts, rule.Host)
	}

	e.Callback(action, config.ClientPayload{
		Namespace: ing.Namespace,
		Name:      ing.Name,
		Hosts:     hosts,
	})
}

func (e Sender) Add(obj interface{}) {
	e.action(config.ACTION_ADD, obj)
}

func (e Sender) Delete(obj interface{}) {
	e.action(config.ACTION_DELETE, obj)
}

func (e Sender) Update(oldObj, newObj interface{}) {
	e.Delete(oldObj)
	e.Add(newObj)
}
