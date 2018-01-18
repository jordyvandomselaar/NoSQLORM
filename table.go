package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Table struct {
	Path         string
	Data         []byte
	Query        *Query
	SelectFields []string
}

func (t *Table) Get() map[string]map[string]string {
	t.Load()

	// If no wheres were added, return entire resultset.
	if t.Query == nil || t.Query.Filters == nil {
		return t.DataToMap()
	}

	// Store matches here
	result := make(map[string]map[string]string)

	// Loop through available rows
	for primaryKey, row := range t.DataToMap() {

		match := true

		// Loop through added filter statements
		for _, filter := range t.Query.Filters {

			// Run our filters on the dataset.
			if !filter.Check(row[filter.GetColumn()]) {
				match = false

				break
			}
		}

		if match {
			// Append the row to our resultset if it's what we want.
			result[primaryKey] = row
		}
	}

	return result
}

func (t *Table) Select(f []string) *Table {
	t.SelectFields = f

	return t
}

func (t *Table) AddQuery(q *Query) *Table {
	t.Query = q

	return t
}

func (t *Table) DataToMap() map[string]map[string]string {
	var result map[string]map[string]string

	err := json.Unmarshal(t.Data, &result)

	if err != nil {
		log.Fatal(err)
	}

	if t.SelectFields != nil {
		for key, row := range result {
			result[key] = GetFieldsFromMap(row, t.SelectFields)
		}
	}

	return result
}

func (t *Table) Load() {
	if t.Data != nil {
		return
	}

	tableJson, err := ioutil.ReadFile(t.Path)

	if err != nil {
		panic(err)
	}

	t.Data = tableJson
}

func NewTable(path string) *Table {
	return &Table{
		Path: path,
	}
}

type Query struct {
	Filters []FilterInterface
}

func NewQuery(filters ...FilterInterface) *Query {
	return &Query{
		Filters: filters,
	}
}
