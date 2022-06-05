# Stat
## Description
The Stat collector parses the content of /proc/stat and picks CPU usage statistics (total and per CPU) and other kernel activity metrics.
System metrics:

| Name             | Type    | Data source    | Description                                                           |
|------------------|---------|----------------|-----------------------------------------------------------------------|
| Intr             | counter | /proc/stat     | The total number of all interrupts serviced.                          |
| Ctxt             | counter | /proc/stat     | The total number of context switches across all CPUs.                 |
| BootTime         | counter | /proc/stat     | The time at which the system booted, in seconds since the Unix epoch. |
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

## Configuration
```json
{
  "collectors": {
    "Stat": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector