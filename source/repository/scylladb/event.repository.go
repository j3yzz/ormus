package scylladb

import (
	"github.com/ormushq/ormus/source/adapter/scylladb"
	"github.com/ormushq/ormus/source/model"
)

type EventRepository interface {
	Save(event model.Event) (*model.Event, error)
	GetById(id string) (*model.Event, error)
}

type eventRepository struct {
	session scylladb.SessionxInterface
}

func NewEventRepository(session scylladb.SessionxInterface) EventRepository {
	return &eventRepository{session: session}
}

func (e *eventRepository) Save(event model.Event) (*model.Event, error) {
	var query string = "INSERT INTO ormus_source.events (id, key, value, created_at, updated_at)" +
		"VALUES (" +
		"uuid(), 'your_key_valu2e','your_value_valu2e', toTimestamp(now()), toTimestamp(now()))"

	insertQuery := e.session.Query(query, nil)

	if err := insertQuery.Exec(); err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *eventRepository) GetById(id string) (*model.Event, error) {
	return nil, nil
}
