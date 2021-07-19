package dto

// https://www.kernel.org/doc/html/v4.15/admin-guide/md.html
// https://github.com/torvalds/linux/blob/master/Documentation/admin-guide/md.rst
type MDStats struct {
	Stats map[string]MDStat
}

type MDStat struct {
	// Level represents the content of the file sys/block/mdN/md/level
	// This text file indicating the raid level. e.g. raid0, raid1, raid5, linear, multipath, faulty.
	// If no raid level has been set yet (array is still being assembled), the value will reflect whatever has been written to it,
	// which may be a name like the above, or may be a number such as 0, 5, etc.
	Level string

	// NumDisks represents the content of the file /sys/block/mdN/md/raid_disks
	// This is a text file with a simple number indicating the number of devices in a fully functional array.
	// If this is not yet known, the file will be empty. If an array is being resized this will contain the new number of devices.
	NumDisks int64

	// ArrayState represents the content of the file /sys/block/mdN/md/array_state
	// This file contains a single word which describes the current state of the array.
	//   "clear" - No devices, no size, no level
	//	 "inactive" - May have some settings, but array is not active all IO results in error
	//   "suspended" (not supported yet) - All IO requests will block. The array can be reconfigured.
	//   "readonly" - no resync can happen. no superblocks get written.
	//   "read-auto" - like readonly, but behaves like clean on a write request.
	//   "clean" - no pending writes, but otherwise active.
	//   "active" - fully active: IO and resync can be happening. When written to inactive array, starts with resync
	//   "write-pending" - clean, but writes are blocked waiting for active to be written.
	//   "active-idle" - like active, but no writes have been seen for a while (safe_mode_delay).
	ArrayState string

	// ArraySizeKBytes represents the content of the file /sys/block/mdN/md/array_size
	// The word "default" means the effective size of the array to be whatever size is actually available based on level, chunk_size, and component_size.
	ArraySizeKBytes int64

	// ComponentSizeKBytes represents the content of the file /sys/block/mdN/md/component_size
	ComponentSizeKBytes    int64
	DevStats               map[string]DevStats     // in files /sys/block/mdN/md/dev-XXX
	StatRaidWithRedundancy *StatRaidWithRedundancy // /sys/block/mdN/md/... for (RAID1,4,5,6,10)
}

// StatRaidWithRedundancy only for raids with redundancy (RAID1,4,5,6,10)
type StatRaidWithRedundancy struct {
	// SyncAction represents the content of the file /sys/block/mdN/md/sync_action
	// This is a text file that can be used to monitor and control the rebuild process. It contains one word which can be one of:
	//   "resync" - redundancy is being recalculated after unclean shutdown or creation
	//   "recover" - a hot spare is being built to replace a failed/missing device
	//   "idle" - nothing is happening
	//   "check" - A full check of redundancy was requested and is happening. This reads all blocks and checks them. A repair may also happen for some raid levels.
	//   "repair" - A full check and repair is happening. This is similar to resync, but was requested by the user, and the write-intent bitmap is NOT used to optimise the process.
	SyncAction string

	// NumDegraded represents the content of the file /sys/block/mdN/md/degraded
	// This contains a count of the number of devices by which the arrays is degraded.
	// So an optimal array will show 0. A single failed/missing drive will show 1, etc.
	NumDegraded int64

	// MismatchCnt represents the content of the file /sys/block/mdN/md/mismatch_cnt
	// When performing check and repair, and possibly when performing resync, md will count the number of errors that are found.
	// The count in mismatch_cnt is the number of sectors that were re-written, or (for check) would have been re-written.
	// As most raid levels work in units of pages rather than sectors, this may be larger than the number of actual errors by a factor of the number of sectors in a page.
	MismatchCnt int64

	// SyncCompletedSectors and NumSectors are represents the content of the file /sys/block/mdN/md/sync_completed
	// This shows the number of sectors that have been completed of whatever the current sync_action is, followed by the number of sectors in total that could need to be processed.
	// The two numbers are separated by a / thus effectively showing one value, a fraction of the process that is complete.
	// May contain the word "none"
	SyncCompletedSectors int64 // first value in /sys/block/mdN/md/sync_completed
	NumSectors           int64 // second value in /sys/block/mdN/md/sync_completed

	// SyncSpeed represents the content of the file /sys/block/mdN/md/sync_speed
	// This shows the current actual speed, in K/sec, of the current sync_action. It is averaged over the last 30 seconds.
	SyncSpeed int64 // /sys/block/mdN/md/sync_speed
}

type DevStats struct {
	// Slot represents the content of the file /sys/block/md*/md/dev-*/slot
	// This gives the role that the device has in the array. It will either be none if the device is not active in the array
	// (i.e. is a spare or has failed) or an integer less than the raid_disks number for the array indicating which position it currently fills.
	// This can only be set while assembling an array. A device for which this is set is assumed to be working.
	Slot string

	// State represents the content of the file /sys/block/md*/md/dev-*/state
	// This file recording the current state of the device in the array which can be a comma separated list of:
	//   "faulty" - device has been kicked from active use due to a detected fault, or it has unacknowledged bad blocks
	//   "in_sync" - device is a fully in-sync member of the array
	//   "writemostly" - device will only be subject to read requests if there are no other options. This applies only to raid1 arrays.
	//   "blocked" - device has failed, and the failure hasn’t been acknowledged yet by the metadata handler.
	//   "spare" - device is working, but not a full member.
	//   "write_error" - device has ever seen a write error.
	//   "want_replacement" - device is (mostly) working but probably should be replaced, either due to errors or due to user request.
	//   "replacement" - device is a replacement for another active device with same raid_disk.
	State string

	// Errors represents the content of the file /sys/block/md*/md/dev-*/errors
	// An approximate count of read errors that have been detected on this device but have not caused the device to be evicted from the array
	// (either because they were corrected or because they happened while the array was read-only).
	// When using version-1 metadata, this value persists across restarts of the array.
	Errors int64
}
