package database

import (
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/infrastructure/gorm/repository"
	data_repository "github.com/vinigracindo/fiber-gorm-clean-architecture/internal/infrastructure/gorm/repository/data"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectGORMDB(dialector gorm.Dialector) (*gorm.DB, error) {
	var gormDB *gorm.DB
	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	gormDB.AutoMigrate(&repository.UserGORM{}, &data_repository.DataGORM{})
	return gormDB, nil
}
