package repository

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/icrowley/fake"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var listStocks []*domain.Stocks
var idBrandsStocks int64
var idCategoryStocks int64
var idProductStocks int64
var idStoreStocks int64

func TestInsertStocks(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing stocks on database : %s", err.Error())
	}

	stocksRepository := NewGenerateStocksRepository(database)
	brandsRepository := NewGenerateBrandsRepository(database)
	categoriesRepository := NewGenerateCategoriesRepository(database)
	productsRepository := NewGenerateProductsRepository(database)
	storeRepository := NewGenerateStoresRepository(database)

	if idBrandsStocks, err = brandsRepository.Insert(domain.Brands{
		Name: fake.Brand(),
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idCategoryStocks, err = categoriesRepository.Insert(&domain.Categories{
		Name: fake.Brand(),
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idProductStocks, err = productsRepository.Insert(domain.Products{
		ListPrice:   1000,
		ModelYear:   strconv.Itoa((fake.Year(1990, 2000))),
		ProductName: fake.ProductName(),
		CategoryID:  int(idCategoryStocks),
		BrandID:     int(idBrandsStocks),
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idStoreStocks, err = storeRepository.Insert(domain.Stores{
		StoreName: fake.FullName(),
		Phone:     fake.Phone(),
		Email:     fake.EmailAddress(),
		Street:    fake.Street(),
		City:      fake.City(),
		State:     fake.State(),
		ZipCode:   fake.Zip(),
	}); err != nil {
		t.Error(err.Error())
	}

	for loopingStocks := 0; loopingStocks < 1; loopingStocks++ {
		stocks := domain.Stocks{
			StoreID:   int(idStoreStocks),
			ProductID: int(idProductStocks),
			Quantity:  100,
		}

		_, err := stocksRepository.Insert(stocks)

		if err != nil {
			t.Errorf("error on testing stocks : %s", err.Error())
		}

		listStocks = append(listStocks, &stocks)

		stocksOnDatabase, err := stocksRepository.Get(int(idStoreStocks), int(idProductStocks))

		if err != nil {
			t.Errorf("error on testing : %s", err.Error())
		}

		if stocks.Compare(*stocksOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}

}

func TestUpdateStocks(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing stocks on database : %s", err.Error())
	}

	stocksRepository := NewGenerateStocksRepository(database)

	for stocks := range listStocks {

		stocks1 := domain.Stocks{
			StoreID:   int(*&listStocks[stocks].StoreID),
			ProductID: int(*&listStocks[stocks].ProductID),
			Quantity:  110,
		}
		countEffect, err := stocksRepository.Update(stocks1)

		if err != nil {
			t.Errorf("error on testing stocks : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		listStocks[stocks] = &stocks1
		stocksOnDatabase, err := stocksRepository.Get(int(*&listStocks[stocks].StoreID), int(*&listStocks[stocks].ProductID))

		if err != nil {
			t.Errorf("error on testing stocks : %s", err.Error())
		}

		if stocks1.Compare(*stocksOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}
}

func TestAllStocks(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing stocks on database : %s", err.Error())
	}

	stocksRepository := NewGenerateStocksRepository(database)

	stocksOnDatabase, err := stocksRepository.Fetch()

	for stocks := range listStocks {
		stock := *&listStocks[stocks]

		var exists bool

		exists = false

		for stocks := range stocksOnDatabase {
			if exists = stock.Compare(*stocksOnDatabase[stocks]) || exists; exists {
				break
			}

		}

		if exists == false {
			t.Errorf("Found filed not exists %d -> %d", stock.ProductID, stock.StoreID)
		}
	}
}

func TestSingleStocks(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing stocks on database : %s", err.Error())
	}

	stocksRepository := NewGenerateStocksRepository(database)

	for stocks := range listStocks {
		stocksOnDatabase, err := stocksRepository.Get(int(*&listStocks[stocks].StoreID), int(*&listStocks[stocks].ProductID))

		if err != nil {
			t.Errorf("error on testing stocks : %s", err.Error())
		}

		if stocksOnDatabase == nil {
			t.Errorf("error field not exists : %s", err.Error())
		}
	}
}

func TestDeleteStocks(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing stocks on database : %s", err.Error())
	}

	stocksRepository := NewGenerateStocksRepository(database)
	brandsRepository := NewGenerateBrandsRepository(database)
	categoriesRepository := NewGenerateCategoriesRepository(database)

	for stocks := range listStocks {

		countEffect, err := stocksRepository.Delete(int(*&listStocks[stocks].StoreID), int(*&listStocks[stocks].ProductID))

		if err != nil {
			t.Errorf("error on testing stocks : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		stocksOnDatabase, err := stocksRepository.Get(int(idStoreOrders), int(idProductStocks))

		if err != sql.ErrNoRows {
			t.Errorf("error on testing stocks : %s", err.Error())
		}

		if stocksOnDatabase != nil {
			t.Errorf("error field still exists : %s", err.Error())
		}
	}

	if _, err := brandsRepository.Delete(int(idBrandsStocks)); err != nil {
		t.Fatal(err.Error())
	}

	if _, err := categoriesRepository.Delete(int(idCategoryStocks)); err != nil {
		t.Fatal(err.Error())
	}
}
