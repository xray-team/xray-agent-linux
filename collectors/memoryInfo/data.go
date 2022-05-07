package memoryInfo

// MemoryInfo partially describes the content of the file /proc/meminfo
type MemoryInfo struct {
	MemTotal     int64
	MemFree      int64
	MemAvailable int64 // absent in old kernels
	Buffers      int64
	Cached       int64
	SwapTotal    int64
	SwapFree     int64
}
