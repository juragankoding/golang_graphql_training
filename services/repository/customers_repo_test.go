package repository

import (
	"database/sql"
	"testing"

	"github.com/icrowley/fake"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var listCustomers []*domain.Customers

func TestInsertCustomers(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing customers on database : %s", err.Error())
	}

	customersRepository := NewGenerateCustomersRepository(database)

	for loopingCustomers := 0; loopingCustomers < 10; loopingCustomers++ {
		customers := domain.Customers{
			FirstName: fake.FirstName(),
			LastName:  fake.LastName(),
			Phone:     fake.Phone(),
			Email:     fake.EmailAddress(),
			Street:    fake.Street(),
			City:      fake.City(),
			State:     sql.NullString{String: fake.State()},
			ZipCode:   111,
		}

		lastID, err := customersRepository.Insert(customers)

		if err != nil {
			t.Errorf("error on testing customers : %s", err.Error())
		}

		customers.CustomerID = int(lastID)
		listCustomers = append(listCustomers, &customers)

		customersOnDatabase, err := customersRepository.Get(int(lastID))

		if err != nil {
			t.Errorf("error on testing customers : %s", err.Error())
		}

		if customers.Compare(*customersOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}

}

func TestUpdateCustomers(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing customers on database : %s", err.Error())
	}

	customersRepository := NewGenerateCustomersRepository(database)

	for customers := range listCustomers {
		id := *&listCustomers[customers].CustomerID

		customers1 := domain.Customers{
			CustomerID: id,
			FirstName:  fake.FirstName(),
			LastName:   fake.LastName(),
			Phone:      fake.Phone(),
			Email:      fake.EmailAddress(),
			Street:     fake.Street(),
			City:       fake.City(),
			State:      sql.NullString{String: fake.State()},

			ZipCode: 111,
		}

		countEffect, err := customersRepository.Update(customers1)

		if err != nil {
			t.Errorf("error on testing customers : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		listCustomers[customers] = &customers1

		customersOnDatabase, err := customersRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing customers : %s", err.Error())
		}

		if customers1.Compare(*customersOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}
}

func TestAllCustomers(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing customers on database : %s", err.Error())
	}

	customersRepository := NewGenerateCustomersRepository(database)

	customersOnDatabase, err := customersRepository.All()

	for customers := range listCustomers {
		customer := *&listCustomers[customers]

		var exists bool

		exists = false

		for customers := range customersOnDatabase {
			exists = customer.Compare(*customersOnDatabase[customers]) || exists
		}

		if exists == false {
			t.Errorf("Found filed not exists %d -> %s", customer.CustomerID, customer.FirstName)
		}
	}
}

func TestSingleCustomers(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing customers on database : %s", err.Error())
	}

	customersRepository := NewGenerateCustomersRepository(database)

	for customers := range listCustomers {
		id := *&listCustomers[customers].CustomerID
		customersOnDatabase, err := customersRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing customers : %s", err.Error())
		}

		if customersOnDatabase == nil {
			t.Errorf("error field not exists : %s", err.Error())
		}
	}
}

func TestDeleteCustomers(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing customers on database : %s", err.Error())
	}

	customersRepository := NewGenerateCustomersRepository(database)

	for customers := range listCustomers {
		id := *&listCustomers[customers].CustomerID
		countEffect, err := customersRepository.Delete(id)

		if err != nil {
			t.Errorf("error on testing customers : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		customersOnDatabase, err := customersRepository.Get(int(id))

		if err != sql.ErrNoRows {
			t.Errorf("error on testing customers : %s", err.Error())
		}

		if customersOnDatabase != nil {
			t.Errorf("error field still exists : %s", err.Error())
		}
	}
}
