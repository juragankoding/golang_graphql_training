// package repository

// import (
// 	"database/sql"
// 	"testing"

// 	"github.com/icrowley/fake"
// 	"github.com/juragankoding/golang_graphql_training/db"
// 	"github.com/juragankoding/golang_graphql_training/domain"
// )

// var listOrders []*domain.Orders

// func TestInsertOrders(t *testing.T) {
// 	database, err := db.GetDatabase()

// 	if err != nil {
// 		t.Errorf("error on testing orders on database : %s", err.Error())
// 	}

// 	ordersRepository := NewGenerateOrdersRepository(database)

// 	for loopingOrders := 0; loopingOrders < 10; loopingOrders++ {
// 		orders := domain.Orders{
// 			FirstName: fake.FirstName(),
// 			LastName:  fake.LastName(),
// 			Phone:     fake.Phone(),
// 			Email:     fake.EmailAddress(),
// 			Street:    fake.Street(),
// 			City:      fake.City(),
// 			State:     sql.NullString{String: fake.State()},
// 			ZipCode:   111,
// 		}

// 		lastID, err := ordersRepository.Insert(orders)

// 		if err != nil {
// 			t.Errorf("error on testing orders : %s", err.Error())
// 		}

// 		orders.CustomerID = int(lastID)
// 		listOrders = append(listOrders, &orders)

// 		ordersOnDatabase, err := ordersRepository.Get(int(lastID))

// 		if err != nil {
// 			t.Errorf("error on testing orders : %s", err.Error())
// 		}

// 		if orders.Compare(*ordersOnDatabase) == false {
// 			t.Errorf("field on database not same with on field input")
// 		}
// 	}

// }

// func TestUpdateOrders(t *testing.T) {
// 	database, err := db.GetDatabase()

// 	if err != nil {
// 		t.Errorf("error on testing orders on database : %s", err.Error())
// 	}

// 	ordersRepository := NewGenerateOrdersRepository(database)

// 	for orders := range listOrders {
// 		id := *&listOrders[orders].CustomerID

// 		orders := domain.Orders{
// 			CustomerID: id,
// 			FirstName:  fake.FirstName(),
// 			LastName:   fake.LastName(),
// 			Phone:      fake.Phone(),
// 			Email:      fake.EmailAddress(),
// 			Street:     fake.Street(),
// 			City:       fake.City(),
// 			State:      sql.NullString{String: fake.State()},

// 			ZipCode: 111,
// 		}

// 		countEffect, err := ordersRepository.Update(orders)

// 		if err != nil {
// 			t.Errorf("error on testing orders : %s", err.Error())
// 		}

// 		if countEffect <= 0 {
// 			t.Errorf("nothing row update")
// 		}

// 		ordersOnDatabase, err := ordersRepository.Get(int(id))

// 		if err != nil {
// 			t.Errorf("error on testing orders : %s", err.Error())
// 		}

// 		if orders.Compare(*ordersOnDatabase) == false {
// 			t.Errorf("field on database not same with on field input")
// 		}
// 	}
// }

// func TestAllOrders(t *testing.T) {
// 	database, err := db.GetDatabase()

// 	if err != nil {
// 		t.Errorf("error on testing orders on database : %s", err.Error())
// 	}

// 	ordersRepository := NewGenerateOrdersRepository(database)

// 	ordersOnDatabase, err := ordersRepository.All()

// 	for orders := range listOrders {
// 		customer := *&listOrders[orders]
// 		for orders := range ordersOnDatabase {
// 			if customer.Compare(*ordersOnDatabase[orders]) {
// 				t.Errorf("Found filed not exists %d -> %s", customer.CustomerID, customer.FirstName)
// 			}
// 		}
// 	}
// }

// func TestSingleOrders(t *testing.T) {
// 	database, err := db.GetDatabase()

// 	if err != nil {
// 		t.Errorf("error on testing orders on database : %s", err.Error())
// 	}

// 	ordersRepository := NewGenerateOrdersRepository(database)

// 	for orders := range listOrders {
// 		id := *&listOrders[orders].CustomerID
// 		ordersOnDatabase, err := ordersRepository.Get(int(id))

// 		if err != nil {
// 			t.Errorf("error on testing orders : %s", err.Error())
// 		}

// 		if ordersOnDatabase == nil {
// 			t.Errorf("error field not exists : %s", err.Error())
// 		}
// 	}
// }

// func TestDeleteOrders(t *testing.T) {
// 	database, err := db.GetDatabase()

// 	if err != nil {
// 		t.Errorf("error on testing orders on database : %s", err.Error())
// 	}

// 	ordersRepository := NewGenerateOrdersRepository(database)

// 	for orders := range listOrders {
// 		id := *&listOrders[orders].CustomerID
// 		countEffect, err := ordersRepository.Delete(id)

// 		if err != nil {
// 			t.Errorf("error on testing orders : %s", err.Error())
// 		}

// 		if countEffect <= 0 {
// 			t.Errorf("nothing row update")
// 		}

// 		ordersOnDatabase, err := ordersRepository.Get(int(id))

// 		if err != sql.ErrNoRows {
// 			t.Errorf("error on testing orders : %s", err.Error())
// 		}

// 		if ordersOnDatabase != nil {
// 			t.Errorf("error field still exists : %s", err.Error())
// 		}
// 	}
// }

package repository
