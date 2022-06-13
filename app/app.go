package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"ports/ports"

	"github.com/gorilla/mux"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type PortsApp struct {
	config ConfigManager
	server http.Server
	db     db.Session
}

func Init(configManager ConfigManager) (*PortsApp, error) {
	log.SetOutput(os.Stdout)

	connURL, err := postgresql.ParseURL(dbURL)
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to db")

	database, err := postgresql.Open(connURL)
	if err != nil {
		fmt.Println("could not connect to database ", err)
		return nil, err
	}

	return &PortsApp{
		config: configManager,
		db:     database,
	}, nil
}

func (app *PortsApp) Start() error {
	app.server = http.Server{
		Addr:    app.config.ListenAddr,
		Handler: app.Router(),
	}

	listener, err := net.Listen("tcp", app.server.Addr)
	if err == nil {
		fmt.Println("service started ", app.server.Addr)
		err = app.server.Serve(listener)
	}
	if err == http.ErrServerClosed {
		fmt.Println("http server is closed")
		return nil
	}
	if err != nil {
		fmt.Println("http server failed", err)
	}
	return err
}

func (app *PortsApp) Shutdown() {
	err := app.db.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func (app *PortsApp) Router() *mux.Router {
	router := mux.NewRouter()

	portController := ports.NewController(ports.NewRepository(app.db))

	router.
		Path("/ports").
		Methods(http.MethodPost).
		HandlerFunc(portController.UpdatePorts)

	router.
		Path("/ports/import").
		Methods(http.MethodPost).
		HandlerFunc(portController.ImportPorts)

	return router
}
