package event

import (
	"encoding/json"
	"fmt"
)

var (
	_ Event = JsonPrinter{}
)

type JsonPrinter struct {}

func (e JsonPrinter) Add(obj interface{}) error {
	jsn, err := processJson(obj)
	if err != nil {
		return err
	}

	fmt.Printf("Add: %s\n", jsn)
	return nil
}

func (e JsonPrinter) Delete(obj interface{}) error {
	jsn, err := processJson(obj)
	if err != nil {
		return err
	}

	fmt.Printf("Delete: %s\n", jsn)
	return nil
}

func (e JsonPrinter) Update(oldObj, newObj interface{}) error {
	var jsnOld, jsnNew []byte
	var err error

	jsnOld, err = processJson(oldObj)
	if err != nil {
		return err
	}

	jsnNew, err = processJson(newObj)
	if err != nil {
		return err
	}

	fmt.Printf("Update: \n")
	fmt.Printf("  Old: %s\n", jsnOld)
	fmt.Printf("  New: %s\n", jsnNew)
	return nil
}

func processJson(obj interface{}) ([]byte, error) {
	var jsn []byte
	var err error

	if jsn, err = json.Marshal(obj); err != nil {
		return nil, err
	}

	return jsn, nil
}
