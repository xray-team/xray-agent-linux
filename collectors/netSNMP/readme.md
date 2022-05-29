# netSNMP
## Description
The NetSNMP collector parses the content of /proc/net/snmp and picks the network protocols (IP, ICMP, TCP, UDP) statistics.  
A detailed description of the metrics: [https://www.kernel.org/doc/html/latest/networking/snmp_counter.html](https://www.kernel.org/doc/html/latest/networking/snmp_counter.html)
## Configuration
```json
{
  "collectors": {
    "NetSNMP": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector