package gorputil

import (
	"reflect"

	"gopkg.in/gorp.v2"
)

type TypeRegister interface {
	AddTable(i interface{}) *gorp.TableMap
	AddTableWithName(i interface{}, name string) *gorp.TableMap
	AddTableWithNameAndSchema(i interface{}, schema string, name string) *gorp.TableMap
}

type Typer struct {
	typeRegister TypeRegister
	toType       map[string]reflect.Type
}

func NewTyper(typeRegister TypeRegister) *Typer {
	return &Typer{
		typeRegister: typeRegister,
		toType:       map[string]reflect.Type{},
	}
}

func (t *Typer) sniff(i interface{}, tableMap *gorp.TableMap) *gorp.TableMap {
	gotype := reflect.TypeOf(i)
	t.toType[tableMap.TableName] = gotype
	return tableMap
}

func (t *Typer) Lookup(tableName string) (reflect.Type, bool) {
	gotype, ok := t.toType[tableName]
	return gotype, ok
}

func (t *Typer) AddTable(i interface{}) *gorp.TableMap {
	return t.sniff(i, t.typeRegister.AddTable(i))
}

func (t *Typer) AddTableWithName(i interface{}, name string) *gorp.TableMap {
	return t.sniff(i, t.typeRegister.AddTableWithName(i, name))
}

func (t *Typer) AddTableWithNameAndSchema(i interface{}, schema string, name string) *gorp.TableMap {
	return t.sniff(i, t.typeRegister.AddTableWithNameAndSchema(i, schema, name))
}
