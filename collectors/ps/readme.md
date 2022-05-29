# PS collector
## Description
The PS collector picks the following metrics:

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

## Configuration
```json
{
  "collectors": {
    "PS": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector