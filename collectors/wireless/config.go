package wireless

type Config struct {
	Enabled            bool     `json:"enabled"`
	ExcludeByName      []string `json:"excludeByName"`
	ExcludeByOperState []string `json:"excludeByOperState" validate:"dive,oneof=unknown notpresent down lowerlayerdown testing dormant up"`
}
