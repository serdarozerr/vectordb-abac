package main

import (
	"context"
	"fmt"
	"github.com/serdarozerr/vectordb-abac/config"
	"github.com/serdarozerr/vectordb-abac/internal/instance"
	"github.com/serdarozerr/vectordb-abac/internal/server"
	"github.com/serdarozerr/vectordb-abac/internal/service"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

//func CreateCollection(w http.ResponseWriter, r *http.Request) {
//
//	fmt.Println("CreateCollection")
//}
//
//func LogMiddleware(next http.Handler) http.Handler {
//
//	//
//	var hf http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
//		println("Logged the request")
//		next.ServeHTTP(w, r)
//	}
//	return hf
//}

func run(ctx context.Context, conf *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Configure logger
	l := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// based on the configuration create repository instance
	r, err := instance.NewRepository(conf.VectorDB.Type, conf)
	if err != nil {
		return err
	}
	//Create services
	ds := service.NewDBService(r)

	// Create llm model
	llm := service.NewLLM(conf.LLM.ApiKey)

	// Init Cache Repo.
	c, err := instance.NewMemeDbInstance(conf, "redis")
	if err != nil {
		return err
	}

	// Get the  server handler
	srv := server.NewServer(l, conf, ds, llm, c)

	httpServer := &http.Server{
		Handler: srv,
		Addr:    net.JoinHostPort("localhost", strconv.Itoa(conf.App.Port)),
	}

	go func() {
		log.Printf("Server start to listening on Port %d", httpServer.Addr)

		// gonna be blocked on ListenAndServe method until it canceled
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Error listening on Port %d: %v", httpServer.Addr, err)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done() // gonna block until we canceled it. we  increase the wg.Add by one because after getting Done signal
		// we can do some clean up below this line. Like disconnect DB etc.

		fmt.Println("Context canceled. Performing cleanup...")
		shutdownCtx := context.Background()
		shutdownCtx, shutdownCancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer shutdownCancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "Error shutting down server on Port %d: %v", httpServer.Addr, err)
		}

	}()

	wg.Wait()

	return nil
}

func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/CreateCollection", CreateCollection)
	//loginMW := LogMiddleware(mux)
	//
	//println("Server Started")
	//http.ListenAndServe("localhost:8000", loginMW)

	// Get the config values
	conf, err := config.ReadYamlFile("./config.yaml")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	ctx := context.Background()
	if err := run(ctx, conf); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
