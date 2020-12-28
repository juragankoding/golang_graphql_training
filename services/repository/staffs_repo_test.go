package repository

import (
	"database/sql"
	"testing"

	"github.com/icrowley/fake"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var listStaffs []*domain.Staffs

var idStores int64

func TestInsertStaffs(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing staffs on database : %s", err.Error())
	}

	staffsRepository := NewGenerateStaffsRepository(database)
	storesRepository := NewGenerateStoresRepository(database)

	idStores, err = storesRepository.Insert(domain.Stores{
		StoreName: fake.DomainName(),
		Phone:     fake.Phone(),
		Email:     fake.EmailAddress(),
		Street:    fake.Street(),
		City:      fake.City(),
		State:     fake.State(),
		ZipCode:   fake.Zip(),
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	for loopingStaffs := 0; loopingStaffs < 10; loopingStaffs++ {
		staffs := domain.Staffs{
			StoreID:   int(idStores),
			FirstName: fake.FirstName(),
			LastName:  fake.LastName(),
			Email:     fake.LastName(),
			Phone:     fake.Phone(),
			Active:    "1",
		}

		lastID, err := staffsRepository.Insert(staffs)

		if err != nil {
			t.Errorf("error on testing staffs : %s", err.Error())
		}

		staffs.StaffID = int(lastID)
		listStaffs = append(listStaffs, &staffs)

		staffsOnDatabase, err := staffsRepository.Single(int(lastID))

		if err != nil {
			t.Errorf("error on testing staffs : %s", err.Error())
		}

		if staffs.Compare(*staffsOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}

}

func TestUpdateStaffs(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing staffs on database : %s", err.Error())
	}

	staffsRepository := NewGenerateStaffsRepository(database)

	for staffs := range listStaffs {
		id := *&listStaffs[staffs].StaffID

		staffs := domain.Staffs{
			StaffID:   id,
			StoreID:   int(idStores),
			FirstName: fake.FirstName(),
			LastName:  fake.LastName(),
			Email:     fake.LastName(),
			Phone:     fake.Phone(),
			Active:    "1",
		}

		countEffect, err := staffsRepository.Update(staffs)

		if err != nil {
			t.Errorf("error on testing staffs : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		staffsOnDatabase, err := staffsRepository.Single(int(id))

		if err != nil {
			t.Errorf("error on testing staffs : %s", err.Error())
		}

		if staffs.Compare(*staffsOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}
}

func TestAllStaffs(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing staffs on database : %s", err.Error())
	}

	staffsRepository := NewGenerateStaffsRepository(database)

	staffsOnDatabase, err := staffsRepository.All()

	for staffs := range listStaffs {
		staff := *&listStaffs[staffs]
		for staffs := range staffsOnDatabase {
			if staff.Compare(*staffsOnDatabase[staffs]) {
				t.Errorf("Found filed not exists %d -> %s %s", staff.StaffID, staff.FirstName, staff.LastName)
			}
		}
	}
}

func TestSingleStaffs(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing staffs on database : %s", err.Error())
	}

	staffsRepository := NewGenerateStaffsRepository(database)

	for staffs := range listStaffs {
		id := *&listStaffs[staffs].StaffID
		staffsOnDatabase, err := staffsRepository.Single(int(id))

		if err != nil {
			t.Errorf("error on testing staffs : %s", err.Error())
		}

		if staffsOnDatabase == nil {
			t.Errorf("error field not exists : %s", err.Error())
		}
	}
}

func TestDeleteStaffs(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing staffs on database : %s", err.Error())
	}

	staffsRepository := NewGenerateStaffsRepository(database)
	storesRepository := NewGenerateStoresRepository(database)

	for staffs := range listStaffs {
		id := *&listStaffs[staffs].StaffID
		countEffect, err := staffsRepository.Delete(id)

		if err != nil {
			t.Errorf("error on testing staffs : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		staffsOnDatabase, err := staffsRepository.Single(int(id))

		if err != sql.ErrNoRows {
			t.Errorf("error on testing staffs : %s", err.Error())
		}

		if staffsOnDatabase != nil {
			t.Errorf("error field still exists : %s", err.Error())
		}
	}

	if countEffect, err := storesRepository.Delete(int(idStores)); err != nil {
		t.Error(err.Error())
	} else if countEffect == 0 {
		t.Error("nothing store delete on delete staffs testing")
	}

}
