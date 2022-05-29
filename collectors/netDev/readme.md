# NetDev
## Description
The NetDev collector parses the content of /proc/net/dev and picks statistics for each network interface.  
A description of all counters: [https://www.kernel.org/doc/html/latest/networking/statistics.html#struct-rtnl-link-stats64](https://www.kernel.org/doc/html/latest/networking/statistics.html#struct-rtnl-link-stats64)
## Configuration
```json
{
  "collectors": {
    "NetDev": {
      "enabled": true,
      "excludeLoopbacks": true,
      "excludeWireless": false,
      "excludeBridges": false,
      "excludeVirtual": false,
      "excludeByName": [],
      "excludeByOperState": []
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"excludeLoopbacks"**
    * **true** - exclude loopback interfaces
    * **false** - do not exclude loopback interfaces
* **"excludeWireless"**
    * **true** - exclude wireless interfaces
    * **false** - do not exclude wireless interfaces
* **"excludeBridges"**
    * **true** - exclude bridges
    * **false** - do not exclude bridges
* **"excludeVirtual"**
    * **true** - exclude virtual interfaces
    * **false** - do not exclude virtual interfaces
* **"excludeByName"** - list of interface names to exclude
* **"excludeByOperState"** - list of interface states to exclude.