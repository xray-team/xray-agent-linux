# LoadAvg
## Description
The LoadAvg collector picks the following metrics from /proc/loadavg:

| Name                     | Type   | Data source   | Description                                                                                             |
|--------------------------|--------|---------------|---------------------------------------------------------------------------------------------------------|
| Last                     | gauge  | /proc/loadavg | The system load during the last one-minute period.                                                      |
| Last5m                   | gauge  | /proc/loadavg | The system load during the last five-minute period.                                                     |
| Last15m                  | gauge  | /proc/loadavg | The system load during the last fifteen-minute period.                                                  |
| KernelSchedulingEntities | gauge  | /proc/loadavg | The total number of kernel scheduling entities (processes, threads) that currently exist on the system. |

## Configuration
```json
{
  "collectors": {
    "LoadAvg": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
