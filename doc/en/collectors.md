[Russian version](../ru/collectors.md)

---
# Collectors

List of collectors:
- [uptime](collectors.md#uptime)
- [loadAvg](collectors.md#loadavg)
- [ps](collectors.md#ps)
- [psStat](collectors.md#psstat)
- [stat](collectors.md#stat)
- [cpuInfo](collectors.md#cpuinfo)
- [memoryInfo](collectors.md#memoryinfo)
- [diskStat](collectors.md#diskstat)
- [diskSpace](collectors.md#diskspace)
- [mdStat](collectors.md#mdstat)
- [netDev](collectors.md#netdev)
- [netDevStatus](collectors.md#netdevstatus)
- [netStat](collectors.md#netstat)
- [netSNMP](collectors.md#netsnmp)
- [netSNMP6](collectors.md#netsnmp6)
- [netARP](collectors.md#netarp)
- [wireless](collectors.md#wireless)
- [entropy](collectors.md#entropy)
- [nginxStubStatus](collectors.md#nginxstubstatus)
- [cmd](collectors.md#cmd)

## uptime
### Description
The uptime collector picks the following metrics from /proc/uptime:

| Name   | Type    | Data source  | Description                                                                                                                                             |
|--------|---------|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| Uptime | counter | /proc/uptime | The total number of seconds the system has been up.                                                                                                     |
| Idle   | counter | /proc/uptime | The sum of how much time each core has spent idle, in seconds. This value may be greater than the overall system uptime on systems with multiple cores. |

### Configuration
```json
{
  "collectors": {
    "uptime": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## loadAvg
### Description
The loadAvg collector picks the following metrics from /proc/loadavg:

| Name                     | Type   | Data source   | Description                                                                                             |
|--------------------------|--------|---------------|---------------------------------------------------------------------------------------------------------|
| Last                     | gauge  | /proc/loadavg | The system load during the last one-minute period.                                                      |
| Last5m                   | gauge  | /proc/loadavg | The system load during the last five-minute period.                                                     |
| Last15m                  | gauge  | /proc/loadavg | The system load during the last fifteen-minute period.                                                  |
| KernelSchedulingEntities | gauge  | /proc/loadavg | The total number of kernel scheduling entities (processes, threads) that currently exist on the system. |

### Configuration
```json
{
  "collectors": {
    "loadAvg": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## ps
### Description
The ps collector picks the following metrics:

| Name             | Type  | Data source             | Description                                                     |
|------------------|-------|-------------------------|-----------------------------------------------------------------|
| Count            | gauge | /proc/\[pid\]/status    | The process count.                                              |
| Limit            | gauge | /sys/kernel/pid_max     | The maximum number of processes.                                |
| InStateRunning   | gauge | /proc/\[pid\]/status    | The number of processes in running state.                       |
| InStateIdle      | gauge | /proc/\[pid\]/status    | The number of processes in idle state.                          |
| InStateSleeping  | gauge | /proc/\[pid\]/status    | The number of processes in sleeping state.                      |
| InStateDiskSleep | gauge | /proc/\[pid\]/status    | The number of processes in disk sleep state.                    |
| InStateStopped   | gauge | /proc/\[pid\]/status    | The number of processes in stopped state.                       |
| InStateZombie    | gauge | /proc/\[pid\]/status    | The number of processes in zombie state.                        |
| InStateDead      | gauge | /proc/\[pid\]/status    | The number of processes in dead state.                          |
| Threads          | gauge | /proc/\[pid\]/status    | The threads count.                                              |
| ThreadsLimit     | gauge | /sys/kernel/threads-max | The maximum number of threads that can be created using fork(). |

### Configuration
```json
{
  "collectors": {
    "ps": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## psStat
### Description
The psStat collector picks CPU and memory usage statistics for given applications (processes).
The following metrics are collected for each application:

| Name               | Type  | Data source        | Description                    |
|--------------------|-------|--------------------|--------------------------------|
| Processes          | gauge | /proc/\[pid\]/stat | The number of processes.       |
| Threads            | gauge | /proc/\[pid\]/stat | The number of threads.         |
| System             | gauge | /proc/\[pid\]/stat | CPU usage (system).            |
| User               | gauge | /proc/\[pid\]/stat | CPU usage (user).              |
| Guest              | gauge | /proc/\[pid\]/stat | CPU usage (guest).             |
| Total              | gauge | /proc/\[pid\]/stat | CPU usage (system+user+guest). |
| ResidentMemorySize | gauge | /proc/\[pid\]/stat | Resident memory size in bytes. |
| VirtualMemorySize  | gauge | /proc/\[pid\]/stat | Virtual memory size in bytes.  |

### Configuration
```json
{
  "collectors": {
    "psStat": {
      "enabled": true,
      "collectPerPidStat": false,
      "processList": ["xray-agent"]
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **collectPerPidStat**
    * **true** - enable detailed (per process) metrics
    * **false** - disable detailed metrics. Only aggregated metrics will be collected
* **"processList"** - list of process names

## stat
### Description
The stat collector parses the content of /proc/stat and picks CPU usage statistics (total and per CPU) and other kernel activity metrics.
System metrics:

| Name             | Type    | Data source    | Description                                                           |
|------------------|---------|----------------|-----------------------------------------------------------------------|
| Intr             | counter | /proc/stat     | The total number of all interrupts serviced.                          |
| Ctxt             | counter | /proc/stat     | The total number of context switches across all CPUs.                 |
| Btime            | counter | /proc/stat     | The time at which the system booted, in seconds since the Unix epoch. |
| Processes        | counter | /proc/stat     | The number of processes and threads created.                          |
| ProcessesRunning | gauge   | /proc/stat     | The number of processes currently running on CPUs.                    |
| ProcessesBlocked | gauge   | /proc/stat     | The number of processes currently blocked.                            |

CPU usage metrics (total and per CPU):

| Name      | Type    | Data source  | Description                                                                                  |
|-----------|---------|--------------|----------------------------------------------------------------------------------------------|
| User      | counter | /proc/stat   | The amount of CPU time in milliseconds used by user space processes (in User mode).          |
| Nice      | counter | /proc/stat   | The amount of CPU time in milliseconds used by processes in Nice mode.                       |
| System    | counter | /proc/stat   | The amount of CPU time in milliseconds used by the kernel (in System mode).                  |
| Idle      | counter | /proc/stat   | The amount of CPU time in milliseconds used in Idle mode.                                    |
| IOwait    | counter | /proc/stat   | The amount of CPU time in milliseconds used in IOwait mode.                                  |
| IRQ       | counter | /proc/stat   | The amount of CPU time in milliseconds used for serving the hardware interrupts.             |
| SoftIRQ   | counter | /proc/stat   | The amount of CPU time in milliseconds used for serving the software interrupts.             |
| Steal     | counter | /proc/stat   | The amount of CPU time in milliseconds spent waiting for a virtual CPU in a virtual machine. |
| Guest     | counter | /proc/stat   | The amount of CPU time in milliseconds used by guest OS.                                     |
| GuestNice | counter | /proc/stat   | The amount of CPU time in milliseconds used by guest OS in Nice mode.                        |

Softirq metrics:

| Name    | Type    | Data source  | Description                         |
|---------|---------|--------------|-------------------------------------|
| Total   | counter | /proc/stat   | The total number of all interrupts. |
| Hi      | counter | /proc/stat   | The number of HI interrupts.        |
| Timer   | counter | /proc/stat   | The number of TIMER interrupts.     |
| NetTx   | counter | /proc/stat   | The number of NET_TX interrupts.    |
| NetRx   | counter | /proc/stat   | The number of NET_RX interrupts.    |
| Block   | counter | /proc/stat   | The number of BLOCK interrupts.     |
| IRQPoll | counter | /proc/stat   | The number of IRQ_POLL interrupts.  |
| Tasklet | counter | /proc/stat   | The number of TASKLET interrupts.   |
| Sched   | counter | /proc/stat   | The number of SCHED interrupts.     |
| HRTimer | counter | /proc/stat   | The number of HRTIMER interrupts.   |
| RCU     | counter | /proc/stat   | The number of RCU interrupts.       |

### Configuration
```json
{
  "collectors": {
    "stat": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## cpuInfo
### Description
The cpuInfo collector picks the following metrics for each CPU:

| Name   | Type  | Data source   | Description                               |
|--------|-------|---------------|-------------------------------------------|
| MHz    | gauge | /proc/cpuinfo | The speed in megahertz for the processor. |

### Configuration
```json
{
  "collectors": {
    "cpuInfo": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## memoryInfo
### Description
The memoryInfo collector picks the following memory usage metrics:

| Name      | Type  | Data source   | Description                                              |
|-----------|-------|---------------|----------------------------------------------------------|
| Total     | gauge | /proc/meminfo | Total RAM size in kB.                                    |
| Free      | gauge | /proc/meminfo | Free memory in kB.                                       |
| Available | gauge | /proc/meminfo | Memory in kB is available for starting new applications. |
| Used      | gauge | /proc/meminfo | Total - Free.                                            |
| Buffers   | gauge | /proc/meminfo | Memory in buffer cache (in kB).                          |
| Cached    | gauge | /proc/meminfo | Memory in the pagecache (in kB).                         |
| SwapTotal | gauge | /proc/meminfo | Total SWAP space available in kB.                        |
| SwapFree  | gauge | /proc/meminfo | The remaining swap space available in kB                 |

### Configuration
```json
{
  "collectors": {
    "memoryInfo": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## diskStat
### Description
The diskStat collector picks the following metrics for each disk:

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
### Configuration
```json
{
  "collectors": {
    "diskStat": {
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

## diskSpace
### Description
The diskSpace collector picks the following metrics for each disk partition:

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
### Configuration
```json
{
  "collectors": {
    "diskSpace": {
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

## mdStat
### Description
The mdStat collector picks information about the state of the linux software RAID. The following metrics are collected for each md-device:

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
### Configuration
```json
{
  "collectors": {
    "mdStat": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## netDev
### Description
The netDev collector parses the content of /proc/net/dev and picks statistics for each network interface.  
A description of all counters: [https://www.kernel.org/doc/html/latest/networking/statistics.html#struct-rtnl-link-stats64](https://www.kernel.org/doc/html/latest/networking/statistics.html#struct-rtnl-link-stats64)
### Configuration
```json
{
  "collectors": {
    "netDev": {
      "enabled": true,
      "excludeLoopbacks": true,
      "excludeWireless": false,
      "excludeBridges": false,
      "excludeVirtual": false,
      "excludeByName": [],
      "excludeByOperState": []
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"excludeLoopbacks"**
    * **true** - exclude loopback interfaces
    * **false** - do not exclude loopback interfaces
* **"excludeWireless"**
    * **true** - exclude wireless interfaces
    * **false** - do not exclude wireless interfaces
* **"excludeBridges"**
    * **true** - exclude bridges
    * **false** - do not exclude bridges
* **"excludeVirtual"**
    * **true** - exclude virtual interfaces
    * **false** - do not exclude virtual interfaces
* **"excludeByName"** - list of interface names to exclude
* **"excludeByOperState"** - list of interface states to exclude.

## netDevStatus
### Description
The netDevStatus collector picks information about the state of network interfaces.  
The following metrics are available for each interface:

| Name      | Type    | Data source                              | Description                                                                                                                                                                                                                                               |
|-----------|---------|------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| OperState | gauge   | /sys/class/net/\[iface\]/operstate       | The oper state of interface: 1 - "up", 2 - "lowerlayerdown", 3 - "dormant", 4 - "down", 5 - "unknown", 6 - "testing", 7 - "notpresent". Details: [https://tools.ietf.org/html/rfc2863#section-3.1.14](https://tools.ietf.org/html/rfc2863#section-3.1.14) |
| LinkFlaps | counter | /sys/class/net/\[iface\]/carrier_changes | The number of link flaps.                                                                                                                                                                                                                                 |
| Speed     | gauge   | /sys/class/net/\[iface\]/speed           | The link speed in MB/s.                                                                                                                                                                                                                                   |
| Duplex    | gauge   | /sys/class/net/\[iface\]/duplex          | Duplex: 1 - "full", 2 - "half", 3 - "unknown". Details: https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-class-net.                                                                                                                             |
| MTU       | gauge   | /sys/class/net/\[iface\]/mtu             | The MTU value for the interface.                                                                                                                                                                                                                          |
### Configuration
```json
{
  "collectors": {
    "netDev": {
      "enabled": true,
      "excludeWireless": true,
      "excludeByName": []
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"excludeWireless"**
    * **true** - exclude wireless interfaces
    * **false** - do not exclude wireless interfaces
* **"excludeByName"** - list of interface names to exclude. Possible states: "unknown", "notpresent", "down", "lowerlayerdown", "testing", "dormant", "up" (Details: [https://tools.ietf.org/html/rfc2863#section-3.1.14](https://tools.ietf.org/html/rfc2863#section-3.1.14))

## netStat
### Description
The netStat collector parses the content of /proc/net/netstat and picks the network statistics of the kernel.
The number of metrics may differ depending on the version and settings of the Linux kernel.  
A detailed description of the metrics: [https://www.kernel.org/doc/html/latest/networking/snmp_counter.html](https://www.kernel.org/doc/html/latest/networking/snmp_counter.html)
### Configuration
```json
{
  "collectors": {
    "netStat": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## netSNMP
### Description
The netSNMP collector parses the content of /proc/net/snmp and picks the network protocols (IP, ICMP, TCP, UDP) statistics.  
A detailed description of the metrics: [https://www.kernel.org/doc/html/latest/networking/snmp_counter.html](https://www.kernel.org/doc/html/latest/networking/snmp_counter.html)
### Configuration
```json
{
  "collectors": {
    "netSNMP": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## netSNMP6
### Description
The netSNMP6 collector parses the content of /proc/net/snmp6.  
This collector is IPv6 version of netSNMP collector.
### Configuration
```json
{
  "collectors": {
    "netSNMP6": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## netARP
### Description
The netARP collector parses the content of /proc/net/arp and picks the following metrics for each network interface:

| Name              | Type  | Data source   | Description                                 |
|-------------------|-------|---------------|---------------------------------------------|
| Entries           | gauge | /proc/net/arp | The total number of ARP entries.            |
| IncompleteEntries | gauge | /proc/net/arp | The total number of incomplete ARP entries. |

### Configuration
```json
{
  "collectors": {
    "netARP": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## wireless
### Description
The wireless collector picks metrics for wireless interfaces by using iwconfig utility.

| Name               | Type    | Data source     | Description                                                                                                                               |
|--------------------|---------|-----------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| Frequency          | gauge   | iwconfig        | Wireless channel.                                                                                                                         |
| BitRate            | gauge   | iwconfig        | Wireless speed.                                                                                                                           |
| TxPower            | gauge   | iwconfig        | TX power in dBm.                                                                                                                          |
| LinkQuality        | gauge   | iwconfig        | The overall quality of the link.                                                                                                          |
| LinkQualityLimit   | gauge   | iwconfig        | The link quality limit.                                                                                                                   |
| SignalLevel        | gauge   | iwconfig        | Received signal strength.                                                                                                                 |
| RxInvalidNwid      | counter | iwconfig        | The number of packets received with a different NWID or ESSID.                                                                            |
| RxInvalidCrypt     | counter | iwconfig        | The number of packets that the hardware was unable to decrypt.                                                                            |
| RxInvalidFrag      | counter | iwconfig        | The number of packets for which the hardware was not able to properly re-assemble the link layer fragments (most likely one was missing). |
| TxExcessiveRetries | counter | iwconfig        | The number of packets that the hardware failed to deliver.                                                                                |
| InvalidMisc        | counter | iwconfig        | The number of other packets lost in relation with specific wireless operations.                                                           |
| MissedBeacon       | counter | iwconfig        | The number of periodic beacons from the Cell or the Access Point we have missed.                                                          |

Details: man iwconfig

### Configuration
```json
{
  "collectors": {
    "wireless": {
      "enabled": true,
      "excludeByName": [],
      "excludeByOperState": ["down"]
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"excludeByName"** - list of interface names to exclude
* **"excludeByOperState"** - list of interface states to exclude. Possible states: "unknown", "notpresent", "down", "lowerlayerdown", "testing", "dormant", "up" (Details: [https://tools.ietf.org/html/rfc2863#section-3.1.14](https://tools.ietf.org/html/rfc2863#section-3.1.14))

## entropy
### Description
The entropy collector picks the following metrics:

| Name      | Type  | Data source                           | Description                                                          |
|-----------|-------|---------------------------------------|----------------------------------------------------------------------|
| Available | gauge | /proc/sys/kernel/random/entropy_avail | The measure of bits currently available to be read from /dev/random. |

### Configuration
```json
{
  "collectors": {
    "entropy": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector

## nginxStubStatus
### Description
The nginxStubStatus collector picks the basic status information about nginx server from ngx_http_stub_status_module.  
The following status information metrics are available:

| Name     | Type    | Data source                 | Description                                                                               |
|----------|---------|-----------------------------|-------------------------------------------------------------------------------------------|
| Active   | gauge   | ngx_http_stub_status_module | The current number of active client connections including Waiting connections.            |
| Accepts  | counter | ngx_http_stub_status_module | The total number of accepted client connections.                                          |
| Handled  | counter | ngx_http_stub_status_module | The total number of handled connections.                                                  |
| Requests | counter | ngx_http_stub_status_module | The total number of client requests.                                                      |
| Reading  | gauge   | ngx_http_stub_status_module | The current number of connections where nginx is reading the request header.              |
| Writing  | gauge   | ngx_http_stub_status_module | The current number of connections where nginx is writing the response back to the client. |
| Waiting  | gauge   | ngx_http_stub_status_module | The current number of idle client connections waiting for a request.                      |

### Limitations
Before using the nginxStubStatus collector, you must first enable ngx_http_stub_status_module.  
More details in the nginx documentation: [https://nginx.org/en/docs/http/ngx_http_stub_status_module.html](https://nginx.org/en/docs/http/ngx_http_stub_status_module.html)

### Configuration
```json
{
  "collectors": {
    "nginxStubStatus": {
      "enabled": true,
      "endpoint": "http://127.0.0.1/basic_status",
      "timeout": 5
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"endpoint"** - ngx_http_stub_status_module url (location)
* **"timeout"** - timeout in seconds

## cmd
### Description
The cmd collector allows you to execute shell commands or scripts and record the results as separate metrics.
Thus, it is possible to quickly and easily extend the functionality of the xray-agent with your own metrics.
### Configuration
```json
{
  "collectors": {
    "cmd": {
      "enabled": true,
      "timeout": 1,
      "metrics": [
        {
          "names": ["metricName1", "metricName2", "metricNameN"],
          "delimiter": " ",
          "pipeline": [
            ["command1", "argument1ForCommand1", "argument2ForCommand1", "argumentNForCommand1"],
            ["command2", "argument1ForCommand2", "argument2ForCommand2", "argumentNForCommand2"],
            ["commandN", "argument1ForCommandN", "argument2ForCommandN", "argumentNForCommandN"]
          ],
          "attributes": [
            {
              "name": "AttributeName",
              "value": "AttributeValue"
            }
          ]
        }
      ]
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"timeout"** - timeout for pipeline in seconds. Min: 1, max: 120
* **"metrics"** - command configuration
    * **"names"** - list of metric names. The number of elements in this list must be equal to the number of values returned by the command specified in the "pipeline". To ignore (skip) unnecessary values, specify "-" instead of name.
    * **"delimiter"** - delimiter string
    * **"pipeline"** - list of commands  
      Examples:
        * Entropy  
          *shell*:
          ```shell
          cat /proc/sys/kernel/random/entropy_avail
          ```
          *equivalent xray cmd pipeline*:
          ```json
          "pipeline": [
            ["cat", "/proc/sys/kernel/random/entropy_avail"]
          ]
          ```
        * OOM killer  
          *shell*:
          ```shell
          cat /proc/vmstat | grep oom_kill
          ```
          *equivalent xray cmd pipeline*:
          ```json
          "pipeline": [
            ["cat", "/proc/vmstat"],
            ["grep", "oom_kill"]
          ]
          ```
        * ARP Limits  
          *shell*:
          ```shell
          sysctl net.ipv4.neigh.default.gc_thresh1 net.ipv4.neigh.default.gc_thresh2 net.ipv4.neigh.default.gc_thresh3 | awk '{print $3}' ORS=' '
          ```
          *equivalent xray cmd pipeline*:
          ```json
          "pipeline": [
            ["sysctl", "net.ipv4.neigh.default.gc_thresh1", "net.ipv4.neigh.default.gc_thresh2", "net.ipv4.neigh.default.gc_thresh3"],
            ["awk", "{print $3}", "ORS= "]
          ]
          ```
* **"attributes"** - additional attributes

**Example 1.**
For example, if you want to save the output of the following command as a metric named EntropyAvail:
```shel
cat /proc/sys/kernel/random/entropy_avail
```
add to xray-agent config:
```json
{
  "collectors": {
    "cmd": {
      "enabled": true,
      "metrics": [
        {
          "names": ["EntropyAvail"],
          "pipeline": [
            ["cat", "/proc/sys/kernel/random/entropy_avail"]
          ],
          "delimiter": " "
        }
      ]
    }
  }
}
```
**Example 2.** For example, if you need to save the output of the following command as 3 metrics:
```shel
sysctl net.ipv4.neigh.default.gc_thresh1 net.ipv4.neigh.default.gc_thresh2 net.ipv4.neigh.default.gc_thresh3 | awk '{print $3}' ORS=' '
```
add to xray-agent config:
```json
{
  "collectors": {
    "cmd": {
      "metrics": [
        {
          "names": ["GcThresh1", "GcThresh2", "GcThresh3"],
          "pipeline": [
            ["sysctl", "net.ipv4.neigh.default.gc_thresh1", "net.ipv4.neigh.default.gc_thresh2", "net.ipv4.neigh.default.gc_thresh3"],
            ["awk", "{print $3}", "ORS= "]
          ],
          "delimiter": " ",
          "attributes": [
            {
              "name": "Set",
              "value": "IPv4NeighbourTable"
            }
          ]
        }
      ]
    }
  }
}
```