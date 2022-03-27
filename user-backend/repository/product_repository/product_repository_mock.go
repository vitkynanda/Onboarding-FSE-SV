package product_repository

import (
	"go-api/models/dto"
	"go-api/models/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository ProductRepositoryMock) GetAllProducts() ([]entity.Product, error) {
	ret := repository.Mock.Called()

	var r0 []entity.Product
	if rf, ok := ret.Get(0).(func() []entity.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (repository ProductRepositoryMock) GetProductById(id string) (entity.Product, error) {
	ret := repository.Mock.Called(id)

	var r0 entity.Product
	if rf, ok := ret.Get(0).(func(string) entity.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (repository ProductRepositoryMock) UpdateProductData(id string) (entity.Product, error) {
	ret := repository.Mock.Called(id)  

	var r0 entity.Product
	if rf, ok := ret.Get(0).(func(string) entity.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	r0.ID = id
	return r0, r1
}

func (repository ProductRepositoryMock) DeleteProductById(id string) (entity.Product, error) {
	ret := repository.Mock.Called(id)  

	var r0 entity.Product
	if rf, ok := ret.Get(0).(func(string) entity.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	r0.ID = id
	return r0, r1
}


func (repository ProductRepositoryMock) CreateNewProduct(product dto.Product) (entity.Product, error) {
	ret := repository.Mock.Called()  

	var r0 = entity.Product{ID: product.ID, Name:product.Name, Description:product.Description}
	if rf, ok := ret.Get(0).(func() entity.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	
	return r0, r1
}