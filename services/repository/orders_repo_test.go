package repository

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/icrowley/fake"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var listOrders []*domain.Orders

var idProductOrders int64
var idBrandsOrders int64
var idStaffOrders int64
var idStoreOrders int64
var idCategoryOrders int64
var idCustomersOrders int64
var idOrdersOrders int64

func TestInsertOrders(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing orders on database : %s", err.Error())
	}

	ordersRepository := NewGenerateOrdersRepository(database)
	productsRepository := NewGenerateProductsRepository(database)
	categoriesRepository := NewGenerateCategoriesRepository(database)
	brandsRepository := NewGenerateBrandsRepository(database)
	storesRepository := NewGenerateStoresRepository(database)
	staffsRepository := NewGenerateStaffsRepository(database)
	customersRepository := NewGenerateCustomersRepository(database)

	if idBrandsOrders, err = brandsRepository.Insert(domain.Brands{
		Name: fake.Brand(),
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idCategoryOrders, err = categoriesRepository.Insert(&domain.Categories{
		Name: fake.Brand(),
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idProductOrders, err = productsRepository.Insert(domain.Products{
		BrandID:     int(idBrandsOrders),
		CategoryID:  int(idCategoryOrders),
		ProductName: fake.ProductName(),
		ModelYear:   strconv.Itoa(fake.Year(1993, 2020)),
		ListPrice:   1000,
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idStoreOrders, err = storesRepository.Insert(domain.Stores{
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

	if idStaffOrders, err = staffsRepository.Insert(domain.Staffs{
		StoreID:   int(idStoreOrders),
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
		Email:     fake.LastName(),
		Phone:     fake.Phone(),
		Active:    "1",
	}); err != nil {
		t.Errorf(err.Error())
	}

	if idCustomersOrders, err = customersRepository.Insert(domain.Customers{
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

	for loopingOrders := 0; loopingOrders < 1; loopingOrders++ {
		orders := domain.Orders{
			CustomerID:   int(idCustomersOrders),
			OrderStatus:  1,
			OrderDate:    "20-11-2003",
			RequiredDate: "20-11-2003",
			ShippedDate:  "20-11-2003",
			StoreID:      int(idStaffOrders),
			StaffID:      int(idStaffOrders),
		}

		lastID, err := ordersRepository.Insert(orders)

		if err != nil {
			t.Errorf(err.Error())
		}

		if err != nil {
			t.Errorf("error on testing orders : %s", err.Error())
		}

		orders.OrderID = int(lastID)
		listOrders = append(listOrders, &orders)

		ordersOnDatabase, err := ordersRepository.Get(int(lastID))

		if err != nil {
			t.Errorf("error on testing orders : %s", err.Error())
		}

		if orders.Compare(*ordersOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}

}

func TestUpdateOrders(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing orders on database : %s", err.Error())
	}

	ordersRepository := NewGenerateOrdersRepository(database)

	for orders := range listOrders {
		id := *&listOrders[orders].OrderID

		orders1 := domain.Orders{
			OrderID:      id,
			CustomerID:   int(idCustomersOrders),
			OrderStatus:  1,
			OrderDate:    "20-11-2003",
			RequiredDate: "20-11-2003",
			ShippedDate:  "20-11-2003",
			StoreID:      int(idStaffOrders),
			StaffID:      int(idStaffOrders),
		}

		countEffect, err := ordersRepository.Update(orders1)

		if err != nil {
			t.Errorf("error on testing orders : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		listOrders[orders] = &orders1

		ordersOnDatabase, err := ordersRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing orders : %s", err.Error())
		}

		if orders1.Compare(*ordersOnDatabase) == false {
			t.Errorf("field on database not same with on field input")
		}
	}
}

func TestAllOrders(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing orders on database : %s", err.Error())
	}

	ordersRepository := NewGenerateOrdersRepository(database)

	ordersOnDatabase, err := ordersRepository.Fetch()

	if len(ordersOnDatabase) <= 0 {
		panic("error on gettering data")
	}

	for orders := range listOrders {
		order := *&listOrders[orders]

		var exists bool

		exists = false

		for orders := range ordersOnDatabase {
			exists = order.Compare(*ordersOnDatabase[orders]) || exists
		}

		if exists == false {
			t.Errorf("Found filed not exists %d -> %d", order.OrderID, *&ordersOnDatabase[orders].OrderID)
		}
	}
}

func TestSingleOrders(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing orders on database : %s", err.Error())
	}

	ordersRepository := NewGenerateOrdersRepository(database)

	for orders := range listOrders {
		id := *&listOrders[orders].OrderID
		ordersOnDatabase, err := ordersRepository.Get(int(id))

		if err != nil {
			t.Errorf("error on testing orders : %s", err.Error())
		}

		if ordersOnDatabase == nil {
			t.Errorf("error field not exists : %s", err.Error())
		}
	}
}

func TestDeleteOrders(t *testing.T) {
	database, err := db.GetDatabase()

	if err != nil {
		t.Errorf("error on testing orders on database : %s", err.Error())
	}

	ordersRepository := NewGenerateOrdersRepository(database)

	for orders := range listOrders {
		id := *&listOrders[orders].OrderID
		countEffect, err := ordersRepository.Delete(id)

		if err != nil {
			t.Errorf("error on testing orders : %s", err.Error())
		}

		if countEffect <= 0 {
			t.Errorf("nothing row update")
		}

		ordersOnDatabase, err := ordersRepository.Get(int(id))

		if err != sql.ErrNoRows {
			t.Errorf("error on testing orders : %s", err.Error())
		}

		if ordersOnDatabase != nil {
			t.Errorf("error field still exists : %s", err.Error())
		}
	}

	productsRepository := NewGenerateProductsRepository(database)
	categoriesRepository := NewGenerateCategoriesRepository(database)
	brandsRepository := NewGenerateBrandsRepository(database)
	storesRepository := NewGenerateStoresRepository(database)
	staffsRepository := NewGenerateStaffsRepository(database)
	customersRepository := NewGenerateCustomersRepository(database)

	if _, err := productsRepository.Delete(int(idProductOrders)); err != nil {
		t.Error(err.Error())
	}

	if _, err := categoriesRepository.Delete(int(idCategoryOrders)); err != nil {
		t.Error(err.Error())
	}

	if _, err := brandsRepository.Delete(int(idBrandsOrders)); err != nil {
		t.Error(err.Error())
	}

	if _, err := storesRepository.Delete(int(idStoreOrders)); err != nil {
		t.Error(err.Error())
	}

	if _, err := staffsRepository.Delete(int(idStaffOrders)); err != nil {
		t.Error(err.Error())
	}

	if _, err := customersRepository.Delete(int(idCustomersOrders)); err != nil {
		t.Error(err.Error())
	}
}
