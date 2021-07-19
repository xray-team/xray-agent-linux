package dto

type DiskSpaceUsage struct {
	Bytes  DiskSpaceBlockInfo
	Inodes DiskSpaceInodeInfo
}

// DiskSpaceBlockInfo represents a disk block usage statistics. All values are in bytes.
type DiskSpaceBlockInfo struct {
	Available uint64 // stat.Bavail, f_bavail; Free blocks available to unprivileged user
	Free      uint64 // stat.Bfree, f_bfree; Free blocks in filesystem
	Used      uint64 // calculated value: Total - Available
	Total     uint64 // stat.Blocks, f_blocks; Total data blocks in filesystem
}

// DiskSpaceInodeInfo represents a disk inode usage statistics.
type DiskSpaceInodeInfo struct {
	Free  uint64 // stat.Ffree, f_ffree; Free file nodes in filesystem
	Used  uint64 // calculated value: Total - Free
	Total uint64 // stat.Files, f_files; Total file nodes in filesystem
}
