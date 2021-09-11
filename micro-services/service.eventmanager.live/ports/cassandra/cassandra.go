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
		INSERT INTO users_per_stage_by_date(
			app,
			date,
			stage,
			user
		) values(?,?,?,?)
	`
	UpsertAvgTimeStage = `
		UPDATE avg_time_per_stage 
			SET elapsed = elapsed + ?, hit = hit+1
		WHERE app = ? and stage = ?
	`
	UpsertPageHit = `
		UPDATE page_view_per_url_by_date
			SET 
				hit = hit + 1,
				avg_duration = avg_duration + ?
		WHERE app = ? and url = ? and date = ?
	`

	UpsertInterestingBtn = `
			UPDATE action_count_interesting_btn
				SET 
					elapsed_click = elapsed_click + ?,
					elapsed_leave = elapsed_leave + ?,
					hover_click = hover_click + ?,
					hover_leave = hover_leave + ?,
					hit = hit + 1
			WHERE app = ? and xpath = ?
	`
	InsertSession = `
		INSERT INTO session(
			app,
			date,
			user,
			start,
			duration,
			referrer
		) values(?,?,?,?,?,?)
	`

	InsertBatchStart = `
		BEGIN COUNTER BATCH
			UPDATE used_device set count = count+1 where app = ? and device_type = ?;
			UPDATE used_browser set count = count+1 where app = ? and browser = ?;
			UPDATE views_per_date set views = views+1 where app = ? and date = ?;
		APPLY BATCH
	`
)

func (c Client) InsertEvent(ctx context.Context, query string, args ...interface{}) error {
	return c.session.Query(query, args...).Exec()
}
