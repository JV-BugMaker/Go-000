package main

import (
	"fmt"
	"net/http"
	"os"
	"errGroup"
)

func main(){
	g,ctx := errGroup.WithContext(context.Background())
	mux := http.NewServerMutex()
	mux.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Println("hello world")
	})

	server := &http.Server{
		Addr : ":8080",
		Handler:mux
	}

	g.Go(server.ListenAndServer())

	signs := make(chan os.Signal,1)
	signal.Notify(signs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select{
		case <- signs:
			fmt.Println("http shutdown")
			server.Shutdown(ctx)
	}


}
