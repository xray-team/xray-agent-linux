# NetDevStatus
## Description
The NetDevStatus collector picks information about the state of network interfaces.  
The following metrics are available for each interface:

| Name      | Type    | Data source                              | Description                                                                                                                                                                                                                                               |
|-----------|---------|------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| OperState | gauge   | /sys/class/net/\[iface\]/operstate       | The oper state of interface: 1 - "up", 2 - "lowerlayerdown", 3 - "dormant", 4 - "down", 5 - "unknown", 6 - "testing", 7 - "notpresent". Details: [https://tools.ietf.org/html/rfc2863#section-3.1.14](https://tools.ietf.org/html/rfc2863#section-3.1.14) |
| LinkFlaps | counter | /sys/class/net/\[iface\]/carrier_changes | The number of link flaps.                                                                                                                                                                                                                                 |
| Speed     | gauge   | /sys/class/net/\[iface\]/speed           | The link speed in MB/s.                                                                                                                                                                                                                                   |
| Duplex    | gauge   | /sys/class/net/\[iface\]/duplex          | Duplex: 1 - "full", 2 - "half", 3 - "unknown". Details: https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-class-net.                                                                                                                             |
| MTU       | gauge   | /sys/class/net/\[iface\]/mtu             | The MTU value for the interface.                                                                                                                                                                                                                          |
## Configuration
```json
{
  "collectors": {
    "NetDevStatus": {
      "enabled": true,
      "excludeWireless": true,
      "excludeByName": []
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"excludeWireless"**
    * **true** - exclude wireless interfaces
    * **false** - do not exclude wireless interfaces
* **"excludeByName"** - list of interface names to exclude. Possible states: "unknown", "notpresent", "down", "lowerlayerdown", "testing", "dormant", "up" (Details: [https://tools.ietf.org/html/rfc2863#section-3.1.14](https://tools.ietf.org/html/rfc2863#section-3.1.14))
