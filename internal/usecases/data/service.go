package data

import (
	"sync"

	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/entities"
)

type service struct {
	dataRepository DataRepository
}

func NewService(r DataRepository) DataUseCase {
	return &service{
		dataRepository: r,
	}
}

func (s *service) CreateData(value string) error {
	d, err := entities.NewData(value)
	if err != nil {
		return err
	}
	err = s.dataRepository.Create(d)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateDatas(values [][]string, batchSize int) (int, error) {
	created := 0
	divided := chunks(values, batchSize)
	for _, batch := range divided {
		d, err := entities.NewDatas(batch)
		if err != nil {
			return created, err
		}
		err = s.dataRepository.CreateMany(d)
		if err != nil {
			return created, err
		}
		created += len(batch)
	}
	return created, nil
}

func (s *service) saveToDB(ch chan [][]string, created *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	for batch := range ch {
		d, err := entities.NewDatas(batch)
		if err != nil {
			return
		}
		err = s.dataRepository.CreateMany(d)
		if err != nil {
			return
		}
		mu.Lock()
		*created += len(batch)
		mu.Unlock()
	}
	wg.Done()
}

func (s *service) CreateDatasConcurrent(values [][]string, batchSize int, numWorker int) (int, error) {
	var mu sync.Mutex
	var wg sync.WaitGroup

	ch := make(chan [][]string)

	created := 0
	divided := chunks(values, batchSize)

	for t := 0; t < numWorker; t++ {
		wg.Add(1)
		go s.saveToDB(ch, &created, &mu, &wg)
	}

	for _, batch := range divided {
		ch <- batch
	}
	close(ch)
	wg.Wait()
	return created, nil
}

func (s *service) ListDatas() ([]*entities.Data, error) {
	return s.dataRepository.List()
}

func (s *service) GetData(id string) (*entities.Data, error) {
	return s.dataRepository.Get(id)
}

func chunks(xs [][]string, chunkSize int) [][][]string {
	if len(xs) == 0 {
		return nil
	}
	divided := make([][][]string, (len(xs)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(xs) - chunkSize
	for prev < till {
		next := prev + chunkSize
		divided[i] = xs[prev:next]
		prev = next
		i++
	}
	divided[i] = xs[prev:]
	return divided
}
