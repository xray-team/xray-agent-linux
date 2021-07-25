package conf

type GraphiteConf struct {
	Servers []GraphiteServerConf `json:"servers" validate:"required,dive"`
	DryRun  bool                 `json:"-"`
}

type GraphiteServerConf struct {
	Mode     string `json:"mode" validate:"oneof=tree tags"` // "tree"|"tags"
	Address  string `json:"address" validate:"required"`
	Protocol string `json:"protocol" validate:"oneof=tcp udp"` // "tcp"|"udp"
	Timeout  int    `json:"timeout" validate:"min=1,max=300"`  // in seconds
}
