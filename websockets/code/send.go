func send(conn net.Conn, message string) {
	msg := "new server message"
	err := wsutil.WriteServerText(conn, message)
	if err != nil {
		// handle error
	}
}
