package repository

import (
	"database/sql"
	"testing"

	"github.com/icrowley/fake"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var listBrands []*domain.Brands

func TestInsertBrands(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing brands on database : %s", err.Error())
	}

	brandsRepository := NewGenerateBrandsRepository(database)

	for loopingBrands := 0; loopingBrands < 10; loopingBrands++ {
		brands := domain.Brands{
			Name: fake.Brand(),
		}

		lastID, err := brandsRepository.Insert(brands)

		if err != nil {
			t.Errorf("error on testing brands : %s", err.Error())
		}

		brands.ID = int(lastID)
		listBrands = append(listBrands, &brands)

		brandsOnDatabase, err := brandsRepository.Single(int(lastID))

		if err != nil {
			t.Errorf("error on testing brands : %s", err.Error())
		}

		if brands.Compare(*brandsOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}

}

func TestUpdateBrands(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing brands on database : %s", err.Error())
	}

	brandsRepository := NewGenerateBrandsRepository(database)

	for brands := range listBrands {
		id := *&listBrands[brands].ID

		brands1 := domain.Brands{
			ID:   id,
			Name: fake.Brand(),
		}

		countEffect, err := brandsRepository.Update(brands1)

		if err != nil {
			t.Errorf("error on testing brands : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		listBrands[brands] = &brands1

		brandsOnDatabase, err := brandsRepository.Single(int(id))

		if err != nil {
			t.Errorf("error on testing brands : %s", err.Error())
		}

		if brands1.Compare(*brandsOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}
}

func TestAllBrands(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing brands on database : %s", err.Error())
	}

	brandsRepository := NewGenerateBrandsRepository(database)

	brandsOnDatabase, err := brandsRepository.All()

	for brands := range listBrands {
		brand := *&listBrands[brands]

		var exists bool

		exists = false

		for brands := range brandsOnDatabase {
			exists = brand.Compare(*brandsOnDatabase[brands]) || exists
			if exists {
				break
			}
		}

		if exists == false {
			t.Errorf("Found filed not exists %d -> %s", brand.ID, brand.Name)
		}
	}
}

func TestSingleBrands(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing brands on database : %s", err.Error())
	}

	brandsRepository := NewGenerateBrandsRepository(database)

	for brands := range listBrands {
		id := *&listBrands[brands].ID
		brandsOnDatabase, err := brandsRepository.Single(int(id))

		if err != nil {
			t.Errorf("error on testing brands : %s", err.Error())
		}

		if brandsOnDatabase == nil {
			t.Errorf("error field not exists : %s", err.Error())
		}
	}
}

func TestDeleteBrands(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing brands on database : %s", err.Error())
	}

	brandsRepository := NewGenerateBrandsRepository(database)

	for brands := range listBrands {
		id := *&listBrands[brands].ID
		countEffect, err := brandsRepository.Delete(id)

		if err != nil {
			t.Errorf("error on testing brands : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		brandsOnDatabase, err := brandsRepository.Single(int(id))

		if err != sql.ErrNoRows {
			t.Errorf("error on testing brands : %s", err.Error())
		}

		if brandsOnDatabase != nil {
			t.Errorf("error field still exists : %s", err.Error())
		}
	}
}
