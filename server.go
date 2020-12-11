package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/juragankoding/golang_graphql_training/categories/repository"
	"github.com/juragankoding/golang_graphql_training/categories/usecase"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/generated"
	"github.com/juragankoding/golang_graphql_training/graph"
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
	respositoryCategories := repository.NewCategoriesRepository(db)

	//prepare usecase
	categoriesUseCase := usecase.NewCategoriesUserCase(respositoryCategories)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DomainCategoriUseCase: categoriesUseCase,
	}}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	http.Handle("/", srv)

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
