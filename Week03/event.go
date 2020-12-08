package main


func search(term string) (string,error){
	time.Sleep(time.Seconds * 20)
	return "some value",nil
}

type result struct {
	record string
	err error
}


func process(term string) error{
	ctx,cancel := context.WithTimeout(context.Background(),time.Seconds * 10)
	defer cancel()

	ch := make(chan result)

	// goroutine
	go func(){
		record,err := search(term)
		ch <- result{record:record,err:err}
	}()

	select{
		case <-ctx.Done():
			return errors.New("timeout")
		case result := <-ch:
			if result.err != nil{
				// do error wrap
			}
			return nil
	}

}
