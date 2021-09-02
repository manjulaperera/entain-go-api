package db

import (
	"database/sql"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"

	"git.neds.sh/matty/entain/sports/proto/sports"
)

// SportsRepo provides repository access to sports.
type SportsRepo interface {
	// Init will initialise our sports repository.
	Init() error

	// List will return a list of sports.
	List(filter *sports.ListEventsRequestFilter, order_by *sports.ListEventsRequestOrderBy) ([]*sports.Sport, error)

	// Get sport details by id
	GetSportById(sportId int64) (*sports.Sport, error)
}

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportsRepo creates a new sports repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the sport repository dummy data.
func (r *sportsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy sports.
		err = r.seed()
	})

	return err
}

// This is used in ListEvents method to return the sports events based on the filter and the order by clause
func (r *sportsRepo) List(filter *sports.ListEventsRequestFilter, orderBy *sports.ListEventsRequestOrderBy) ([]*sports.Sport, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportQueries()[sportsList]

	query, args = r.applyFilter(query, filter)
	query = r.applyOrderByClause(query, orderBy)

	rows, err := r.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return r.scanSportsRows(rows)
}

// Get a single sport by id
func (r *sportsRepo) GetSportById(sportId int64) (*sports.Sport, error) {
	var (
		query string
		args  []interface{}
	)

	query = getSportQueries()[sportById]

	args = append(args, sportId)

	row := r.db.QueryRow(query, args...)

	return r.scanSportsRow(row)
}

func (r *sportsRepo) applyFilter(query string, filter *sports.ListEventsRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		}
	}

	if filter.MeetingVisibility != nil {
		clauses = append(clauses, "visible = "+strconv.FormatBool(*filter.MeetingVisibility))
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

func (m *sportsRepo) scanSportsRows(
	rows *sql.Rows,
) ([]*sports.Sport, error) {
	var sportEvents []*sports.Sport

	for rows.Next() {
		var sport sports.Sport
		var advertisedStart time.Time

		if err := rows.Scan(&sport.Id, &sport.MeetingId, &sport.Name, &sport.Number, &sport.Visible, &sport.HomeTeam, &sport.AwayTeam, &advertisedStart, &sport.BettingClosedTime); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		sport.AdvertisedStartTime = ts

		// This will set Status based on the AdvertisedStartTime. Fake data inserted are not in UTC so we should check the local time.
		if advertisedStart.After(time.Now()) || advertisedStart.Equal(time.Now()) {
			sport.Status = "OPEN"
		} else {
			sport.Status = "CLOSED"
		}

		sportEvents = append(sportEvents, &sport)
	}

	return sportEvents, nil
}

// This will try to read the record from the DB and if successful it'll also set the sport status. Otherwise it'll return an error
func (m *sportsRepo) scanSportsRow(
	row *sql.Row,
) (*sports.Sport, error) {
	var sport sports.Sport
	var advertisedStart time.Time

	if err := row.Scan(&sport.Id, &sport.MeetingId, &sport.Name, &sport.Number, &sport.Visible, &sport.HomeTeam, &sport.AwayTeam, &advertisedStart, &sport.BettingClosedTime); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	ts, err := ptypes.TimestampProto(advertisedStart)
	if err != nil {
		return nil, err
	}

	sport.AdvertisedStartTime = ts

	// This will set Status based on the AdvertisedStartTime. Fake data inserted are not in UTC so we should check the local time.
	if advertisedStart.After(time.Now()) || advertisedStart.Equal(time.Now()) {
		sport.Status = "OPEN"
	} else {
		sport.Status = "CLOSED"
	}

	return &sport, nil
}

/* This will add an ORDER BY clause to the ListEvents query with the fileds specified in the request and their order by direction

NOTE: It's better to return an error if the client sends an invalid column name in the order by expression.
		Therefore the validity of the column name is not checked while adding the order by clause.
		And query execution will fail if atleast one of the order by expressions have an invalid column name.
*/
func (r *sportsRepo) applyOrderByClause(query string, orderBy *sports.ListEventsRequestOrderBy) string {
	var expressions []string

	if orderBy == nil {
		return query
	}

	if len(orderBy.OrderByFields) > 0 {
		for _, orderByField := range orderBy.OrderByFields {
			expressions = append(expressions, orderByField.Field+" "+orderByField.Direction.String())
		}

		query += " ORDER BY " + strings.Join(expressions, ", ")
	}

	return query
}
