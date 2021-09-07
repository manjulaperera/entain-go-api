package service

import (
	"git.neds.sh/matty/entain/sports/db"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"golang.org/x/net/context"
)

type Sports interface {
	// ListEvents will return a collection of sports.
	ListEvents(ctx context.Context, in *sports.ListEventsRequest) (*sports.ListEventsResponse, error)

	// Get sport details by id
	GetSportById(ctx context.Context, in *sports.GetSportRequest) (*sports.GetSportResponse, error)
}

// sportingService implements the Sports interface.
type sportingService struct {
	sportsRepo db.SportsRepo
}

// NewSportsService instantiates and returns a new sportingService.
func NewSportsService(sportsRepo db.SportsRepo) Sports {
	return &sportingService{sportsRepo}
}

// Get a list of sports with filter and order by clauses
func (s *sportingService) ListEvents(ctx context.Context, in *sports.ListEventsRequest) (*sports.ListEventsResponse, error) {
	sportEvents, err := s.sportsRepo.List(in.Filter, in.OrderBy)
	if err != nil {
		return nil, err
	}

	return &sports.ListEventsResponse{Sports: sportEvents}, nil
}

// Get sport details by id
func (s *sportingService) GetSportById(ctx context.Context, in *sports.GetSportRequest) (*sports.GetSportResponse, error) {
	sport, err := s.sportsRepo.GetSportById(in.Id)
	if err != nil {
		return nil, err
	}

	return &sports.GetSportResponse{Sport: sport}, nil
}
