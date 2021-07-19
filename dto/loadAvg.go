package dto

// LoadAvg partially describes the content of the file /proc/loadavg
type LoadAvg struct {
	Last                     float64
	Last5m                   float64
	Last15m                  float64
	KernelSchedulingEntities int64
}
