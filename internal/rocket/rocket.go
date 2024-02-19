//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/anil1226/go-gRPC/internal/rocket Store

package rocket

type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

type Service struct {
	Store Store
}

type Store interface {
	GetRocketByID(string) (Rocket, error)
	InsertRocket(Rocket) error
	DeleteRocket(string) error
}

func New(store Store) Service {
	return Service{
		Store: store,
	}
}

func (s Service) GetRocketByID(id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

func (s Service) InsertRocket(rkt Rocket) error {
	err := s.Store.InsertRocket(rkt)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
