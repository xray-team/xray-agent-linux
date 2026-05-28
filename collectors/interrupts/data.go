package interrupts

type Interrupts struct {
	Total  InterruptsStat
	PerCPU map[int]InterruptsStat
}

type InterruptsStat struct {
	Total int64 // The total number of interrupts
}
