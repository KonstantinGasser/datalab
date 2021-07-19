package cassandra

import (
	"context"

	"github.com/gocql/gocql"
)

type Client struct {
	session *gocql.Session
}

func WithKeySpace(ks string) func(*gocql.ClusterConfig) {
	return func(c *gocql.ClusterConfig) {
		c.Keyspace = ks
	}
}

func WithConsistency(consistency gocql.Consistency) func(*gocql.ClusterConfig) {
	return func(c *gocql.ClusterConfig) {
		c.Consistency = consistency
	}
}

func New(host string, port int, opts ...func(*gocql.ClusterConfig)) (*Client, error) {
	cluster := gocql.NewCluster(host)
	cluster.Port = port

	for _, opt := range opts {
		opt(cluster)
	}

	s, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &Client{
		session: s,
	}, nil
}

func (c Client) Close() {
	c.session.Close()
}

const (
	InsertFunnelChange = `
		INSERT INTO funnel_count_by_user(
			stage,
			app,
			user,
			action,
			leaving,
			elapsed,
			timestamp
		) values(?,?,?,?,?,?,?);
	`
)

func (c Client) InsertEvent(ctx context.Context, query string, args ...interface{}) error {
	return c.session.Query(query, args...).Exec()
}
