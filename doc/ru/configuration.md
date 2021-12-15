# Конфигурация
Конфигурация xray-agent производится путем задания параметров в конфигурационном файле. Логически конфигурация состоит из трех секций:
  - agent - [общая конфигурация агента](configuration.md#параметры-общей-конфигурации-агента)
  - collectors - [конфигурация коллекторов](configuration.md#конфигурация-коллекторов)
  - tsDB - конфигурация базы данных
## Пример конфигурационного файла
```json
{
  "agent": {
    "getStatIntervalSec": 60,
    "hostAttributes": [
      {
        "name": "Source",
        "value": "xray"
      }
    ]
  },
  "collectors": {
    "rootPath": "/",
    "enableSelfMetrics": true,
    "uptime": {
      "enabled": true
    },
    "loadAvg": {
      "enabled": true
    },
    "ps": {
      "enabled": true
    },
    "psStat": {
      "enabled": true,
      "collectPerPidStat": false,
      "processList": ["xray-agent"]
    },
    "stat": {
      "enabled": true
    },
    "cpuInfo": {
      "enabled": true
    },
    "memoryInfo": {
      "enabled": true
    },
    "netARP": {
      "enabled": true
    },
    "netStat": {
      "enabled": true
    },
    "netSNMP": {
      "enabled": true
    },
    "netSNMP6": {
      "enabled": true
    },
    "netDev": {
      "enabled": true,
      "excludeLoopbacks": true,
      "excludeWireless": false,
      "excludeBridges": false,
      "excludeVirtual": false,
      "excludeByName": [],
      "excludeByOperState": []
    },
    "netDevStatus": {
      "enabled": true,
      "excludeWireless": true,
      "excludeByName": []
    },
    "wireless": {
      "enabled": true,
      "excludeByName": [],
      "excludeByOperState": ["down"]
    },
    "diskStat": {
      "enabled": true,
      "diskTypes": [
        8,
        9
      ],
      "excludePartitions": false,
      "excludeByName": []
    },
    "diskSpace": {
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
    "mdStat": {
      "enabled": true
    },
    "cmd": {
      "enabled": false,
      "timeout": 10,
      "metrics": []
    },
    "nginxStubStatus": {
      "enabled": true,
      "endpoint": "http://127.0.0.1/basic_status",
      "timeout": 5
    },
    "entropy": {
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

## Параметры общей конфигурации агента
### hostAttributes
hostAttributes - структура, которая описывает административные свойства хоста.  
Данные атрибуты задаются администратором и могут служить для группировки хостов по типу, назначению, размещению и т.д. Количество атрибутов может быть произвольным.  
Если для хранения данных используется graphite в режиме "tags" (с поддержкой тегов) - hostAttributes будут добавлены ко всем метрикам хоста в виде тегов.  
Если для хранения данных используется graphite в режиме "tree" - hostAttributes будут добавлены ко всем метрикам хоста в дерево перед хостнеймом.  

Несколько примеров задания hostAttributes:
```json
{
  "agent": {
    "hostAttributes": [
      {
        "name": "Source",
        "value": "xray"
      }
    ]
  }
}
```
```json
{
  "agent": {
    "hostAttributes": [
      {
        "name": "Source",
        "value": "xray"
      },
      {
        "name": "Type",
        "value": "Server"
      },
      {
        "name": "Location",
        "value": "Room229"
      }
    ]
  }
}
```
## Конфигурация коллекторов
Подробней о конфигурации каждого коллектора [здесь](collectors.md).