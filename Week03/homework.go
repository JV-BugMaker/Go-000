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


	g.Go(HttpServer(ctx,":8080",mux))
	g.Go(HttpServer(ctx,":8081",mux))
	g.Go(Signal(ctx))

	if err := g.Wait();err != nil{
		
		fmt.Println("something went wrong~")
	}
}

func Signal(ctx context.Context) error{
	sg := make(chan os.Signal,1)
	signal.Notify(sg, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer signal.Stop(sg)
	select{
		case <- sg:
			return errors.New("receive signal to stop")
		case <- ctx.Done():
			signal
			return nil
	}
}


func HttpServer(ctx context.Context,addr string,handler http.Handler )error{
	server := &http.Server{
		Addr : addr,
		Handler: handler,
	}
	defer server.Shutdown(ctx)
	return server.ListenAndServer()

}
