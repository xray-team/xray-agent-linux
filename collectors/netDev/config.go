package netDev

type Config struct {
	Enabled          bool `json:"enabled"`
	ExcludeLoopbacks bool `json:"excludeLoopbacks"`
	ExcludeWireless  bool `json:"excludeWireless"`
	ExcludeBridges   bool `json:"excludeBridges"`
	// Virtual interfaces are network interfaces that are not associated with an any physical interface.
	// Virtual interface examples: loopback, bridge, tun, vlan, ...
	ExcludeVirtual     bool     `json:"excludeVirtual"`
	ExcludeByName      []string `json:"excludeByName"`
	ExcludeByOperState []string `json:"excludeByOperState" validate:"dive,oneof=unknown notpresent down lowerlayerdown testing dormant up"`
}
