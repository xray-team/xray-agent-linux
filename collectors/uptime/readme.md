# Uptime
## Description
The Uptime collector picks the following metrics from /proc/uptime:

| Name   | Type    | Data source  | Description                                                                                                                                             |
|--------|---------|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| Uptime | counter | /proc/uptime | The total number of seconds the system has been up.                                                                                                     |
| Idle   | counter | /proc/uptime | The sum of how much time each core has spent idle, in seconds. This value may be greater than the overall system uptime on systems with multiple cores. |

## Configuration
```json
{
  "collectors": {
    "Uptime": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector