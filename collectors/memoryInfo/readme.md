# MemoryInfo
## Description
The MemoryInfo collector picks the following memory usage metrics:

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

## Configuration
```json
{
  "collectors": {
    "MemoryInfo": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector