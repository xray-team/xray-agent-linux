# Конфигурация
Конфигурация xray-agent производится путем задания параметров в конфигурационном файле. Логически конфигурация состоит из трех секций:
  - agent - [глобальные настройки агента](configuration.md#глобальные-настройки-агента)
  - collectors - [настройки коллекторов](configuration.md#настройки-коллекторов)
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

## Глобальные настройки агента
### getStatIntervalSec
Параметр getStatIntervalSec задает интервал сбора метрик в секундах.
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
## Настройки коллекторов
Подробней о настройках каждого коллектора [здесь](collectors.md).
### Отключение коллекторов
В некоторых случаях, может быть целесообразно отключить часть коллекторов.  
Например, если метрики, которые собирает коллектор Вам не интересны.  
Для отключения любого из коллекторов - установите параметру **"enabled"** значение false в соответствующей секции конфигурационного файла.
Например, чтобы отключить коллектор ps:
```json
{
  "collectors": {
    "ps": {
      "enabled": false
    }
  }
}
```
### Собственные метрики агента
В xray-agent предусмотрена возможность сбора телеметрии работы агента в виде отдельных метрик.  
Метрики собираются для агента в целом и для каждого коллектора отдельно.

Для агента собираются следующие метрики:

| Имя метрики | Тип   | Описание                                |
|-------------|-------|-----------------------------------------|
| DurationNs  | gauge | Время сбора всех метрик в наносекундах. |
| Metrics     | gauge | Количество всех собранных метрик.       |

Для каждого коллектора собираются следующие метрики:

| Имя метрики    | Тип   | Описание                                                        |
|----------------|-------|-----------------------------------------------------------------|
| DurationNs     | gauge | Время сбора всех метрик коллектором в наносекундах.             |
| Metrics        | gauge | Количество метрик, собранных коллектором.                       |
| CollectorState | gauge | Статус работы коллектора (0 - unknown, 1 - success, 2 - error). |

Чтобы включить эту опцию, добавьте в конфигурационный файл следующий параметр:
```json
{
  "collectors": {
    "enableSelfMetrics": true
  }
}
```