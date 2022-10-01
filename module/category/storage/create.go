package categorystorage

import (
	"context"
	categorymodel "go-api/module/category/model"
)

func (s *sqlStorage) CreateCatgory(context context.Context, data *categorymodel.CategoryCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}