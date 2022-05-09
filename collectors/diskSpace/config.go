package diskSpace

type Config struct {
	Enabled bool `json:"enabled"`
	// MonitoredFileSystemTypes is used by procMounts
	MonitoredFileSystemTypes []string `json:"fsTypes" validate:"dive,oneof=ext4 ext3 ext2 btrfs xfs jfs ufs zfs vfat squashfs fuseblk ntfs msdos hfs hfsplus"`
}
