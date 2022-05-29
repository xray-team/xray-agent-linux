# NetStat
## Description
The NetStat collector parses the content of /proc/net/netstat and picks the network statistics of the kernel.
The number of metrics may differ depending on the version and settings of the Linux kernel.  
A detailed description of the metrics: [https://www.kernel.org/doc/html/latest/networking/snmp_counter.html](https://www.kernel.org/doc/html/latest/networking/snmp_counter.html)
## Configuration
```json
{
  "collectors": {
    "NetStat": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector