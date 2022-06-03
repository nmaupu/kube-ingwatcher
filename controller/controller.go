package controller

import (
	"fmt"
	"github.com/nmaupu/kube-ingwatcher/event"
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"log"
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

	client := clientset.CoreV1().RESTClient()
	watchlist := cache.NewListWatchFromClient(
		client,
		"ingresses",
		metav1.NamespaceAll,
		fields.Everything())

	_, controller := cache.NewInformer(
		watchlist,
		&networking.Ingress{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    func(obj interface{}) {
				err := getEvent().Add(obj)
				if err != nil {
					log.Printf("%v", err)
				}
			},
			DeleteFunc: func(obj interface{}) {
				err := getEvent().Delete(obj)
				if err != nil {
					log.Printf("%v", err)
				}
			},
			UpdateFunc: func(old, new interface{}) {
				err := getEvent().Update(old, new)
				if err != nil {
					fmt.Printf("%v", err)
				}
			},
		},
	)

	stopCh := make(chan struct{})
	go controller.Run(stopCh)
	<-stopCh
}
