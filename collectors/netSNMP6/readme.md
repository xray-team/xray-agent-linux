# NetSNMP6
### Description
The NetSNMP6 collector parses the content of /proc/net/snmp6.  
This collector is IPv6 version of netSNMP collector.
### Configuration
```json
{
  "collectors": {
    "NetSNMP6": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector