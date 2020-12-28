package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/generated"
	"github.com/juragankoding/golang_graphql_training/graph"
	"github.com/juragankoding/golang_graphql_training/middleware"
	"github.com/juragankoding/golang_graphql_training/services/repository"
	"github.com/juragankoding/golang_graphql_training/services/usecase"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	//prepare conect database on repository
	db, err := db.GetDatabase()

	if err != nil {
		log.Fatal(err)

		return
	}

	//prepare repository
	repositoryCategories := repository.NewGenerateCategoriesRepository(db)
	repositoryCustomers := repository.NewGenerateCustomersRepository(db)
	repositoryBrands := repository.NewGenerateBrandsRepository(db)
	repositoryOrderItem := repository.NewGenerateOrderItemRepository(db)
	repositoryOrders := repository.NewGenerateOrdersRepository(db)
	repositoryProducts := repository.NewGenerateProductsRepository(db)
	repositoryStaffs := repository.NewGenerateStaffsRepository(db)
	repositoryStocks := repository.NewGenerateStocksRepository(db)
	repositoryUser := repository.NewGenerateUserRepository(db)
	repositoryStores := repository.NewGenerateStoresRepository(db)

	resolver := graph.Resolver{
		CategoriesUseCase: usecase.NewGenerateCategoriesUserCase(repositoryCategories),
		CustomersUseCase:  usecase.NewGenerateCustomerUseCase(repositoryCustomers),
		BrandsUseCase:     usecase.NewGenerateBrandsUseCase(repositoryBrands),
		OrderItemUseCase:  usecase.NewGenerateOderItemUseCase(repositoryOrderItem),
		OrdersUseCase:     usecase.NewGenerateOrdersUseCase(repositoryOrders),
		ProductUseCase:    usecase.NewGenerateProductUseCase(repositoryProducts),
		StaffsUseCase:     usecase.NewGenerateStaffsUseCase(repositoryStaffs),
		StocksUseCase:     usecase.NewGenerateStockUseCase(repositoryStocks),
		UserUseCase:       usecase.NewGenerateUserUseCase(&repositoryUser),
		StoreUseCase:      usecase.NewGenereateStoreUseCase(&repositoryStores),
	}

	router := chi.NewRouter()

	router.Use(middleware.JwtMiddleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))

	router.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	router.Handle("/", srv)

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
