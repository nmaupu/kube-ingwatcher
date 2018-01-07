package event

type Event interface {
	Add(obj interface{})
	Delete(obj interface{})
	Update(oldObj, newObj interface{})
}
