package netSNMP6

// NetSNMP6 describes the content of the file /proc/net/snmp6 (/proc/$PID/net/snmp6)
type NetSNMP6 struct {
	Counters map[string]int64
}
