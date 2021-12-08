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

| Имя метрики              | Тип    | Источник данных | Описание                                                                                                 |
|--------------------------|--------|-----------------|----------------------------------------------------------------------------------------------------------|
| Last                     | gauge  | /proc/loadavg   | Значение Load Average усредненное за 1 минуту.                                                           |
| Last5m                   | gauge  | /proc/loadavg   | Значение Load Average усредненное за 5 минут.                                                            |
| Last15m                  | gauge  | /proc/loadavg   | Значение Load Average усредненное за 15 минут.                                                           |
| KernelSchedulingEntities | gauge  | /proc/loadavg   | Количество объектов планирования ядра(процессы, потоки), которые в настоящее время существуют в системе. |

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

| Имя метрики      | Тип    | Источник данных         | Описание                                      |
|------------------|--------|-------------------------|-----------------------------------------------|
| Count            | gauge  | /proc/\[pid\]/status    | Количество запущенных процессов в системе.    |
| Limit            | gauge  | /sys/kernel/pid_max     | Максимальное допустимое количество процессов. |
| InStateRunning   | gauge  | /proc/\[pid\]/status    | Количество процессов в состоянии running.     | 
| InStateIdle      | gauge  | /proc/\[pid\]/status    | Количество процессов в состоянии idle.        |
| InStateSleeping  | gauge  | /proc/\[pid\]/status    | Количество процессов в состоянии sleeping.    |
| InStateDiskSleep | gauge  | /proc/\[pid\]/status    | Количество процессов в состоянии disk sleep.  |
| InStateStopped   | gauge  | /proc/\[pid\]/status    | Количество процессов в состоянии stopped.     |
| InStateZombie    | gauge  | /proc/\[pid\]/status    | Количество процессов в состоянии zombie.      |
| InStateDead      | gauge  | /proc/\[pid\]/status    | Количество процессов в состоянии dead.        |
| Threads          | gauge  | /proc/\[pid\]/status    | Количество запущенных потоков.                |
| ThreadsLimit     | gauge  | /sys/kernel/threads-max | Максимальное допустимое количество потоков.   | 
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

| Имя метрики        | Тип    | Источник данных         | Описание                                            |
|--------------------|--------|-------------------------|-----------------------------------------------------|
| Processes          | gauge  | /proc/\[pid\]/stat      | Количество процессов.                               |
| Threads            | gauge  | /proc/\[pid\]/stat      | Количество потоков.                                 |
| System             | gauge  | /proc/\[pid\]/stat      | Использование CPU(system).                          |
| User               | gauge  | /proc/\[pid\]/stat      | Использование CPU(user).                            |
| Guest              | gauge  | /proc/\[pid\]/stat      | Использование CPU(guest).                           |
| Total              | gauge  | /proc/\[pid\]/stat      | Использование CPU(system+user+guest).               |
| ResidentMemorySize | gauge  | /proc/\[pid\]/stat      | Размер памяти (в байтах) выделенной процессу в RAM. |
| VirtualMemorySize  | gauge  | /proc/\[pid\]/stat      | Размер виртуальной памяти в байтах.                 |

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
### Конфигурация

## cpuInfo
### Описание
### Конфигурация

## memoryInfo
### Описание
### Конфигурация

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

## netSNMP
### Описание
### Конфигурация

## netSNMP6
### Описание
### Конфигурация

## netARP
### Описание
### Конфигурация

## entropy
### Описание
### Конфигурация

## nginxStubStatus
### Описание
### Конфигурация

## cmd
### Описание
### Конфигурация
