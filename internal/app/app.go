package app

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"graphql-comments/internal/config"
	"graphql-comments/internal/graphql/graph"
	"graphql-comments/internal/graphql/graph/generated"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

type App struct {
	serviceProvider *serviceProvider
	gqlServer       *handler.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Run() error {
	return a.runServer()
}

func (a *App) initServer(ctx context.Context) error {
	a.gqlServer = handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(a.serviceProvider.postService, a.serviceProvider.commentService),
	}))
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.InitConfig()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(config.CONFIG.UseInMemory)
	return nil
}

func (a *App) runServer() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", a.gqlServer)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	return nil
}
