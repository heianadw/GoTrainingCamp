package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	svr := http.Server{Addr: "localhost:8080"}
	//http
	eg.Go(func() error {
		fmt.Println("http server start")
		go func() {
			<-ctx.Done()
			fmt.Println("http context done")
			svr.Shutdown(context.TODO())
		}()
		return svr.ListenAndServe()
	})

	//signal
	eg.Go(func() error {
		signals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
		sig := make(chan os.Signal, len(signals))
		signal.Notify(sig, signals...)
		for {
			fmt.Println("signal transmission")
			select {
			case <-ctx.Done():
				fmt.Println("signal context done")
				return ctx.Err()
			case <-sig:
				return nil
			}
		}
	})

	//interrupt
	eg.Go(func() error {
		fmt.Println("interrupt the server")
		time.Sleep(3 * time.Second)
		fmt.Println("interrupt completed")
		return errors.New("interrupt error")
	})

	err := eg.Wait()
	fmt.Println(err)
}
