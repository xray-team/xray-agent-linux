# Configuration
## Configuration file example
```json
{
  "agent": {
    "getStatIntervalSec": 60,
    "enableSelfMetrics": true,
    "hostAttributes": [
      {
        "name": "Source",
        "value": "xray"
      }
    ],
    "logLevel": "default",
    "logOut": "syslog"
  },
  "collectors": {
    "Uptime": {
      "enabled": true
    },
    "LoadAvg": {
      "enabled": true
    },
    "PS": {
      "enabled": true
    },
    "PSStat": {
      "enabled": true,
      "collectPerPidStat": false,
      "processList": ["xray-agent"]
    },
    "Stat": {
      "enabled": true
    },
    "CPUInfo": {
      "enabled": true
    },
    "MemoryInfo": {
      "enabled": true
    },
    "NetARP": {
      "enabled": true
    },
    "NetStat": {
      "enabled": true
    },
    "NetSNMP": {
      "enabled": true
    },
    "NetSNMP6": {
      "enabled": true
    },
    "NetDev": {
      "enabled": true,
      "excludeLoopbacks": true,
      "excludeWireless": false,
      "excludeBridges": false,
      "excludeVirtual": false,
      "excludeByName": [],
      "excludeByOperState": []
    },
    "NetDevStatus": {
      "enabled": true,
      "excludeWireless": true,
      "excludeByName": []
    },
    "Wireless": {
      "enabled": true,
      "excludeByName": [],
      "excludeByOperState": ["down"]
    },
    "DiskStat": {
      "enabled": true,
      "diskTypes": [
        8,
        9
      ],
      "excludePartitions": false,
      "excludeByName": []
    },
    "DiskSpace": {
      "enabled": true,
      "fsTypes": [
        "ext4",
        "ext3",
        "ext2",
        "xfs",
        "jfs",
        "btrfs"
      ]
    },
    "MDStat": {
      "enabled": true
    },
    "CMD": {
      "enabled": false,
      "timeout": 10,
      "metrics": []
    },
    "Nginx": {
      "enabled": true,
      "endpoint": "http://127.0.0.1/basic_status",
      "timeout": 5
    },
    "Entropy": {
      "enabled": true
    }
  },
  "tsDB": {
    "graphite": {
      "servers": [
        {
          "mode": "tree",
          "address": "127.0.0.1:2003",
          "protocol": "tcp",
          "timeout": 10
        }
      ]
    }
  }
}
```

## Collectors settings
Learn more about the settings for each collector [here](collectors.md).