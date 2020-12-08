package main

import "fmt"

func main() {
	done := make(chan error,2)
	stop := make(chan struct{}))
	
	//使用goroutine来开启服务
	go func(){
		done <- HttpServer(stop)
	}()

	go func(){
		done <- HttpServer(stop)
	}()

	var stopped bool

	for i := 0;i < cap(done);i++{
		if err := <-done;err != nil {
			fmt.Println("error:%v",err)
		}

		if !stopped{
			stopped = true
			close(stop)
		}
	}

}

func HttpServer(addr string,handler http.Handler,stop <-chan struct{}) error{
	s := http.Server{
		Addr:addr,
		Handler:handler,
	}
	
	// 使用chan 来接收信号进行shutdown 优雅退出
	go func(){
		// 从chan中读取空值 其实就是在接收信号
		<-stop
		s.Shutdown(context.Background())
	}()
	// 正常这边应该是阻塞 当有值返回到时候 就一定是一个error
	return s.ListenAndServer()
}


