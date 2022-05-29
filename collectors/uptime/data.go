package uptime

// Uptime describes the content of the file /proc/uptime
type Uptime struct {
	Uptime float64 // The total number of seconds the system has been up.
	Idle   float64 // The sum of how much time each core has spent idle, in seconds
}
