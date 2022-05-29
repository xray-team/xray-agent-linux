# MDStat
## Description
The MDStat collector picks information about the state of the linux software RAID. The following metrics are collected for each md-device:

| Name                 | Type    | Data source                          | Description                                                                                                                                                                       |
|----------------------|---------|--------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Level                | gauge   | /sys/block/\[md*\]/md/level          | RAID type: raid1 - 1, raid6 - 6, raid10 - 10, и т.д.                                                                                                                              |
| NumDisks             | gauge   | /sys/block/\[md*\]/md/raid_disks     | The number of disks in a fully functional array.                                                                                                                                  |
| ArrayState           | gauge   | /sys/block/\[md*\]/md/array_state    | The current state of the array: "clear" - 1, "inactive" - 2, "suspended" - 3, "readonly" - 4, "read-auto" - 5, "clean" - 6, "active" - 7, "write-pending" - 8, "active-idle" - 9. |
| ArraySize            | gauge   | /sys/block/\[md*\]/md/               | The array size in kB.                                                                                                                                                             |
| SyncAction           | gauge   | /sys/block/\[md*\]/md/sync_action    | The sync state of the array: "resync" - 1, "recover" - 2, "idle" - 3, "check" - 4, "repair" - 5.                                                                                  |
| NumDegraded          | gauge   | /sys/block/\[md*\]/md/degraded       | The number of devices by which the arrays is degraded.                                                                                                                            |
| MismatchCnt          | counter | /sys/block/\[md*\]/md/mismatch_cnt   | The number of sectors that were re-written, or (for check) would have been re-written                                                                                             |
| SyncCompletedSectors | counter | /sys/block/\[md*\]/md/sync_completed | The number of sectors that have been completed of whatever the current sync_action is.                                                                                            |
| NumSectors           | gauge   | /sys/block/\[md*\]/md/               | The total number of sectors to process.                                                                                                                                           |
| SyncSpeed            | gauge   | /sys/block/\[md*\]/md/sync_speed     | The current actual speed, in K/sec, of the current sync_action.                                                                                                                   |

And for each disk:

| Name    | Type   | Data source                         | Description                                                                                                                                                                                                                          |
|---------|--------|-------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Slot    | gauge  | /sys/block/\[mdN\]/md/\[dev\]slot   | The role that the device has in the array.                                                                                                                                                                                           |
| State   | gauge  | /sys/block/\[mdN\]/md/\[dev\]state  | The disk state: "faulty" - 1, "in_sync" - 2, "writemostly" - 3, "blocked" - 4, "spare" - 5, "write_error" - 6, "want_replacement" - 7, "replacement" - 8.                                                                            |
| Errors  | gauge  | /sys/block/\[mdN\]/md/\[dev\]errors | An approximate count of read errors that have been detected on this device but have not caused the device to be evicted from the array (either because they were corrected or because they happened while the array was read-only).  |

Details:
* [https://www.kernel.org/doc/html/latest/admin-guide/md.html](https://www.kernel.org/doc/html/latest/admin-guide/md.html)
## Configuration
```json
{
  "collectors": {
    "MDStat": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector