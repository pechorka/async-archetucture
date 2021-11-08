package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/jackc/pgx/v4"
	"github.com/pechorka/async-architecture/app/auth/handlers"
	"github.com/pkg/errors"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "")
	if err != nil {
		return errors.Wrap(err, "failed to connect to database")
	}
	defer func() {
		cerr := conn.Close(ctx)
		if cerr != nil {
			log.Println("[ERROR] while closing db connection")
		}
	}()
	adapter := pgx4adapter.NewConn(conn)
	tokenStore, err := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	if err != nil {
		return errors.Wrap(err, "init token store")
	}
	defer func() {
		cerr := tokenStore.Close()
		if cerr != nil {
			log.Println("[ERROR] while closing tokenStore")
		}
	}()

	clientStore, err := pg.NewClientStore(adapter)
	if err != nil {
		return errors.Wrap(err, "init client store")
	}
	manager := manage.NewDefaultManager()

	manager.MapClientStorage(clientStore)
	manager.MapTokenStorage(tokenStore)

	router := handlers.New(manager)

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	srvErr := make(chan error, 1)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, os.Interrupt)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			srvErr <- err
		}
	}()

	select {
	case err := <-srvErr:
		return errors.Wrap(err, "stopping server failed")
	case <-shutdown:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := srv.Shutdown(ctx)
		errors.Wrap(err, "couln's shutdown server gracefully")
	}

	return nil
}
