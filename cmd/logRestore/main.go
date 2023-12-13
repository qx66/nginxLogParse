package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	watch(ctx)
	
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	
	exitChan := make(chan int)
	
	go func() {
		for {
			s := <-signalChan
			switch s {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				fmt.Println("hungup")
				exitChan <- 0
			
			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				fmt.Println("Warikomi")
				exitChan <- 0
			
			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				fmt.Println("force stop")
				exitChan <- 0
			
			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				fmt.Println("stop and core dump")
				exitChan <- 0
			
			default:
				fmt.Println("Unknown signal.")
				exitChan <- 1
			}
		}
	}()
	
	code := <-exitChan
	
	os.Exit(code)
}

func watch(ctx context.Context) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	
	//defer watcher.Close()
	
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Printf("name: %s, op: %s.\n ", event.Name, event.Op)
			case err := <-watcher.Errors:
				fmt.Println("err: ", err)
			case <-ctx.Done():
				watcher.Close()
			}
		}
	}()
	
	err = watcher.Add("/tmp")
	if err != nil {
		return
	}
}
