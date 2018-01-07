package event

import (
	"fmt"
)

var (
	_ Event = Printer{}
)

type Printer struct {
}

func (e Printer) Add(obj interface{}) {
	fmt.Printf("Add: %s\n", obj)
}

func (e Printer) Delete(obj interface{}) {
	fmt.Printf("Delete: %s\n", obj)
}

func (e Printer) Update(oldObj, newObj interface{}) {
	fmt.Printf("Update: \n")
	fmt.Printf("  Old: %s\n", oldObj)
	fmt.Printf("  New: %s\n", newObj)
}
