package event

import (
	"encoding/json"
	"fmt"
)

var (
	_ Event = JsonPrinter{}
)

type JsonPrinter struct {
}

func (e JsonPrinter) Add(obj interface{}) {
	jsn, err := processJson(obj)
	if err != nil {
		fmt.Errorf("Ops! Cannot marshal JSON")
		return
	}

	fmt.Printf("Add: %s\n", jsn)
}

func (e JsonPrinter) Delete(obj interface{}) {
	jsn, err := processJson(obj)
	if err != nil {
		fmt.Errorf("Ops! Cannot marshal JSON")
		return
	}

	fmt.Printf("Delete: %s\n", jsn)
}

func (e JsonPrinter) Update(oldObj, newObj interface{}) {
	var jsnOld, jsnNew []byte
	var err error

	jsnOld, err = processJson(oldObj)
	if err != nil {
		fmt.Errorf("Ops! Cannot marshal JSON")
		return
	}

	jsnNew, err = processJson(newObj)
	if err != nil {
		fmt.Errorf("Ops! Cannot marshal JSON")
		return
	}

	fmt.Printf("Update: \n")
	fmt.Printf("  Old: %s\n", jsnOld)
	fmt.Printf("  New: %s\n", jsnNew)
}

func processJson(obj interface{}) ([]byte, error) {
	var jsn []byte
	var err error

	if jsn, err = json.Marshal(obj); err != nil {
		return nil, err
	}

	return jsn, nil
}
