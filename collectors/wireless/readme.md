# Wireless
## Description
The Wireless collector picks metrics for wireless interfaces by using iwconfig utility.

| Name               | Type    | Data source     | Description                                                                                                                               |
|--------------------|---------|-----------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| Frequency          | gauge   | iwconfig        | Wireless channel.                                                                                                                         |
| BitRate            | gauge   | iwconfig        | Wireless speed.                                                                                                                           |
| TxPower            | gauge   | iwconfig        | TX power in dBm.                                                                                                                          |
| LinkQuality        | gauge   | iwconfig        | The overall quality of the link.                                                                                                          |
| LinkQualityLimit   | gauge   | iwconfig        | The link quality limit.                                                                                                                   |
| SignalLevel        | gauge   | iwconfig        | Received signal strength.                                                                                                                 |
| RxInvalidNwid      | counter | iwconfig        | The number of packets received with a different NWID or ESSID.                                                                            |
| RxInvalidCrypt     | counter | iwconfig        | The number of packets that the hardware was unable to decrypt.                                                                            |
| RxInvalidFrag      | counter | iwconfig        | The number of packets for which the hardware was not able to properly re-assemble the link layer fragments (most likely one was missing). |
| TxExcessiveRetries | counter | iwconfig        | The number of packets that the hardware failed to deliver.                                                                                |
| InvalidMisc        | counter | iwconfig        | The number of other packets lost in relation with specific wireless operations.                                                           |
| MissedBeacon       | counter | iwconfig        | The number of periodic beacons from the Cell or the Access Point we have missed.                                                          |

Details: man iwconfig

## Configuration
```json
{
  "collectors": {
    "Wireless": {
      "enabled": true,
      "excludeByName": [],
      "excludeByOperState": ["down"]
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"excludeByName"** - list of interface names to exclude
* **"excludeByOperState"** - list of interface states to exclude. Possible states: "unknown", "notpresent", "down", "lowerlayerdown", "testing", "dormant", "up" (Details: [https://tools.ietf.org/html/rfc2863#section-3.1.14](https://tools.ietf.org/html/rfc2863#section-3.1.14))
