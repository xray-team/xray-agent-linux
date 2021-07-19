package dto

type PSStat struct {
	PS []ProcessStat
}

type ProcessStat struct {
	// https://man7.org/linux/man-pages/man5/proc.5.html
	// PID - The process ID
	PID int64 // (1)
	// Name - The filename of the executable, in parentheses.
	// Strings longer than TASK_COMM_LEN (16) characters (including the terminating null byte) are silently truncated.
	Name string // (2)
	// State - The current state of process
	// One of the following characters, indicating process state:
	//  R  Running
	//  S  Sleeping in an interruptible wait
	//  D  Waiting in uninterruptible disk sleep
	//  Z  Zombie
	//  T  Stopped (on a signal) or (before Linux 2.6.33) trace stopped
	//  t  Tracing stop (Linux 2.6.33 onward)
	//  W  Paging (only before Linux 2.6.0)
	//  X  Dead (from Linux 2.6.0 onward)
	//  x  Dead (Linux 2.6.33 to 3.13 only)
	//  K  Wakekill (Linux 2.6.33 to 3.13 only)
	//  W  Waking (Linux 2.6.33 to 3.13 only)
	//  P  Parked (Linux 3.9 to 3.13 only)
	State string // (3)
	// UTime - Amount of time that this process has been scheduled in user mode.
	UTime int64 // (14)
	// STime Amount of time that this process has been scheduled in kernel mode.
	STime int64 // (15)
	// CuTime - Amount of time that this process's waited-for children have been scheduled in user mode.
	CuTime int64 // (16)
	// CsTime - Amount of time that this process's waited-for children have been scheduled in kernel mode
	CsTime int64 // (16)
	// GuestTime of the process (time spent running a virtual CPU for a guest operating system)
	GuestTime int64 // (43)
	// CGuestTime - Guest time of the process's children.
	CGuestTime int64 // (44)
	// Threads - Number of threads in this process
	Threads int64 // (20)
	// VSize - Virtual memory size in bytes.
	VSize int64 // (23)
	// Rss - Resident Set Size: number of pages the process has in real memory.
	// This is just the pages which count toward text, data, or stack space.
	// This does not include pages which have not been demand-loaded in, or which are swapped out.
	Rss int64 // (24)
}
