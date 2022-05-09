package nginx

type Config struct {
	Enabled  bool   `json:"enabled"`
	Endpoint string `json:"endpoint" validate:"required"`
	Timeout  int    `json:"timeout" validate:"required,min=1,max=120"`
}
