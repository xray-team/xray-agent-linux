# Entropy
## Description
The Entropy collector picks the following metrics:

| Name      | Type  | Data source                           | Description                                                          |
|-----------|-------|---------------------------------------|----------------------------------------------------------------------|
| Available | gauge | /proc/sys/kernel/random/entropy_avail | The measure of bits currently available to be read from /dev/random. |

## Configuration
```json
{
  "collectors": {
    "Entropy": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
