package dto

type PS struct {
	Count            int64
	Limit            int64
	InStateRunning   int64
	InStateIdle      int64
	InStateSleeping  int64
	InStateDiskSleep int64
	InStateStopped   int64
	InStateZombie    int64
	InStateDead      int64
	Threads          int64
	ThreadsLimit     int64
}
