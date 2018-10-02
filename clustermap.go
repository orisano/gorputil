package gorputil

import (
	"database/sql"

	"gopkg.in/gorp.v2"
)

type ClusterMap struct {
	*gorp.DbMap
	slaves   []*gorp.DbMap
	balancer Balancer
}

func NewClusterMap(master *gorp.DbMap, slaves []*gorp.DbMap, balancer Balancer) *ClusterMap {
	return &ClusterMap{
		DbMap:    master,
		slaves:   slaves,
		balancer: balancer,
	}
}

func (m *ClusterMap) applyCluster(f func(*gorp.DbMap) *gorp.TableMap) *gorp.TableMap {
	for _, slave := range m.slaves {
		f(slave)
	}
	return f(m.DbMap)
}

func (m *ClusterMap) AddTable(i interface{}) *gorp.TableMap {
	return m.applyCluster(func(dbMap *gorp.DbMap) *gorp.TableMap {
		return dbMap.AddTable(i)
	})
}

func (m *ClusterMap) AddTableWith(i interface{}, name string) *gorp.TableMap {
	return m.applyCluster(func(dbMap *gorp.DbMap) *gorp.TableMap {
		return dbMap.AddTableWithName(i, name)
	})
}

func (m *ClusterMap) AddTableWithNameAndSchema(i interface{}, schema string, name string) *gorp.TableMap {
	return m.applyCluster(func(dbMap *gorp.DbMap) *gorp.TableMap {
		return dbMap.AddTableWithNameAndSchema(i, schema, name)
	})
}

func (m *ClusterMap) AddTableDynamic(inp gorp.DynamicTable, schema string) *gorp.TableMap {
	return m.applyCluster(func(dbMap *gorp.DbMap) *gorp.TableMap {
		return dbMap.AddTableDynamic(inp, schema)
	})
}

func (m *ClusterMap) node() *gorp.DbMap {
	if len(m.slaves) == 0 {
		return m.DbMap
	}
	id := m.balancer.Balance(len(m.slaves))
	return m.slaves[id]
}

func (m *ClusterMap) Get(i interface{}, keys ...interface{}) (interface{}, error) {
	return m.node().Get(i, keys...)
}

func (m *ClusterMap) SelectInt(query string, args ...interface{}) (int64, error) {
	return m.node().SelectInt(query, args...)
}

func (m *ClusterMap) SelectNullInt(query string, args ...interface{}) (sql.NullInt64, error) {
	return m.node().SelectNullInt(query, args...)
}

func (m *ClusterMap) SelectFloat(query string, args ...interface{}) (float64, error) {
	return m.node().SelectFloat(query, args...)
}

func (m *ClusterMap) SelectNullFloat(query string, args ...interface{}) (sql.NullFloat64, error) {
	return m.node().SelectNullFloat(query, args...)
}

func (m *ClusterMap) SelectStr(query string, args ...interface{}) (string, error) {
	return m.node().SelectStr(query, args...)
}

func (m *ClusterMap) SelectNullStr(query string, args ...interface{}) (sql.NullString, error) {
	return m.node().SelectNullStr(query, args...)
}

func (m *ClusterMap) SelectOne(holder interface{}, query string, args ...interface{}) error {
	return m.node().SelectOne(holder, query, args...)
}

func (m *ClusterMap) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if isSelectQuery(query) {
		return m.node().Query(query, args...)
	}
	return m.DbMap.Query(query, args...)
}

func (m *ClusterMap) QueryRow(query string, args ...interface{}) *sql.Row {
	if isSelectQuery(query) {
		return m.node().QueryRow(query, args...)
	}
	return m.DbMap.QueryRow(query, args...)
}
