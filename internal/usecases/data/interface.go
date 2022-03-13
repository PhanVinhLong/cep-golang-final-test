package data

import "github.com/vinigracindo/fiber-gorm-clean-architecture/internal/entities"

type Reader interface {
	Get(dataID string) (*entities.Data, error)
	List() ([]*entities.Data, error)
}

type Writer interface {
	Create(u *entities.Data) error
	CreateMany(u *[]entities.Data) error
	//Update(u *entities.User) error
	//Delete(u *entities.User) error
}

type DataRepository interface {
	Reader
	Writer
}

type DataUseCase interface {
	CreateData(value string) error
	CreateDatas(values [][]string, batchSize int) (int, error)
	CreateDatasConcurrent(values [][]string, batchSize int, numWorker int) (int, error)
	ListDatas() ([]*entities.Data, error)
	GetData(id string) (*entities.Data, error)
}
