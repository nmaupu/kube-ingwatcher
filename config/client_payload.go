package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	ACTION_ADD    = 0
	ACTION_DELETE = 1
)

type ClientPayload struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Hosts     []string          `json:"hosts"`
	Labels    map[string]string `json:"labels"`
}

func (cp ClientPayload) GenerateTemplate(tpl, dest string) error {
	funcMap := template.FuncMap{
		"Join": strings.Join,
	}

	t, err := template.New("").Funcs(funcMap).ParseFiles(tpl)
	if err != nil {
		return err
	}

	file, _ := os.Create(dest)
	defer file.Close()
	w := bufio.NewWriter(file)

	err = t.ExecuteTemplate(w, filepath.Base(tpl), cp)
	w.Flush()

	return err
}
