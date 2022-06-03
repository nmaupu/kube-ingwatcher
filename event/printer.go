package event

import (
	"fmt"
)

var (
	_ Event = Printer{}
)

type Printer struct {
}

func (e Printer) Add(obj interface{}) error {
	fmt.Printf("Add: %s\n", obj)
	return nil
}

func (e Printer) Delete(obj interface{}) error {
	fmt.Printf("Delete: %s\n", obj)
	return nil
}

func (e Printer) Update(oldObj, newObj interface{}) error {
	fmt.Printf("Update: \n")
	fmt.Printf("  Old: %s\n", oldObj)
	fmt.Printf("  New: %s\n", newObj)
	return nil
}
