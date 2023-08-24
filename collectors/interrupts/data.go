package interrupts

type Interrupts struct {
	Total  int64         // Total interrupts
	PerCPU map[int]int64 // Number of interrupts for each cpu since system startup.
}
