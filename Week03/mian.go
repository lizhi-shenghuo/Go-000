package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, errCtx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return startServer1(errCtx)
	})
	g.Go(func() error {
		return startServer2(errCtx)
	})
	g.Go(func() error {
		return ReceiveSignal(errCtx)
	})
	err := g.Wait()
	fmt.Println(err)
}

func startServer1(ctx context.Context) error {
	m := http.NewServeMux()
	m.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hello, i am Server1")
	})
	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	go func() {
		<-ctx.Done()
		fmt.Println("Server1 shut down!")
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func startServer2(ctx context.Context) error {
	m := http.NewServeMux()
	m.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hello, i am Server2")
	})
	s := &http.Server{
		Addr:    ":8081",
		Handler: m,
	}
	go func() {
		<-ctx.Done()
		fmt.Println("Server2 shut down!")
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func ReceiveSignal(ctx context.Context) error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigs:
		fmt.Println("the whole system will shut down")
		return errors.New("receive exit signal")
	case <-ctx.Done():
		return ctx.Err()
	}
}
