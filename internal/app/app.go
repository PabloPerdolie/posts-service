package app

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
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
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(a.serviceProvider.postService, a.serviceProvider.commentService),
	}))

	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	a.gqlServer = srv

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", a.gqlServer)

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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	return nil
}
