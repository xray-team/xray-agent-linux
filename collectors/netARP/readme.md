# NetARP
## Description
The NetARP collector parses the content of /proc/net/arp and picks the following metrics for each network interface:

| Name              | Type  | Data source   | Description                                 |
|-------------------|-------|---------------|---------------------------------------------|
| Entries           | gauge | /proc/net/arp | The total number of ARP entries.            |
| IncompleteEntries | gauge | /proc/net/arp | The total number of incomplete ARP entries. |

## Configuration
```json
{
  "collectors": {
    "NetARP": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
