// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/juragankoding/golang_graphql_training/domain"
	"github.com/juragankoding/golang_graphql_training/models"
)

type ResultInsert interface {
	IsResultInsert()
}

type ResultAllBrands struct {
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Data   []*domain.Brands `json:"data"`
}

func (ResultAllBrands) IsResultInsert() {}

type ResultAllProducts struct {
	Status string             `json:"status"`
	Code   int                `json:"code"`
	Data   []*domain.Products `json:"data"`
}

func (ResultAllProducts) IsResultInsert() {}

type ResultAllStores struct {
	Status string    `json:"status"`
	Code   int       `json:"code"`
	Data   []*Stores `json:"data"`
}

func (ResultAllStores) IsResultInsert() {}

type ResultDeleteBrands struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func (ResultDeleteBrands) IsResultInsert() {}

type ResultDeleteCategories struct {
	Status string             `json:"status"`
	Code   int                `json:"code"`
	Data   *domain.Categories `json:"data"`
}

func (ResultDeleteCategories) IsResultInsert() {}

type ResultDeleteProducts struct {
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Data   *domain.Products `json:"data"`
}

func (ResultDeleteProducts) IsResultInsert() {}

type ResultDeleteStaffs struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Staffs `json:"data"`
}

func (ResultDeleteStaffs) IsResultInsert() {}

type ResultDeleteStocks struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Stocks `json:"data"`
}

func (ResultDeleteStocks) IsResultInsert() {}

type ResultDeleteStores struct {
	Status string  `json:"status"`
	Code   int     `json:"code"`
	Data   *Stores `json:"data"`
}

func (ResultDeleteStores) IsResultInsert() {}

type ResultFetchCategories struct {
	Status string               `json:"status"`
	Code   int                  `json:"code"`
	Data   []*domain.Categories `json:"data"`
}

func (ResultFetchCategories) IsResultInsert() {}

type ResultFetchStaffs struct {
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Data   []*domain.Staffs `json:"data"`
}

func (ResultFetchStaffs) IsResultInsert() {}

type ResultFetchStocks struct {
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Data   []*domain.Stocks `json:"data"`
}

func (ResultFetchStocks) IsResultInsert() {}

type ResultGetCategories struct {
	Status string             `json:"status"`
	Code   int                `json:"code"`
	Data   *domain.Categories `json:"data"`
}

func (ResultGetCategories) IsResultInsert() {}

type ResultGetStaffs struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Staffs `json:"data"`
}

func (ResultGetStaffs) IsResultInsert() {}

type ResultGetStocks struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Stocks `json:"data"`
}

func (ResultGetStocks) IsResultInsert() {}

type ResultInsertBrands struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Brands `json:"data"`
}

func (ResultInsertBrands) IsResultInsert() {}

type ResultInsertCategories struct {
	Status string             `json:"status"`
	Code   int                `json:"code"`
	Data   *domain.Categories `json:"data"`
}

func (ResultInsertCategories) IsResultInsert() {}

type ResultInsertProducts struct {
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Data   *domain.Products `json:"data"`
}

func (ResultInsertProducts) IsResultInsert() {}

type ResultInsertStaffs struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Staffs `json:"data"`
}

func (ResultInsertStaffs) IsResultInsert() {}

type ResultInsertStocks struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Stocks `json:"data"`
}

func (ResultInsertStocks) IsResultInsert() {}

type ResultInsertStores struct {
	Status string  `json:"status"`
	Code   int     `json:"code"`
	Data   *Stores `json:"data"`
}

func (ResultInsertStores) IsResultInsert() {}

type ResultSingleBrands struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Brands `json:"data"`
}

func (ResultSingleBrands) IsResultInsert() {}

type ResultSingleProducts struct {
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Data   *domain.Products `json:"data"`
}

func (ResultSingleProducts) IsResultInsert() {}

type ResultSingleStores struct {
	Status string  `json:"status"`
	Code   int     `json:"code"`
	Data   *Stores `json:"data"`
}

func (ResultSingleStores) IsResultInsert() {}

type ResultUpdateBrands struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Brands `json:"data"`
}

func (ResultUpdateBrands) IsResultInsert() {}

type ResultUpdateCategories struct {
	Status string             `json:"status"`
	Code   int                `json:"code"`
	Data   *domain.Categories `json:"data"`
}

func (ResultUpdateCategories) IsResultInsert() {}

type ResultUpdateProducts struct {
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Data   *domain.Products `json:"data"`
}

func (ResultUpdateProducts) IsResultInsert() {}

type ResultUpdateStaffs struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Staffs `json:"data"`
}

func (ResultUpdateStaffs) IsResultInsert() {}

type ResultUpdateStocks struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   *domain.Stocks `json:"data"`
}

func (ResultUpdateStocks) IsResultInsert() {}

type ResultUpdateStores struct {
	Status string  `json:"status"`
	Code   int     `json:"code"`
	Data   *Stores `json:"data"`
}

func (ResultUpdateStores) IsResultInsert() {}

type Stores struct {
	StoreID   int    `json:"StoreID"`
	StoreName string `json:"StoreName"`
	Phone     string `json:"Phone"`
	Email     string `json:"Email"`
	City      string `json:"City"`
	State     string `json:"State"`
	ZipCode   string `json:"ZipCode"`
}

type ResultGetAllJenisBarang struct {
	Status string                `json:"status"`
	Code   int                   `json:"code"`
	Data   []*models.JenisBarang `json:"data"`
}

func (ResultGetAllJenisBarang) IsResultInsert() {}

type ResultJenisBarang struct {
	Status string              `json:"status"`
	Code   int                 `json:"code"`
	Data   *models.JenisBarang `json:"data"`
}

func (ResultJenisBarang) IsResultInsert() {}
