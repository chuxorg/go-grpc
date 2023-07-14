package rocket

import "context"

// Defines a Rocket
type Rocket struct {
   
	ID string
	Name string
	Type string
	Flights int
}

// Defines a Store interface for database implementation
type Store interface {
	GetRocketById(id string) (*Rocket, error)
	InsertRocket(r *Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Rocket Services responsible for managing rockets
type Service struct {
	Store Store
}


func New (store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetRocketById returns a rocket by id
func (s *Service) GetRocketById(ctx context.Context, id string) (Rocket, error) {
	rocket, err := s.Store.GetRocketById(id)
	if err != nil {
		return Rocket{}, err
	}
	return *rocket, nil
}

// InsertRocket inserts a rocket into store
func (s *Service) InsertRocket(ctx context.Context, r *Rocket) (Rocket, error) {
	rocket, err := s.Store.InsertRocket(r)
	if err != nil {
		return Rocket{}, err
	}	
	return rocket, nil
}

// DeleteRocket deletes a rocket from store
func (s *Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
