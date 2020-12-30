package repository

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/icrowley/fake"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var listProducts []*domain.Products
var idBrandsProduct int64
var idCategoryProduct int64

func TestInsertProducts(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing products on database : %s", err.Error())
	}

	productsRepository := NewGenerateProductsRepository(database)
	brandsRepository := NewGenerateBrandsRepository(database)
	categoriesRepository := NewGenerateCategoriesRepository(database)

	idBrandsProduct, err = brandsRepository.Insert(domain.Brands{
		Name: fake.Brand(),
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	idCategoryProduct, err = categoriesRepository.Insert(&domain.Categories{
		Name: fake.Brand(),
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	for loopingProducts := 0; loopingProducts < 10; loopingProducts++ {
		products := domain.Products{
			ListPrice:   1000,
			ModelYear:   strconv.Itoa((fake.Year(1990, 2000))),
			ProductName: fake.ProductName(),
			CategoryID:  int(idCategoryProduct),
			BrandID:     int(idBrandsProduct),
		}

		lastID, err := productsRepository.Insert(products)

		if err != nil {
			t.Errorf("error on testing products : %s", err.Error())
		}

		products.ProductID = int(lastID)
		listProducts = append(listProducts, &products)

		productsOnDatabase, err := productsRepository.Get(int(lastID))

		if err != nil {
			t.Errorf("error on testing products : %s", err.Error())
		}

		if products.Compare(*productsOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}

}

func TestUpdateProducts(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing products on database : %s", err.Error())
	}

	productsRepository := NewGenerateProductsRepository(database)

	for products := range listProducts {
		id := *&listProducts[products].ProductID

		products1 := domain.Products{
			ProductID:   id,
			ListPrice:   1000,
			ModelYear:   strconv.Itoa((fake.Year(1990, 2000))),
			ProductName: fake.ProductName(),
			CategoryID:  int(idCategoryProduct),
			BrandID:     int(idBrandsProduct),
		}
		countEffect, err := productsRepository.Update(products1)

		if err != nil {
			t.Errorf("error on testing products : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		listProducts[products] = &products1

		productsOnDatabase, err := productsRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing products : %s", err.Error())
		}

		if products1.Compare(*productsOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}
}

func TestAllProducts(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing products on database : %s", err.Error())
	}

	productsRepository := NewGenerateProductsRepository(database)

	productsOnDatabase, err := productsRepository.Fetch()

	for products := range listProducts {
		product := *&listProducts[products]

		var exists bool

		exists = false

		for products := range productsOnDatabase {
			exists = product.Compare(*productsOnDatabase[products]) || exists
		}

		if exists == false {
			t.Errorf("Found filed not exists %d -> %s", product.ProductID, product.ProductName)
		}
	}
}

func TestSingleProducts(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing products on database : %s", err.Error())
	}

	productsRepository := NewGenerateProductsRepository(database)

	for products := range listProducts {
		id := *&listProducts[products].ProductID
		productsOnDatabase, err := productsRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing products : %s", err.Error())
		}

		if productsOnDatabase == nil {
			t.Errorf("error field not exists : %s", err.Error())
		}
	}
}

func TestDeleteProducts(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing products on database : %s", err.Error())
	}

	productsRepository := NewGenerateProductsRepository(database)
	brandsRepository := NewGenerateBrandsRepository(database)
	categoriesRepository := NewGenerateCategoriesRepository(database)

	for products := range listProducts {
		id := *&listProducts[products].ProductID

		countEffect, err := productsRepository.Delete(id)

		if err != nil {
			t.Errorf("error on testing products : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		productsOnDatabase, err := productsRepository.Get(int(id))

		if err != sql.ErrNoRows {
			t.Errorf("error on testing products : %s", err.Error())
		}

		if productsOnDatabase != nil {
			t.Errorf("error field still exists : %s", err.Error())
		}
	}

	if _, err := brandsRepository.Delete(int(idBrandsProduct)); err != nil {
		t.Fatal(err.Error())
	}
	if _, err := categoriesRepository.Delete(int(idCategoryProduct)); err != nil {
		t.Fatal(err.Error())
	}
}
