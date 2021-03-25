func receive(conn net.Conn){
 // receive message
   for { 
	reader := wsutil.NewReader(conn, ws.StateServerSide)
	_, err := reader.NextFrame()
	if err != nil {
		// handle error
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		// handle error
	}

	println(string(data))
}   