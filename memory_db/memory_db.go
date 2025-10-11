package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Row struct {
	id    string
	value map[string]string
	mu    sync.RWMutex
}

type Table struct {
	Rows    []*Row
	mu      sync.RWMutex
	indexes map[string]map[string][]*Row // key -> values -> []Rows?
}

func (t *Table) CreateIndex(key string) {
	// t.mu.Lock()
	// defer t.mu.Unlock()
	index := map[string][]*Row{}
	for _, row := range t.Rows {
		value, ok := row.value[key]
		if ok {
			index[value] = append(index[value], row)
		}
	}
	t.indexes[key] = index
}

func (t *Table) SyncIndex() {
	for key := range t.indexes {
		t.CreateIndex(key)
	}
}

type Database struct {
	tables map[string]*Table // tableName to Table
}

func NewDatabase() *Database {
	return &Database{
		tables: make(map[string]*Table, 0),
	}
}

func (d *Database) CreateTables(tableName string) bool {
	_, ok := d.tables[tableName]
	fmt.Println("ok", ok)
	if ok {
		fmt.Println("table already exist")
		return false
	}
	fmt.Println("here")
	d.tables[tableName] = &Table{
		Rows: make([]*Row, 0),
	}

	return true
}

func (d *Database) PrintTable(tableName string) {
	table, ok := d.tables[tableName]
	if !ok {
		fmt.Println("table does not exist")
		return
	}
	table.mu.RLock()
	defer table.mu.RUnlock()
	fmt.Println("printing table")
	fmt.Println(len(table.Rows))
	for _, row := range table.Rows {
		row.mu.RLock()
		fmt.Println("row id: ", row.id, "values : ", row.value)
		row.mu.RUnlock()
	}
}

func (d *Database) INSERT(tableName string, value map[string]string) bool {
	// check if the table exists
	table, ok := d.tables[tableName]
	if !ok {
		fmt.Println("table does not exist")
		return false
	}

	fmt.Println("appending")
	table.mu.Lock()
	defer table.mu.Unlock()
	table.Rows = append(table.Rows, &Row{
		id:    uuid.NewString(),
		value: value,
	})

	table.SyncIndex()

	fmt.Println("appended")
	return true
}

func (d *Database) SELECT(tableName string, key string, value string) []*Row {
	table, ok := d.tables[tableName]
	if !ok {
		fmt.Println("table does not exist")
		return []*Row{}
	}

	var ans []*Row
	table.mu.RLock()
	defer table.mu.RUnlock()

	index, ok := table.indexes[key]
	if ok {
		return index[value]
	}

	for _, row := range table.Rows {
		tableData := row.value
		val, ok := tableData[key]
		if !ok {
			continue
		}
		if val == value {
			ans = append(ans, row)
		}
	}
	return ans
}

func (d *Database) UPDATE(tableName string, key string, value string, values map[string]string) bool {
	table, ok := d.tables[tableName]
	if !ok {
		fmt.Println("table does not exist")
		return false
	}
	table.mu.Lock()
	defer table.mu.Unlock()
	updated := false
	for _, row := range table.Rows {
		row.mu.Lock()
		tableData := row.value
		val, ok := tableData[key]
		if ok && val == value {
			updated = true
			row.value = values
		}
		row.mu.Unlock()
	}

	table.SyncIndex()
	return updated
}

func (d *Database) DELETE(tableName string, key string, value string) bool {
	table, ok := d.tables[tableName]
	if !ok {
		fmt.Println("table does not exist")
		return false
	}
	table.mu.Lock()
	defer table.mu.Unlock()

	deleted := false
	newRows := []*Row{}
	for _, row := range table.Rows {
		row.mu.Lock()
		val, ok := row.value[key]
		if !ok || val != value {
			newRows = append(newRows, row)
		} else {
			deleted = true
		}
		row.mu.Unlock()
	}
	table.Rows = newRows

	table.SyncIndex()
	return deleted
}

func main() {
	database := NewDatabase()
	database.CreateTables("users")
	database.INSERT("users", map[string]string{
		"name":  "avinash",
		"ph":    "fasdfasdf",
		"email": "avinash@gmail.com",
	})
	database.INSERT("users", map[string]string{
		"name":  "test",
		"ph":    "fasdfdf",
		"email": "avinash@gmail.com",
	})
	database.INSERT("users", map[string]string{
		"name":  "asdfasdf",
		"ph":    "asdfasdf",
		"email": "asdfasd@gmail.com",
	})

	database.INSERT("users", map[string]string{
		"name":  "asdfawerqwerasdf",
		"ph":    "asdfasdfasdfasdf",
		"email": "asdfasd@gmail.com",
	})

	database.UPDATE("users", "name", "test", map[string]string{
		"name":  "test1",
		"ph":    "fasdfdf",
		"email": "avinash@gmail.com",
	})

	database.tables["users"].CreateIndex("name")

	database.DELETE("users", "name", "test1")
	fmt.Println(database.SELECT("users", "ph", "fasdfdf"))

	database.PrintTable("users")
}
