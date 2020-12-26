package graph

import "github.com/juragankoding/golang_graphql_training/domain"

type Resolver struct {
	CategoriesUseCase domain.CategoriesUseCase
	CustomersUseCase  domain.CustomersUseCase
	BrandsUseCase     domain.BrandsUseCase
	OrderItemUseCase  domain.OrderItemUseCase
	OrdersUseCase     domain.OrdersUseCase
	ProductUseCase    domain.ProductsUseCase
	StaffsUseCase     domain.StaffsUseCase
	StocksUseCase     domain.StocksUseCase
	UserUseCase       domain.UserUseCase
	StoreUseCase      domain.StoresUseCase
}
