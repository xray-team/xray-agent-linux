# CPUInfo
## Description
The CPUInfo collector picks the following metrics for each CPU:

| Name   | Type  | Data source   | Description                               |
|--------|-------|---------------|-------------------------------------------|
| MHz    | gauge | /proc/cpuinfo | The speed in megahertz for the processor. |

## Configuration
```json
{
  "collectors": {
    "CPUInfo": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
