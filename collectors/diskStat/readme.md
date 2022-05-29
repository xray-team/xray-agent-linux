# DiskStat
## Description
The DiskStat collector picks the following metrics for each disk:

| Name                               | Type    | Data source     | Description                                                                                                                                                                   | Limitations                             |
|------------------------------------|---------|-----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------|
| ReadsCompletedSuccessfully         | counter | /proc/diskstats | The total number of reads completed successfully.                                                                                                                             |                                         |
| ReadsMerged                        | counter | /proc/diskstats | The total number of reads merged. Reads and writes which are adjacent to each other may be merged for efficiency. These operations are counted (and queued) as only one I/O.  |                                         |
| SectorsRead                        | counter | /proc/diskstats | The total number of sectors read successfully.                                                                                                                                |                                         |
| TimeSpentReading                   | counter | /proc/diskstats | The total number of milliseconds spent by all reads.                                                                                                                          |                                         |
| WritesCompleted                    | counter | /proc/diskstats | The total number of writes completed successfully.                                                                                                                            |                                         |
| WritesMerged                       | counter | /proc/diskstats | The total number of writes merged. Reads and writes which are adjacent to each other may be merged for efficiency. These operations are counted (and queued) as only one I/O. |                                         |
| SectorsWritten                     | counter | /proc/diskstats | The total number of sectors written successfully.                                                                                                                             |                                         |
| TimeSpentWriting                   | counter | /proc/diskstats | The total number of milliseconds spent by all writes.                                                                                                                         |                                         |
| IOsCurrentlyInProgress             | gauge   | /proc/diskstats | The number of I/Os currently in progress.                                                                                                                                     |                                         |
| TimeSpentDoingIOs                  | counter | /proc/diskstats | The number of milliseconds spent doing I/Os. This field increases so long as field IOsCurrentlyInProgress is nonzero.                                                         |                                         |
| WeightedTimeSpentDoingIOs          | counter | /proc/diskstats | The weighted number of milliseconds spent doing I/Os                                                                                                                          | Kernel 4.18+                            |
| DiscardsCompletedSuccessfully      | counter | /proc/diskstats | The total number of discards completed successfully.                                                                                                                          | Kernel 4.18+                            |
| DiscardsMerged                     | counter | /proc/diskstats | The total number of discards merged..                                                                                                                                         | Kernel 4.18+                            |
| SectorsDiscarded                   | counter | /proc/diskstats | The total number of sectors discarded successfully.                                                                                                                           | Kernel 4.18+                            |
| TimeSpentDiscarding                | counter | /proc/diskstats | The total number of milliseconds spent by all discards.                                                                                                                       | Kernel 4.18+                            |
| FlushRequestsCompletedSuccessfully | counter | /proc/diskstats | The total number of flush requests completed successfully.                                                                                                                    | Kernel 5.5+, Not tracked for partitions |
| TimeSpentFlushing                  | counter | /proc/diskstats | The total number of milliseconds spent by all flush requests.                                                                                                                 | Kernel 5.5+, Not tracked for partitions |

Details:
* [https://www.kernel.org/doc/Documentation/ABI/testing/procfs-diskstats](https://www.kernel.org/doc/Documentation/ABI/testing/procfs-diskstats)
* [https://www.kernel.org/doc/Documentation/admin-guide/iostats.rst](https://www.kernel.org/doc/Documentation/admin-guide/iostats.rst).
## Configuration
```json
{
  "collectors": {
    "DiskStat": {
      "enabled": true,
      "diskTypes": [
        8,
        9
      ],
      "excludePartitions": false,
      "excludeByName": []
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"diskTypes"** - list of allowed disk types (major numbers).  
  More about disk types: [https://www.kernel.org/doc/html/latest/admin-guide/devices.html](https://www.kernel.org/doc/html/latest/admin-guide/devices.html).  
  Common disk types:
    * 8 - SCSI(sata)-disks. For example: sda, sda1, sdb, sdb1, etc.
    * 9 - Metadisk (RAID). For example: md0, md1, etc.
* **"excludePartitions"**
    * **true** - exclude disk partitions from statistics
    * **false** - do not exclude disk partitions from statistics
* **"excludeByName"** - list of disk names to exclude from statistics. For example: "excludeByName": \["sda2", "sda3"\]
