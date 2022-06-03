package event

type Event interface {
	Add(obj interface{}) error
	Delete(obj interface{}) error
	Update(oldObj, newObj interface{}) error
}
