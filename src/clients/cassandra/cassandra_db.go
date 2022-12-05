package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	cluster = gocql.NewCluster("172.17.0.3")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum // gocql.One // gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
