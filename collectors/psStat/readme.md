# psStat
## Description
The PSStat collector picks CPU and memory usage statistics for given applications (processes).
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

## Configuration
```json
{
  "collectors": {
    "PSStat": {
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