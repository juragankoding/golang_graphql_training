package repository

// import (
// 	"database/sql"
// 	"testing"

// 	"github.com/icrowley/fake"
// 	"github.com/juragankoding/golang_graphql_training/db"
// 	"github.com/juragankoding/golang_graphql_training/domain"
// )

// var listOrderItem []*domain.OrderItem

// func TestInsertOrderItem(t *testing.T) {
// 	database, err := db.GetDatabase()

// 	if err != nil {
// 		t.Errorf("error on testing orderItem on database : %s", err.Error())
// 	}

// 	orderItemRepository := NewGenerateOrderItemRepository(database)

// 	for loopingOrderItem := 0; loopingOrderItem < 10; loopingOrderItem++ {
// 		orderItem := domain.OrderItem{
// 			orderitem
// 		}

// 		lastID, err := orderItemRepository.Insert(orderItem)

// 		if err != nil {
// 			t.Errorf("error on testing orderItem : %s", err.Error())
// 		}

// 		orderItem.CustomerID = int(lastID)
// 		listOrderItem = append(listOrderItem, &orderItem)

// 		orderItemOnDatabase, err := orderItemRepository.Get(int(lastID))

// 		if err != nil {
// 			t.Errorf("error on testing orderItem : %s", err.Error())
// 		}

// 		if orderItem.Compare(*orderItemOnDatabase) == false {
// 			t.Errorf("field on database not same with on field input")
// 		}
// 	}

// }

// func TestUpdateOrderItem(t *testing.T) {
// 	database, err := db.GetDatabase()

// 	if err != nil {
// 		t.Errorf("error on testing orderItem on database : %s", err.Error())
// 	}

// 	orderItemRepository := NewGenerateOrderItemRepository(database)

// 	for orderItem := range listOrderItem {
// 		id := *&listOrderItem[orderItem].CustomerID

// 		orderItem := domain.OrderItem{
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

// 		countEffect, err := orderItemRepository.Update(orderItem)

// 		if err != nil {
// 			t.Errorf("error on testing orderItem : %s", err.Error())
// 		}

// 		if countEffect <= 0 {
// 			t.Errorf("nothing row update")
// 		}

// 		orderItemOnDatabase, err := orderItemRepository.Get(int(id))

// 		if err != nil {
// 			t.Errorf("error on testing orderItem : %s", err.Error())
// 		}

// 		if orderItem.Compare(*orderItemOnDatabase) == false {
// 			t.Errorf("field on database not same with on field input")
// 		}
// 	}
// }

// func TestAllOrderItem(t *testing.T) {
// 	database, err := db.GetDatabase()

// 	if err != nil {
// 		t.Errorf("error on testing orderItem on database : %s", err.Error())
// 	}

// 	orderItemRepository := NewGenerateOrderItemRepository(database)

// 	orderItemOnDatabase, err := orderItemRepository.All()

// 	for orderItem := range listOrderItem {
// 		customer := *&listOrderItem[orderItem]
// 		for orderItem := range orderItemOnDatabase {
// 			if customer.Compare(*orderItemOnDatabase[orderItem]) {
// 				t.Errorf("Found filed not exists %d -> %s", customer.CustomerID, customer.FirstName)
// 			}
// 		}
// 	}
// }

// func TestSingleOrderItem(t *testing.T) {
// 	database, err := db.GetDatabase()

// 	if err != nil {
// 		t.Errorf("error on testing orderItem on database : %s", err.Error())
// 	}

// 	orderItemRepository := NewGenerateOrderItemRepository(database)

// 	for orderItem := range listOrderItem {
// 		id := *&listOrderItem[orderItem].CustomerID
// 		orderItemOnDatabase, err := orderItemRepository.Get(int(id))

// 		if err != nil {
// 			t.Errorf("error on testing orderItem : %s", err.Error())
// 		}

// 		if orderItemOnDatabase == nil {
// 			t.Errorf("error field not exists : %s", err.Error())
// 		}
// 	}
// }

// func TestDeleteOrderItem(t *testing.T) {
// 	database, err := db.GetDatabase()

// 	if err != nil {
// 		t.Errorf("error on testing orderItem on database : %s", err.Error())
// 	}

// 	orderItemRepository := NewGenerateOrderItemRepository(database)

// 	for orderItem := range listOrderItem {
// 		id := *&listOrderItem[orderItem].CustomerID
// 		countEffect, err := orderItemRepository.Delete(id)

// 		if err != nil {
// 			t.Errorf("error on testing orderItem : %s", err.Error())
// 		}

// 		if countEffect <= 0 {
// 			t.Errorf("nothing row update")
// 		}

// 		orderItemOnDatabase, err := orderItemRepository.Get(int(id))

// 		if err != sql.ErrNoRows {
// 			t.Errorf("error on testing orderItem : %s", err.Error())
// 		}

// 		if orderItemOnDatabase != nil {
// 			t.Errorf("error field still exists : %s", err.Error())
// 		}
// 	}
// }
