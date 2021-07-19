package dto

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
