package data_repository

import (
	"errors"

	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/entities"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/pkg/id"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/usecases/data"
	"gorm.io/gorm"
)

type DataGORM struct {
	gorm.Model
	ID    id.ID `gorm:"type:uuid"`
	Value string
}

// Set tablename (GORM)
func (DataGORM) TableName() string {
	return "data"
}

func (u DataGORM) toEntitiesData() *entities.Data {
	return &entities.Data{
		ID:    u.ID,
		Value: u.Value,
	}
}

func NewDataGORM(entityData *entities.Data) *DataGORM {
	d := DataGORM{}
	d.ID = entityData.ID
	d.Value = entityData.Value
	return &d
}

type repository struct {
	DB *gorm.DB
}

func NewDataGORMRepository(db *gorm.DB) data.DataRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(entityData *entities.Data) error {
	d := NewDataGORM(entityData)

	err := r.DB.Create(&d).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateMany(entityDatas *[]entities.Data) error {
	datas := make([]*DataGORM, len(*entityDatas))
	for i, d := range *entityDatas {
		datas[i] = NewDataGORM(&d)
	}

	err := r.DB.Create(&datas).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) List() ([]*entities.Data, error) {
	var datas []DataGORM

	err := r.DB.Find(&datas).Error
	if err != nil {
		return nil, err
	}

	// TODO: Refactor. maybe inefficient.
	result := make([]*entities.Data, 0, len(datas))
	for _, data := range datas {
		result = append(result, data.toEntitiesData())
	}

	return result, nil
}

func (r *repository) Get(dataID string) (*entities.Data, error) {
	var data DataGORM

	r.DB.Find(&data, "id = ?", dataID)

	// If no such user present return an error
	if id.UUIDIsNil(data.ID) {
		return nil, errors.New("Data does not exists.")
	}

	return data.toEntitiesData(), nil
}
