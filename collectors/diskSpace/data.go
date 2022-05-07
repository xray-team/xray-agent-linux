package diskSpace

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

// Mounts describes content of the files /proc/mounts, /etc/fstab, /etc/mtab files
// http://man7.org/linux/man-pages/man5/fstab.5.html
type Mounts struct {
	Dev            string // Dev describes the block special device or remote filesystem to be mounted.
	MountPoint     string // MountPoint describes the mount point (target) for the	filesystem.
	FileSystemType string // FileSystemType describes the type of the filesystem.
	MountOptions   string // MountOptions describes the mount options associated with the filesystem. It is formatted as a comma-separated list of options.
	Dump           int64  // Dump field is used by dump to determine which filesystems	need to be dumped
	Pass           int64  // Pass is used by fsck to determine the order in which filesystem checks are done at boot time.
}
