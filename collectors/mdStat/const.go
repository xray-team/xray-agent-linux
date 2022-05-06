package mdStat

const (
	MDStatPath  = "/sys/block"
	MDSubFolder = "md"
	// in md directory
	ArraySizeFile     = "array_size"
	ArrayStateFile    = "array_state"
	ComponentSizeFile = "component_size"
	DegradedFile      = "degraded"
	LevelFile         = "level"
	MismatchCntFile   = "mismatch_cnt"
	RaidDisksFile     = "raid_disks"
	SyncActionFile    = "sync_action"
	SyncCompletedFile = "sync_completed"
	SyncSpeedFile     = "sync_speed"

	// in dev-xxx directory
	ErrorsFile = "errors"
	StateFile  = "state"
	SlotFile   = "slot"
)

var (
	MDStatsArrayStates = map[string]int{
		"clear":         1, // no devices, no size, no level
		"inactive":      2, // may have some settings, but array is not active all IO results in error
		"suspended":     3, // (not supported yet) - All IO requests will block. The array can be reconfigured.
		"readonly":      4, // no resync can happen. no superblocks get written.
		"read-auto":     5, // like readonly, but behaves like clean on a write request.
		"clean":         6, // no pending writes, but otherwise active.
		"active":        7, // fully active: IO and resync can be happening. When written to inactive array, starts with resync
		"write-pending": 8, // clean, but writes are blocked waiting for active to be written.
		"active-idle":   9, // like active, but no writes have been seen for a while (safe_mode_delay).
	}
	MDStatsDevStates = map[string]int{
		"faulty":           1, // device has been kicked from active use due to a detected fault, or it has unacknowledged bad blocks
		"in_sync":          2, // device is a fully in-sync member of the array
		"writemostly":      3, // device will only be subject to read requests if there are no other options. This applies only to raid1 arrays.
		"blocked":          4, // device has failed, and the failure hasn’t been acknowledged yet by the metadata handler.
		"spare":            5, // device is working, but not a full member.
		"write_error":      6, // device has ever seen a write error.
		"want_replacement": 7, // device is (mostly) working but probably should be replaced, either due to errors or due to user request.
		"replacement":      8, // device is a replacement for another active device with same raid_disk.
	}
	MDStatsSyncActions = map[string]int{
		"resync":  1, // redundancy is being recalculated after unclean shutdown or creation
		"recover": 2, // a hot spare is being built to replace a failed/missing device
		"idle":    3, // nothing is happening
		"check":   4, // a full check of redundancy was requested and is happening. This reads all blocks and checks them. A repair may also happen for some raid levels.
		"repair":  5, // a full check and repair is happening. This is similar to resync, but was requested by the user, and the write-intent bitmap is NOT used to optimise the process.
	}
)
