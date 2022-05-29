package nginx

type StubStatus struct {
	Active   uint64 // The current number of active client connections including Waiting connections.
	Accepts  uint64 // The total number of accepted client connections.
	Handled  uint64 // The total number of handled connections.
	Requests uint64 // The total number of client requests.
	Reading  uint64 // The current number of connections where nginx is reading the request header.
	Writing  uint64 // The current number of connections where nginx is writing the response back to the client.
	Waiting  uint64 // The current number of idle client connections waiting for a request.
}
