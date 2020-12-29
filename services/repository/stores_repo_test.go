package repository

import (
	"database/sql"
	"testing"

	"github.com/icrowley/fake"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var listStores []*domain.Stores

func TestInsertStores(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing stores on database : %s", err.Error())
	}

	storesRepository := NewGenerateStoresRepository(database)

	for loopingStores := 0; loopingStores < 10; loopingStores++ {
		stores := domain.Stores{
			StoreName: fake.FullName(),
			Phone:     fake.Phone(),
			Email:     fake.EmailAddress(),
			Street:    fake.Street(),
			City:      fake.City(),
			State:     fake.State(),
			ZipCode:   fake.Zip(),
		}

		lastID, err := storesRepository.Insert(stores)

		if err != nil {
			t.Errorf("error on testing stores : %s", err.Error())
		}

		stores.StoreID = int(lastID)
		listStores = append(listStores, &stores)

		storesOnDatabase, err := storesRepository.Get(int(lastID))

		if err != nil {
			t.Errorf("error on testing stores : %s", err.Error())
		}

		if stores.Compare(*storesOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}

}

func TestUpdateStores(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing stores on database : %s", err.Error())
	}

	storesRepository := NewGenerateStoresRepository(database)

	for stores := range listStores {
		id := *&listStores[stores].StoreID

		stores := domain.Stores{
			StoreID:   id,
			StoreName: fake.FullName(),
			Phone:     fake.Phone(),
			Email:     fake.EmailAddress(),
			Street:    fake.Street(),
			City:      fake.City(),
			State:     fake.State(),
			ZipCode:   fake.Zip(),
		}
		countEffect, err := storesRepository.Update(stores)

		if err != nil {
			t.Errorf("error on testing stores : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		storesOnDatabase, err := storesRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing stores : %s", err.Error())
		}

		if stores.Compare(*storesOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}
}

func TestAllStores(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing stores on database : %s", err.Error())
	}

	storesRepository := NewGenerateStoresRepository(database)

	storesOnDatabase, err := storesRepository.Fetch()

	for stores := range listStores {
		store := *&listStores[stores]
		var exists bool

		exists = false

		for stores := range storesOnDatabase {
			exists = store.Compare(*storesOnDatabase[stores]) || exists
		}

		if exists == false {
			t.Errorf("Found filed not exists %d -> %s", store.StoreID, store.StoreName)
		}
	}
}

func TestSingleStores(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing stores on database : %s", err.Error())
	}

	storesRepository := NewGenerateStoresRepository(database)

	for stores := range listStores {
		id := *&listStores[stores].StoreID
		storesOnDatabase, err := storesRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing stores : %s", err.Error())
		}

		if storesOnDatabase == nil {
			t.Errorf("error field not exists : %s", err.Error())
		}
	}
}

func TestDeleteStores(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing stores on database : %s", err.Error())
	}

	storesRepository := NewGenerateStoresRepository(database)

	for stores := range listStores {
		id := *&listStores[stores].StoreID

		countEffect, err := storesRepository.Delete(id)

		if err != nil {
			t.Errorf("error on testing stores : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		storesOnDatabase, err := storesRepository.Get(int(id))

		if err != sql.ErrNoRows {
			t.Errorf("error on testing stores : %s", err.Error())
		}

		if storesOnDatabase != nil {
			t.Errorf("error field still exists : %s", err.Error())
		}
	}
}
