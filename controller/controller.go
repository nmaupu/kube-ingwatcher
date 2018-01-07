package controller

import (
	"github.com/nmaupu/kube-ingwatcher/event"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"time"
)

func InitIngressController(getEvent func() event.Event) {

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	client := clientset.Extensions().RESTClient()
	watchlist := cache.NewListWatchFromClient(
		client,
		"ingresses",
		v1.NamespaceAll,
		fields.Everything())

	_, controller := cache.NewInformer(
		watchlist,
		&v1beta1.Ingress{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    getEvent().Add,
			DeleteFunc: getEvent().Delete,
			UpdateFunc: getEvent().Update,
		},
	)

	stopCh := make(chan struct{})
	go controller.Run(stopCh)
	<-stopCh
}
