package repository

import (
	"database/sql"
	"testing"

	"github.com/icrowley/fake"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var listCategories []*domain.Categories

func TestInsertCategories(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing categories on database : %s", err.Error())
	}

	categoriesRepository := NewGenerateCategoriesRepository(database)

	for loopingCategories := 0; loopingCategories < 10; loopingCategories++ {
		categories := domain.Categories{
			Name: fake.Brand(),
		}

		lastID, err := categoriesRepository.Insert(&categories)

		if err != nil {
			t.Errorf("error on testing categories : %s", err.Error())
		}

		categories.ID = int(lastID)
		listCategories = append(listCategories, &categories)

		categoriesOnDatabase, err := categoriesRepository.Get(int(lastID))

		if err != nil {
			t.Errorf("error on testing categories : %s", err.Error())
		}

		if categories.Compare(*categoriesOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}

}

func TestUpdateCategories(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing categories on database : %s", err.Error())
	}

	categoriesRepository := NewGenerateCategoriesRepository(database)

	for categories := range listCategories {
		id := *&listCategories[categories].ID

		categories := domain.Categories{
			ID:   id,
			Name: fake.Brand(),
		}

		countEffect, err := categoriesRepository.Update(&categories)

		if err != nil {
			t.Errorf("error on testing categories : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		categoriesOnDatabase, err := categoriesRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing categories : %s", err.Error())
		}

		if categories.Compare(*categoriesOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}
}

func TestAllCategories(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing categories on database : %s", err.Error())
	}

	categoriesRepository := NewGenerateCategoriesRepository(database)

	categoriesOnDatabase, err := categoriesRepository.Fetch()

	for categories := range listCategories {
		brand := *&listCategories[categories]
		for categories := range categoriesOnDatabase {
			if brand.Compare(*categoriesOnDatabase[categories]) {
				t.Errorf("Found filed not exists %d -> %s", brand.ID, brand.Name)
			}
		}
	}
}

func TestSingleCategories(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing categories on database : %s", err.Error())
	}

	categoriesRepository := NewGenerateCategoriesRepository(database)

	for categories := range listCategories {
		id := *&listCategories[categories].ID
		categoriesOnDatabase, err := categoriesRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing categories : %s", err.Error())
		}

		if categoriesOnDatabase == nil {
			t.Errorf("error field not exists : %s", err.Error())
		}
	}
}

func TestDeleteCategories(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing categories on database : %s", err.Error())
	}

	categoriesRepository := NewGenerateCategoriesRepository(database)

	for categories := range listCategories {
		id := *&listCategories[categories].ID
		countEffect, err := categoriesRepository.Delete(id)

		if err != nil {
			t.Errorf("error on testing categories : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		categoriesOnDatabase, err := categoriesRepository.Get(int(id))

		if err != sql.ErrNoRows {
			t.Errorf("error on testing categories : %s", err.Error())
		}

		if categoriesOnDatabase != nil {
			t.Errorf("error field still exists : %s", err.Error())
		}
	}
}
