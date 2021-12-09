# Коллекторы
Все метрики собираются с помощью коллекторов.  
Каждый коллектор представляет собой небольшой модуль, который позволяет собирать определенные метрики, имеет собственную конфигурацию и ограничения.  
Любой из коллекторов можно отключить, если метрики, которые он собирает не нужны.

Список коллекторов:
- [uptime](collectors.md#uptime)
- [loadAvg](collectors.md#loadavg)
- [ps](collectors.md#ps)
- [psStat](collectors.md#psstat)
- [stat](collectors.md#stat)
- [cpuInfo](collectors.md#cpuinfo)
- [memoryInfo](collectors.md#memoryinfo)
- [diskStat](collectors.md#diskstat)
- [diskSpace](collectors.md#diskspace)
- [mdStat](collectors.md#mdstat)
- [netDev](collectors.md#netdev)
- [netDevStatus](collectors.md#netdevstatus)
- [netStat](collectors.md#netstat)
- [netSNMP](collectors.md#netsnmp)
- [netSNMP6](collectors.md#netsnmp6)
- [netARP](collectors.md#netarp)
- [entropy](collectors.md#entropy)
- [nginxStubStatus](collectors.md#nginxstubstatus)
- [cmd](collectors.md#cmd)

## uptime
### Описание
Коллектор uptime собирает следующие метрики из /proc/uptime:

| Имя метрики | Тип     | Источник данных | Описание                                                                                                                              |
|-------------|---------|-----------------|---------------------------------------------------------------------------------------------------------------------------------------|
| Uptime      | counter | /proc/uptime    | Время работы системы с момента последней загрузки в секундах.                                                                         |
| Idle        | counter | /proc/uptime    | Сумма времени простоя всех cpu(ядер) в секундах. Для систем с количеством ядер больше 1 данное значение может быть больше чем Uptime. |

### Конфигурация
```json
{
  "collectors": {
    "uptime": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## loadAvg
### Описание
Коллектор loadAvg собирает следующие метрики из /proc/loadavg:

| Имя метрики              | Тип   | Источник данных | Описание                                                                                                 |
|--------------------------|-------|-----------------|----------------------------------------------------------------------------------------------------------|
| Last                     | gauge | /proc/loadavg   | Значение Load Average усредненное за 1 минуту.                                                           |
| Last5m                   | gauge | /proc/loadavg   | Значение Load Average усредненное за 5 минут.                                                            |
| Last15m                  | gauge | /proc/loadavg   | Значение Load Average усредненное за 15 минут.                                                           |
| KernelSchedulingEntities | gauge | /proc/loadavg   | Количество объектов планирования ядра(процессы, потоки), которые в настоящее время существуют в системе. |

### Конфигурация
```json
{
  "collectors": {
    "loadAvg": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## ps
### Описание
Коллектор ps собирает следующую статистику по процессам в системе:

| Имя метрики      | Тип   | Источник данных         | Описание                                      |
|------------------|-------|-------------------------|-----------------------------------------------|
| Count            | gauge | /proc/\[pid\]/status    | Количество запущенных процессов в системе.    |
| Limit            | gauge | /sys/kernel/pid_max     | Максимальное допустимое количество процессов. |
| InStateRunning   | gauge | /proc/\[pid\]/status    | Количество процессов в состоянии running.     |
| InStateIdle      | gauge | /proc/\[pid\]/status    | Количество процессов в состоянии idle.        |
| InStateSleeping  | gauge | /proc/\[pid\]/status    | Количество процессов в состоянии sleeping.    |
| InStateDiskSleep | gauge | /proc/\[pid\]/status    | Количество процессов в состоянии disk sleep.  |
| InStateStopped   | gauge | /proc/\[pid\]/status    | Количество процессов в состоянии stopped.     |
| InStateZombie    | gauge | /proc/\[pid\]/status    | Количество процессов в состоянии zombie.      |
| InStateDead      | gauge | /proc/\[pid\]/status    | Количество процессов в состоянии dead.        |
| Threads          | gauge | /proc/\[pid\]/status    | Количество запущенных потоков.                |
| ThreadsLimit     | gauge | /sys/kernel/threads-max | Максимальное допустимое количество потоков.   |
### Конфигурация
```json
{
  "collectors": {
    "ps": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## psStat
### Описание
Коллектор psStat собирает статистику использования CPU и памяти для указанных в конфигурации приложений. Сбор метрик происходит аналогично тому, как это делает утилита top (top -p <pid>).  
Для каждого приложения собираются следующие метрики:

| Имя метрики        | Тип   | Источник данных         | Описание                                            |
|--------------------|-------|-------------------------|-----------------------------------------------------|
| Processes          | gauge | /proc/\[pid\]/stat      | Количество процессов.                               |
| Threads            | gauge | /proc/\[pid\]/stat      | Количество потоков.                                 |
| System             | gauge | /proc/\[pid\]/stat      | Использование CPU(system).                          |
| User               | gauge | /proc/\[pid\]/stat      | Использование CPU(user).                            |
| Guest              | gauge | /proc/\[pid\]/stat      | Использование CPU(guest).                           |
| Total              | gauge | /proc/\[pid\]/stat      | Использование CPU(system+user+guest).               |
| ResidentMemorySize | gauge | /proc/\[pid\]/stat      | Размер памяти (в байтах) выделенной процессу в RAM. |
| VirtualMemorySize  | gauge | /proc/\[pid\]/stat      | Размер виртуальной памяти в байтах.                 |

### Конфигурация
```json
{
  "collectors": {
    "psStat": {
      "enabled": true,
      "collectPerPidStat": false,
      "processList": ["xray-agent"]
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор
* **collectPerPidStat**
  * **true** - дополнительно будут собираться метрики по каждому процессу
  * **false** - будут собираться только агрегированные метрики.
* **"processList"** - список имен процессов, для которых нужно собирать статистику

## stat
### Описание
Коллектор stat анализирует содержимое /proc/stat и собирает статистику использования CPU (общую и по каждому CPU отдельно), а также другую активность ядра.  
Общие метрики системы:

| Имя метрики      | Тип     | Источник данных | Описание                                                                                   |
|------------------|---------|-----------------|--------------------------------------------------------------------------------------------|
| Intr             | counter | /proc/stat      | Количество обработанных прерываний с момента загрузки (сумма всех обслуженных прерываний). |
| Ctxt             | counter | /proc/stat      | Общее количество переключений контекста на всех процессорах.                               |
| Btime            | counter | /proc/stat      | Время загрузки системы в секундах с эпохи Unix (Unix epoch).                               |
| Processes        | counter | /proc/stat      | Количество созданных процессов с момента загрузки.                                         |
| ProcessesRunning | gauge   | /proc/stat      | Количество процессов, выполняемых в данный момент.                                         |
| ProcessesBlocked | gauge   | /proc/stat      | Количество процессов, заблокированных в данный момент.                                     |

Метрики по каждому CPU и сводные метрики по всем CPU(Total):

| Имя метрики | Тип     | Источник данных | Описание                                                                                                                                                                                                                                                                                                                                    |
|-------------|---------|-----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| User        | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованого в режиме User.                                                                                                                                                                                                                                                             |
| Nice        | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованого в режиме Nice.                                                                                                                                                                                                                                                             |
| System      | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованого в режиме System.                                                                                                                                                                                                                                                           |
| Idle        | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованого в режиме Idle.                                                                                                                                                                                                                                                             |
| IOwait      | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованого в режиме IOwait.                                                                                                                                                                                                                                                           |
| IRQ         | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованого на обслуживание hardware прерываний.                                                                                                                                                                                                                                       |
| SoftIRQ     | counter | /proc/stat      | Количество процессорного времени, использованого на обслуживание software прерываний.                                                                                                                                                                                                                                                       |
| Steal       | counter | /proc/stat      | Количество процессорного времени в миллисекундах, которое было необходимо гостевой виртуальной машине, но не было предоставлено хостом. Большое количество времени Steal указывает на конкуренцию за CPU, например при переподписке CPU. Большое количество времени Steal является индикатором снижения производительности гостевых систем. |
| Guest       | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованого гостевыми операционными системами.                                                                                                                                                                                                                                         |
| GuestNice   | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованого гостевыми операционными системами (в режиме Nice).                                                                                                                                                                                                                         |

Метрики softirq:

| Имя метрики | Тип     | Источник данных | Описание                        |
|-------------|---------|-----------------|---------------------------------|
| Total       | counter | /proc/stat      | Количество всех прерываний.     |
| Hi          | counter | /proc/stat      | Количество прерываний HI.       |
| Timer       | counter | /proc/stat      | Количество прерываний TIMER.    |
| NetTx       | counter | /proc/stat      | Количество прерываний NET_TX.   |
| NetRx       | counter | /proc/stat      | Количество прерываний NET_RX.   |
| Block       | counter | /proc/stat      | Количество прерываний BLOCK.    |
| IRQPoll     | counter | /proc/stat      | Количество прерываний IRQ_POLL. |
| Tasklet     | counter | /proc/stat      | Количество прерываний TASKLET.  |
| Sched       | counter | /proc/stat      | Количество прерываний SCHED.    |
| HRTimer     | counter | /proc/stat      | Количество прерываний HRTIMER.  |
| RCU         | counter | /proc/stat      | Количество прерываний RCU.      |

### Конфигурация
```json
{
  "collectors": {
    "stat": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## cpuInfo
### Описание
Коллектор cpuInfo для каждого CPU собирает следующие метрики:

| Имя метрики | Тип   | Источник данных | Описание                                |
|-------------|-------|-----------------|-----------------------------------------|
| MHz         | gauge | /proc/cpuinfo   | Частота работы процессора (ядра) в МГц. |

### Конфигурация
```json
{
  "collectors": {
    "cpuInfo": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## memoryInfo
### Описание
### Конфигурация
```json
{
  "collectors": {
    "memoryInfo": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## diskStat
### Описание
### Конфигурация

## diskSpace
### Описание
### Конфигурация

## mdStat
### Описание
### Конфигурация

## netDev
### Описание
### Конфигурация

## netDevStatus
### Описание
### Конфигурация

## netStat
### Описание
### Конфигурация
```json
{
  "collectors": {
    "netStat": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## netSNMP
### Описание
Коллектор netSNMP собирает все доступные метрики из /proc/net/snmp. Это статистика по: IP, ICMP, TCP, UDP.  
Подробное описание метрик можно найти здесь: [https://www.kernel.org/doc/html/latest/networking/snmp_counter.html](https://www.kernel.org/doc/html/latest/networking/snmp_counter.html)
### Конфигурация
```json
{
  "collectors": {
    "netSNMP": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## netSNMP6
### Описание
Коллектор netSNMP6 собирает все доступные метрики из /proc/net/snmp6. Данный коллектор собирает аналогичные коллектору netSNMP метрики для IPv6.
### Конфигурация
```json
{
  "collectors": {
    "netSNMP6": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## netARP
### Описание
Коллектор netARP анализирует содержимое /proc/net/arp и для каждого сетевого интерфейса собирает следующие метрики:

| Имя метрики       | Тип   | Источник данных | Описание                                                       |
|-------------------|-------|-----------------|----------------------------------------------------------------|
| Entries           | gauge | /proc/net/arp   | Количество всех ARP записей.                                   |
| IncompleteEntries | gauge | /proc/net/arp   | Количество ARP записей, для которых не удалось получить адрес. |

### Конфигурация
```json
{
  "collectors": {
    "netARP": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## entropy
### Описание
### Конфигурация
```json
{
  "collectors": {
    "entropy": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## nginxStubStatus
### Описание
Коллектор nginxStubStatus собирает базовую информацию о состоянии сервера nginx используя модуль ngx_http_stub_status_module.  
Список доступных метрик:

| Имя метрики | Тип     | Источник данных              | Описание                                                                               |
|-------------|---------|------------------------------|----------------------------------------------------------------------------------------|
| Active      | gauge   | ngx_http_stub_status_module  | Текущее число активных клиентских соединений, включая Waiting-соединения .             |
| Accepts     | counter | ngx_http_stub_status_module  | Суммарное число принятых клиентских соединений.                                        |
| Handled     | counter | ngx_http_stub_status_module  | Суммарное число обработанных соединений.                                               |
| Requests    | counter | ngx_http_stub_status_module  | Суммарное число клиентских запросов.                                                   |
| Reading     | gauge   | ngx_http_stub_status_module  | Текущее число соединений, в которых nginx в настоящий момент читает заголовок запроса. |
| Writing     | gauge   | ngx_http_stub_status_module  | Текущее число соединений, в которых nginx в настоящий момент отвечает клиенту.         |
| Waiting     | gauge   | ngx_http_stub_status_module  | Текущее число бездействующих клиентских соединений в ожидании запроса.                 |

### Ограничения
Перед использованием коллектора nginxStubStatus, необходимо предварительно включить ngx_http_stub_status_module в nginx.  
Подробней в документации nginx: [http://nginx.org/ru/docs/http/ngx_http_stub_status_module.html](http://nginx.org/ru/docs/http/ngx_http_stub_status_module.html)

### Конфигурация
```json
{
  "collectors": {
    "nginxStubStatus": {
      "enabled": true,
      "endpoint": "http://127.0.0.1/basic_status",
      "timeout": 5
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор
* **"endpoint"** - URL к веб-странице ngx_http_stub_status_module
* **"timeout"** - таймаут на выполнение запроса к ngx_http_stub_status_module в секундах

## cmd
### Описание
### Конфигурация
