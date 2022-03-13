package entities

import (
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/pkg/id"
)

type Data struct {
	ID    id.ID  `valid:"required" json:"id"`
	Value string `json:"value"`
}

func NewData(value string) (*Data, error) {
	data := &Data{
		ID:    id.NewID(),
		Value: value,
	}

	return data, nil
}

func NewDatas(values [][]string) (*[]Data, error) {
	datas := make([]Data, len(values))
	for i, value := range values {
		datas[i] = Data{
			ID:    id.NewID(),
			Value: value[0],
		}
	}

	return &datas, nil
}
