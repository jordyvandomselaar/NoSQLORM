package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Db struct {
	DatabasePath string
	Tables       map[string]*Table
}

// Creates a dir to store databases if it doesn't exist, parses json files as tables.
func (db *Db) Open(path string) {
	db.DatabasePath = path

	CreateDir(path)
	infoSlice, _ := ioutil.ReadDir(path)

	tables := make(map[string]*Table)

	for _, info := range infoSlice {
		name := strings.Replace(info.Name(), ".json", "", -1)

		tables[name] = NewTable(filepath.Join(path, info.Name()))
	}

	db.Tables = tables
}

func NewDb() *Db {
	return &Db{}
}
