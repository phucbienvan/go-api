package categorybusiness

import (
	"context"
	"errors"
	categorymodel "go-api/module/category/model"
)

type CreateCategoryStore interface{
	CreateCatgory(context context.Context, data *categorymodel.CategoryCreate) error
}

type createCategoryBusinees struct {
	store CreateCategoryStore
}

func NewCreateCategoryBusiness(store CreateCategoryStore) *createCategoryBusinees{
	return &createCategoryBusinees{store: store}
}

func (business *createCategoryBusinees) CreateCatgory(context context.Context, data *categorymodel.CategoryCreate) error {
	if data.Name == "" {
		return errors.New("Name is required")
	}
	if err := business.store.CreateCatgory(context, data); err != nil {
		return err
	}

	return nil
}
