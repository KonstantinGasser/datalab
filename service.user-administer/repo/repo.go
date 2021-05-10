package repo

import "context"

type Repo interface {
	// InsertOne appends the database with the given data
	InsertOne(ctx context.Context, db, collection string, query interface{}) error
	// UpdateOne updates on document in the database with the given data and based on the given filter
	// returns the number of updated documents and an error. Accepts mongo.UpdateOption.Upsert to upsert documents
	UpdateOne(ctx context.Context, db, collection string, filter, query interface{}, upsert bool) (int, error)
	// FindOne selects zero or one match from the filter an assigns the data in the given pointer to results
	FindOne(ctx context.Context, db, collection string, filter, result interface{}) error
	// FindMany select zero or many matches from the filter an assigns the data in the given pointer to results
	FindMany(ctx context.Context, db, collection string, filter, results interface{}) error
	// Exists checks whether something exists in the storage based on the filter
	Exists(ctx context.Context, db, collection string, filter interface{}) (bool, error)
	// Delete deletes a document based on a given filter
	DeleteOne(ctx context.Context, db, collection string, filter interface{}) error
}
