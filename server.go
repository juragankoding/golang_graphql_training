package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/cache"
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

	var redisUtil *redis.Client

	redisUtil = cache.GetRedisClient()

	contextBackground := context.Background()

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
		CategoriesUseCase: usecase.NewGenerateCategoriesUserCase(repositoryCategories,
			contextBackground,
			redisUtil),
		CustomersUseCase: usecase.NewGenerateCustomerUseCase(repositoryCustomers,
			contextBackground,
			redisUtil),
		BrandsUseCase: usecase.NewGenerateBrandsUseCase(repositoryBrands,
			contextBackground,
			redisUtil),
		OrderItemUseCase: usecase.NewGenerateOderItemUseCase(repositoryOrderItem,
			contextBackground,
			redisUtil),
		OrdersUseCase: usecase.NewGenerateOrdersUseCase(repositoryOrders,
			contextBackground,
			redisUtil),
		ProductUseCase: usecase.NewGenerateProductUseCase(repositoryProducts,
			contextBackground,
			redisUtil),
		StaffsUseCase: usecase.NewGenerateStaffsUseCase(repositoryStaffs,
			contextBackground,
			redisUtil),
		StocksUseCase: usecase.NewGenerateStockUseCase(repositoryStocks,
			contextBackground,
			redisUtil),
		UserUseCase: usecase.NewGenerateUserUseCase(&repositoryUser),
		StoreUseCase: usecase.NewGenereateStoreUseCase(&repositoryStores,
			contextBackground,
			redisUtil),
	}

	router := chi.NewRouter()

	router.Use(middleware.JwtMiddleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))

	router.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	router.Handle("/", srv)

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
