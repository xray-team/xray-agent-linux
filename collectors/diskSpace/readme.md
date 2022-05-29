# DiskSpace
## Description
The DiskSpace collector picks the following metrics for each disk partition:

| Name             | Type  | Data source   | Description                                       |
|------------------|-------|---------------|---------------------------------------------------|
| BytesTotal       | gauge | Statfs        | Partition size in bytes.                          |
| BytesAvailable   | gauge | Statfs        | Bytes available.                                  |
| BytesFree        | gauge | Statfs        | Free space in bytes.                              |
| BytesFreePercent | gauge | Statfs        | Free space in percents.                           |
| BytesUsed        | gauge | Statfs        | Used space in bytes. BytesTotal - BytesAvailable. |
| InodesTotal      | gauge | Statfs        | The total number of inodes.                       |
| InodesFree       | gauge | Statfs        | Free inodes.                                      |
| InodesUsed       | gauge | Statfs        | Used inodes.                                      |

## Configuration
```json
{
  "collectors": {
    "DiskSpace": {
      "enabled": true,
      "fsTypes": [
        "ext4",
        "ext3",
        "ext2",
        "xfs",
        "jfs",
        "btrfs"
      ]
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"fsTypes"** - the list of allowed file systems.  
  More about file systems: [https://man7.org/linux/man-pages/man5/filesystems.5.html](https://man7.org/linux/man-pages/man5/filesystems.5.html)
