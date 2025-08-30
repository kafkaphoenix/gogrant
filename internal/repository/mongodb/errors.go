package mongodb

import "errors"

var (
	ErrEmptyURI = errors.New("mongo URI is empty")
)

type MongoDBError struct {
	Message string
	Err     error
}

func (e *MongoDBError) Error() string {
	return e.Message + ": " + e.Err.Error()
}
func (e *MongoDBError) Unwrap() error {
	return e.Err
}
