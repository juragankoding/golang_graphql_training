package repository

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/icrowley/fake"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var listOrderItem []*domain.OrderItem

var idOrderOrderItem domain.OrderItem
var idCategoryItem domain.Categories

var idProductOrderItem int64
var idBrandsOrderItem int64
var idStaffOrderItem int64
var idStoreOrderItem int64
var idCategoryOrderItem int64
var idCustomersOrderItem int64
var idOrdersOrderItem int64

func TestInsertOrderItem(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing orderItem on database : %s", err.Error())
	}

	orderItemRepository := NewGenerateOrderItemRepository(database)
	ordersRepository := NewGenerateOrdersRepository(database)
	productsRepository := NewGenerateProductsRepository(database)
	categoriesRepository := NewGenerateCategoriesRepository(database)
	brandsRepository := NewGenerateBrandsRepository(database)
	storesRepository := NewGenerateStoresRepository(database)
	staffsRepository := NewGenerateStaffsRepository(database)
	customersRepository := NewGenerateCustomersRepository(database)

	if idBrandsOrderItem, err = brandsRepository.Insert(domain.Brands{
		Name: fake.Brand(),
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idCategoryOrderItem, err = categoriesRepository.Insert(&domain.Categories{
		Name: fake.Brand(),
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idProductOrderItem, err = productsRepository.Insert(domain.Products{
		BrandID:     int(idBrandsOrderItem),
		CategoryID:  int(idCategoryOrderItem),
		ProductName: fake.ProductName(),
		ModelYear:   strconv.Itoa(fake.Year(1993, 2020)),
		ListPrice:   1000,
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idStoreOrderItem, err = storesRepository.Insert(domain.Stores{
		StoreName: fake.DomainName(),
		Phone:     fake.Phone(),
		Email:     fake.EmailAddress(),
		Street:    fake.Street(),
		City:      fake.City(),
		State:     fake.State(),
		ZipCode:   fake.Zip(),
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idStaffOrderItem, err = staffsRepository.Insert(domain.Staffs{
		StoreID:   int(idStoreOrderItem),
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
		Email:     fake.LastName(),
		Phone:     fake.Phone(),
		Active:    "1",
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idCustomersOrderItem, err = customersRepository.Insert(domain.Customers{
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
		Phone:     fake.Phone(),
		Email:     fake.EmailAddress(),
		Street:    fake.Street(),
		City:      fake.City(),
		State:     sql.NullString{String: fake.State()},
		ZipCode:   111,
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idOrdersOrderItem, err = ordersRepository.Insert(domain.Orders{
		CustomerID:   int(idCustomersOrderItem),
		OrderStatus:  1,
		OrderDate:    "20-11-2003",
		RequiredDate: "20-11-2003",
		ShippedDate:  "20-11-2003",
		StoreID:      int(idStaffOrderItem),
		StaffID:      int(idStaffOrderItem),
	}); err != nil {
		t.Errorf(err.Error())
	}

	for loopingOrderItem := 0; loopingOrderItem < 1; loopingOrderItem++ {
		orderItem := domain.OrderItem{
			OrderID:   int(idOrdersOrderItem),
			ProductID: int(idProductOrderItem),
			Quantity:  10,
			ListPrice: 10000,
			Discount:  10,
		}

		lastID, err := orderItemRepository.Insert(orderItem)

		if err != nil {
			t.Errorf("error on testing orderItem : %s", err.Error())
		}

		orderItem.ItemID = int(lastID)
		listOrderItem = append(listOrderItem, &orderItem)

		orderItemOnDatabase, err := orderItemRepository.Get(int(lastID))

		if err != nil {
			t.Errorf("error on testing orderItem : %s", err.Error())
		}

		if orderItem.Compare(*orderItemOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}

}

func TestUpdateOrderItem(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing orderItem on database : %s", err.Error())
	}

	orderItemRepository := NewGenerateOrderItemRepository(database)

	for orderItem := range listOrderItem {
		id := *&listOrderItem[orderItem].ItemID

		orderItem1 := domain.OrderItem{
			ItemID:    id,
			OrderID:   int(idOrdersOrderItem),
			ProductID: int(idProductOrderItem),
			Quantity:  100,
			ListPrice: 1000,
			Discount:  1,
		}

		countEffect, err := orderItemRepository.Update(orderItem1)

		if err != nil {
			t.Errorf("error on testing orderItem : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		listOrderItem[orderItem] = &orderItem1
		orderItemOnDatabase, err := orderItemRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing orderItem : %s", err.Error())
		}

		if orderItem1.Compare(*orderItemOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}
}

func TestAllOrderItem(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing orderItem on database : %s", err.Error())
	}

	orderItemRepository := NewGenerateOrderItemRepository(database)

	orderItemOnDatabase, err := orderItemRepository.Fetch()

	for orderItem := range listOrderItem {
		customer := *&listOrderItem[orderItem]

		var exists bool

		exists = false

		for orderItem := range orderItemOnDatabase {
			exists = customer.Compare(*orderItemOnDatabase[orderItem]) || exists
		}

		if exists == false {
			t.Errorf("Found filed not exists %d -> %d", customer.OrderID, customer.Quantity)

		}
	}
}

func TestSingleOrderItem(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing orderItem on database : %s", err.Error())
	}

	orderItemRepository := NewGenerateOrderItemRepository(database)

	for orderItem := range listOrderItem {
		id := *&listOrderItem[orderItem].ItemID
		orderItemOnDatabase, err := orderItemRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing orderItem : %s", err.Error())
		}

		if orderItemOnDatabase == nil {
			t.Errorf("error field not exists : %s", err.Error())
		}
	}
}

func TestDeleteOrderItem(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing orderItem on database : %s", err.Error())
	}

	orderItemRepository := NewGenerateOrderItemRepository(database)

	for orderItem := range listOrderItem {
		id := *&listOrderItem[orderItem].ItemID
		countEffect, err := orderItemRepository.Delete(id)

		if err != nil {
			t.Errorf("error on testing orderItem : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		orderItemOnDatabase, err := orderItemRepository.Get(int(id))

		if err != sql.ErrNoRows {
			t.Errorf("error on testing orderItem : %s", err.Error())
		}

		if orderItemOnDatabase != nil {
			t.Errorf("error field still exists : %s", err.Error())
		}
	}

	ordersRepository := NewGenerateOrdersRepository(database)
	productsRepository := NewGenerateProductsRepository(database)
	categoriesRepository := NewGenerateCategoriesRepository(database)
	brandsRepository := NewGenerateBrandsRepository(database)
	storesRepository := NewGenerateStoresRepository(database)
	staffsRepository := NewGenerateStaffsRepository(database)
	customersRepository := NewGenerateCustomersRepository(database)

	if _, err := ordersRepository.Delete(int(idOrdersOrderItem)); err != nil {
		t.Error(err.Error())
	}

	if _, err := productsRepository.Delete(int(idProductOrderItem)); err != nil {
		t.Error(err.Error())
	}

	if _, err := categoriesRepository.Delete(int(idCategoryOrderItem)); err != nil {
		t.Error(err.Error())
	}

	if _, err := brandsRepository.Delete(int(idBrandsOrderItem)); err != nil {
		t.Error(err.Error())
	}

	if _, err := storesRepository.Delete(int(idStoreOrderItem)); err != nil {
		t.Error(err.Error())
	}

	if _, err := staffsRepository.Delete(int(idStaffOrderItem)); err != nil {
		t.Error(err.Error())
	}

	if _, err := customersRepository.Delete(int(idCustomersOrderItem)); err != nil {
		t.Error(err.Error())
	}

}
