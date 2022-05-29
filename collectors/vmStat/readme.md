# VMStat
## Description
The VMStat collector picks the following metrics from /proc/vmstat:

| Name           | Type    | Data source  | Description                   |
|----------------|---------|--------------|-------------------------------|
| PgPgIn         | counter | /proc/vmstat | The number of pages paged in  |
| PgPgOut        | counter | /proc/vmstat | The number of pages paged out |
| PSwpIn         | counter | /proc/vmstat |                               |
| PSwpOut        | counter | /proc/vmstat |                               |
| PgFault        | counter | /proc/vmstat |                               |
| PgMajFault     | counter | /proc/vmstat |                               |
| PgFree         | counter | /proc/vmstat |                               |
| PgActivate     | counter | /proc/vmstat |                               |
| PgDeactivate   | counter | /proc/vmstat |                               |
| PgLazyFree     | counter | /proc/vmstat |                               |
| PgLazyFreed    | counter | /proc/vmstat |                               |
| PgRefill       | counter | /proc/vmstat |                               |
| NumaHit        | counter | /proc/vmstat |                               |
| NumaMiss       | counter | /proc/vmstat |                               |
| NumaForeign    | counter | /proc/vmstat |                               |
| NumaInterleave | counter | /proc/vmstat |                               |
| NumaLocal      | counter | /proc/vmstat |                               |
| NumaOther      | counter | /proc/vmstat |                               |
| OOMKill        | counter | /proc/vmstat |                               |

## Configuration
```json
{
  "collectors": {
    "VMStat": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector