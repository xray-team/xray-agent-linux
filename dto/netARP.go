package dto

// NetArp describes the content of the file /proc/net/arp
type NetArp struct {
	Entries           map[string]uint
	IncompleteEntries map[string]uint
}

type ARPEntry struct {
	IP        string
	HWType    string
	Flags     string
	HWAddress string
	Mask      string
	Device    string
}
