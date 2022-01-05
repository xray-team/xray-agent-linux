[English version](../en/collectors.md)

---
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
- [wireless](collectors.md#wireless)
- [entropy](collectors.md#entropy)
- [nginxStubStatus](collectors.md#nginxstubstatus)
- [cmd](collectors.md#cmd)

## uptime
### Описание
Коллектор uptime собирает следующие метрики из /proc/uptime:

| Имя метрики | Тип     | Источник данных | Описание                                                                                                                               |
|-------------|---------|-----------------|----------------------------------------------------------------------------------------------------------------------------------------|
| Uptime      | counter | /proc/uptime    | Время работы системы с момента последней загрузки в секундах.                                                                          |
| Idle        | counter | /proc/uptime    | Сумма времени простоя всех cpu (ядер) в секундах. Для систем с количеством ядер больше 1 данное значение может быть больше чем Uptime. |

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
Коллектор psStat собирает статистику использования CPU и памяти для указанных в конфигурации приложений. Сбор метрик происходит аналогично тому, как это делает утилита top (top -p $pid).  
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

Метрики по каждому CPU и сводные метрики по всем CPU (Total):

| Имя метрики | Тип     | Источник данных | Описание                                                                                                                                                                                                                                                                                                                                    |
|-------------|---------|-----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| User        | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованного в режиме User.                                                                                                                                                                                                                                                            |
| Nice        | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованного в режиме Nice.                                                                                                                                                                                                                                                            |
| System      | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованного в режиме System.                                                                                                                                                                                                                                                          |
| Idle        | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованного в режиме Idle.                                                                                                                                                                                                                                                            |
| IOwait      | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованного в режиме IOwait.                                                                                                                                                                                                                                                          |
| IRQ         | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованного на обслуживание hardware прерываний.                                                                                                                                                                                                                                      |
| SoftIRQ     | counter | /proc/stat      | Количество процессорного времени, использованного на обслуживание software прерываний.                                                                                                                                                                                                                                                      |
| Steal       | counter | /proc/stat      | Количество процессорного времени в миллисекундах, которое было необходимо гостевой виртуальной машине, но не было предоставлено хостом. Большое количество времени Steal указывает на конкуренцию за CPU, например при переподписке CPU. Большое количество времени Steal является индикатором снижения производительности гостевых систем. |
| Guest       | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованного гостевыми операционными системами.                                                                                                                                                                                                                                        |
| GuestNice   | counter | /proc/stat      | Количество процессорного времени в миллисекундах, использованного гостевыми операционными системами (в режиме Nice).                                                                                                                                                                                                                        |

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
Коллектор memoryInfo собирает следующие метрики использования памяти:

| Имя метрики | Тип   | Источник данных | Описание                                                                     |
|-------------|-------|-----------------|------------------------------------------------------------------------------|
| Total       | gauge | /proc/meminfo   | Общий объем ОЗУ в КБ.                                                        |
| Free        | gauge | /proc/meminfo   | Объем свободной ОЗУ в КБ.                                                    |
| Available   | gauge | /proc/meminfo   | Объем доступной ОЗУ в КБ.                                                    |
| Used        | gauge | /proc/meminfo   | Объем использованной ОЗУ в КБ.                                               |
| Buffers     | gauge | /proc/meminfo   | Объем использованной ОЗУ для дискового буфера в КБ.                          |
| Cached      | gauge | /proc/meminfo   | Объем использованной ОЗУ для дискового кеша в КБ.                            |
| SwapTotal   | gauge | /proc/meminfo   | Общий объем пространства подкачки в КБ.                                      |
| SwapFree    | gauge | /proc/meminfo   | Объем пространства подкачки, которое в настоящее время не используется в КБ. |

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
Коллектор diskStat, для каждого диска, собирает следующие метрики:

| Имя метрики                        | Тип       | Источник данных | Описание                                                                                                                                                                                          | Ограничения                                |
|------------------------------------|-----------|-----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------|
| ReadsCompletedSuccessfully         | counter   | /proc/diskstats | Общее количество успешно завершенных операций чтения.                                                                                                                                             |                                            |
| ReadsMerged                        | counter   | /proc/diskstats | Общее количество объединенных операций чтения. Смежные операции могут быть объединены в одну для увеличения эффективности. Такие операции считаются и выполняются как одна операция ввода/вывода. |                                            |
| SectorsRead                        | counter   | /proc/diskstats | Общее количество успешно прочитанных секторов.                                                                                                                                                    |                                            |
| TimeSpentReading                   | counter   | /proc/diskstats | Общее количество времени (в миллисекундах), затраченного на выполнение операций чтения.                                                                                                           |                                            |
| WritesCompleted                    | counter   | /proc/diskstats | Общее количество успешно завершенных операций записи.                                                                                                                                             |                                            |
| WritesMerged                       | counter   | /proc/diskstats | Общее количество объединенных операций записи. Смежные операции могут быть объединены в одну для увеличения эффективности. Такие операции считаются и выполняются как одна операция ввода/вывода. |                                            |
| SectorsWritten                     | counter   | /proc/diskstats | Общее количество успешно записанных секторов.                                                                                                                                                     |                                            |
| TimeSpentWriting                   | counter   | /proc/diskstats | Общее количество времени (в миллисекундах), затраченного на выполнение операций записи.                                                                                                           |                                            |
| IOsCurrentlyInProgress             | gauge     | /proc/diskstats | Общее количество операций ввода-вывода, находящихся в настоящее время в обработке.                                                                                                                |                                            |
| TimeSpentDoingIOs                  | counter   | /proc/diskstats | Общее количество времени (в миллисекундах), когда значение IOsCurrentlyInProgress не было равно нулю.                                                                                             |                                            |
| WeightedTimeSpentDoingIOs          | counter   | /proc/diskstats | Общее количество времени (в миллисекундах), затраченное на выполнение операций ввода/вывода.                                                                                                      | Kernel 4.18+                               |
| DiscardsCompletedSuccessfully      | counter   | /proc/diskstats | Общее количество успешно завершенных операций discard. Операции discard выполняются при выполнении TRIM.                                                                                          | Kernel 4.18+                               |
| DiscardsMerged                     | counter   | /proc/diskstats | Общее количество объединенных операций discard.                                                                                                                                                   | Kernel 4.18+                               |
| SectorsDiscarded                   | counter   | /proc/diskstats | Общее количество успешно очищенных секторов.                                                                                                                                                      | Kernel 4.18+                               |
| TimeSpentDiscarding                | counter   | /proc/diskstats | Общее количество времени (в миллисекундах), затраченного на выполнение операций discard.                                                                                                          | Kernel 4.18+                               |
| FlushRequestsCompletedSuccessfully | counter   | /proc/diskstats | Общее количество успешно завершенных операций flush.                                                                                                                                              | Kernel 5.5+, не отслеживается для разделов |
| TimeSpentFlushing                  | counter   | /proc/diskstats | Общее количество времени (в миллисекундах), затраченного на выполнение операций flush.                                                                                                            | Kernel 5.5+, не отслеживается для разделов |

Подробнее: 
* [https://www.kernel.org/doc/Documentation/ABI/testing/procfs-diskstats](https://www.kernel.org/doc/Documentation/ABI/testing/procfs-diskstats)
* [https://www.kernel.org/doc/Documentation/admin-guide/iostats.rst](https://www.kernel.org/doc/Documentation/admin-guide/iostats.rst).
### Конфигурация
```json
{
  "collectors": {
    "diskStat": {
      "enabled": true,
      "diskTypes": [
        8,
        9
      ],
      "excludePartitions": false,
      "excludeByName": []
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор
* **"diskTypes"** - список типов дисков (major numbers), для которых нужно собирать метрики.  
Подробней о типах дисков: [https://www.kernel.org/doc/html/latest/admin-guide/devices.html](https://www.kernel.org/doc/html/latest/admin-guide/devices.html).  
Расшифровка часто-используемых типов:
  * 8 - SCSI(sata)-диски. Например: sda, sda1, sdb, sdb1 и т.д.
  * 9 - Metadisk (RAID). Например: md0, md1 и т.д.
* **"excludePartitions"**
  * **true** - исключает из статистики разделы дисков
  * **false** - не исключает из статистики разделы дисков
* **"excludeByName"** - список имен дисков, которые нужно исключить из статистики. Например: "excludeByName": \["sda2", "sda3"\]

## diskSpace
### Описание
Коллектор diskSpace для каждого раздела собирает следующую статистику использования дискового пространства:

| Имя метрики      | Тип   | Источник данных | Описание                                                                            |
|------------------|-------|-----------------|-------------------------------------------------------------------------------------|
| BytesTotal       | gauge | Statfs          | Размер раздела в байтах.                                                            |
| BytesAvailable   | gauge | Statfs          | Количество доступных байт.                                                          |
| BytesFree        | gauge | Statfs          | Количество свободных байт.                                                          |
| BytesFreePercent | gauge | Statfs          | Количество свободного пространства в процентах.                                     |
| BytesUsed        | gauge | Statfs          | Количество использованных байт. Вычисляется по формуле BytesTotal - BytesAvailable. |
| InodesTotal      | gauge | Statfs          | Общее количество Inodes.                                                            |
| InodesFree       | gauge | Statfs          | Количество свободных Inodes файловой системы.                                       |
| InodesUsed       | gauge | Statfs          | Количество Inodes, вычисленных по формуле BytesTotal - BytesAvailable.              |
### Конфигурация
```json
{
  "collectors": {
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
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор
* **"fsTypes"** - список типов файловых систем, для которых нужно собирать метрики.  
Подробнее о файловых системах: [https://man7.org/linux/man-pages/man5/filesystems.5.html](https://man7.org/linux/man-pages/man5/filesystems.5.html)

## mdStat
### Описание
Коллектор mdStat собирает информацию о состоянии linux software RAID. Для каждого md-устройства собираются следующие метрики:

| Имя метрики          | Тип     | Источник данных                      | Описание                                                                                                                                                                  |
|----------------------|---------|--------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Level                | gauge   | /sys/block/\[md*\]/md/level          | Тип RAID-масcива: raid1 - 1, raid6 - 6, raid10 - 10, и т.д.                                                                                                               |
| NumDisks             | gauge   | /sys/block/\[md*\]/md/raid_disks     | Количество дисков в масcиве.                                                                                                                                              |
| ArrayState           | gauge   | /sys/block/\[md*\]/md/array_state    | Состояние RAID-масcива: "clear" - 1, "inactive" - 2, "suspended" - 3, "readonly" - 4, "read-auto" - 5, "clean" - 6, "active" - 7, "write-pending" - 8, "active-idle" - 9. |
| ArraySize            | gauge   | /sys/block/\[md*\]/md/               | Эффективный размер массива в килобайтах.                                                                                                                                  |
| SyncAction           | gauge   | /sys/block/\[md*\]/md/sync_action    | Текущее состояние синхронизации RAID-массива: "resync" - 1, "recover" - 2, "idle" - 3, "check" - 4, "repair" - 5.                                                         |
| NumDegraded          | gauge   | /sys/block/\[md*\]/md/degraded       | Количество деградирующих дисков.                                                                                                                                          |
| MismatchCnt          | counter | /sys/block/\[md*\]/md/mismatch_cnt   | Количество секторов, которые были перезаписаны для проверки.                                                                                                              |
| SyncCompletedSectors | counter | /sys/block/\[md*\]/md/sync_completed | Количество секторов, обработка которых была завершена независимо от текущего sync_action.                                                                                 |
| NumSectors           | gauge   | /sys/block/\[md*\]/md/               | Общее количество секторов, которые нужно обработать.                                                                                                                      |
| SyncSpeed            | gauge   | /sys/block/\[md*\]/md/sync_speed     | Текущая фактическая скорость в КБ/сек текущего sync_action. Среднее значение за последние 30 секунд.                                                                      |

Для каждого диска в RAID-массиве собираются следующие метрики:

| Имя метрики | Тип    | Источник данных                     | Описание                                                                                                                                                   |
|-------------|--------|-------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Slot        | gauge  | /sys/block/\[mdN\]/md/\[dev\]slot   | Позиция диска в RAID-масcиве.                                                                                                                              |
| State       | gauge  | /sys/block/\[mdN\]/md/\[dev\]state  | Состояние диска: "faulty" - 1, "in_sync" - 2, "writemostly" - 3, "blocked" - 4, "spare" - 5, "write_error" - 6, "want_replacement" - 7, "replacement" - 8. |
| Errors      | gauge  | /sys/block/\[mdN\]/md/\[dev\]errors | Приблизительное количество ошибок чтения, которые были обнаружены на этом устройстве, но не привели к удалению устройства из raid массива.                 |

Подробнее:
* [https://www.kernel.org/doc/html/latest/admin-guide/md.html](https://www.kernel.org/doc/html/latest/admin-guide/md.html)
### Конфигурация
```json
{
  "collectors": {
    "mdStat": {
      "enabled": true
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор

## netDev
### Описание
Коллектор netDev собирает статистику для каждого сетевого интерфейса из /proc/net/dev.  
Описание всех счетчиков можно найти здесь: [https://www.kernel.org/doc/html/latest/networking/statistics.html#struct-rtnl-link-stats64](https://www.kernel.org/doc/html/latest/networking/statistics.html#struct-rtnl-link-stats64)  
### Конфигурация
```json
{
  "collectors": {
    "netDev": {
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
  * **true** - включить коллектор
  * **false** - отключить коллектор
* **"excludeLoopbacks"**
  * **true** - исключает петлевые (loopback) интерфейсы из статистики
  * **false** - не исключает петлевые (loopback) интерфейсы из статистики
* **"excludeWireless"**
  * **true** - исключает беспроводные интерфейсы из статистики
  * **false** - не исключает беспроводные интерфейсы из статистики
* **"excludeBridges"**
  * **true** - исключает Bridge-интерфейсы из статистики
  * **false** - не исключает Bridge-интерфейсы из статистики
* **"excludeVirtual"**
  * **true** - исключает виртуальные интерфейсы из статистики
  * **false** - не исключает виртуальные интерфейсы из статистики
* **"excludeByName"** - список имен интерфейсов, которые нужно исключить из статистики
* **"excludeByOperState"** - список состояний интерфейсов, которые нужно исключить из статистики исходя из состояния интерфейса. Возможные состояния: "unknown", "notpresent", "down", "lowerlayerdown", "testing", "dormant", "up" (расшифровка здесь [https://tools.ietf.org/html/rfc2863#section-3.1.14](https://tools.ietf.org/html/rfc2863#section-3.1.14))

## netDevStatus
### Описание
Коллектор netDevStatus собирает информацию о состоянии сетевых интерфейсов.  
Для каждого интерфейса доступны следующие метрики:

| Имя метрики | Тип     | Источник данных                          | Описание                                                                                                                                                                                                                                                                                 |
|-------------|---------|------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| OperState   | gauge   | /sys/class/net/\[iface\]/operstate       | Состояние интерфейса. Может принимать следующие значения: 1 - "up", 2 - "lowerlayerdown", 3 - "dormant", 4 - "down", 5 - "unknown", 6 - "testing", 7 - "notpresent". Подробнее: [https://tools.ietf.org/html/rfc2863#section-3.1.14](https://tools.ietf.org/html/rfc2863#section-3.1.14) |
| LinkFlaps   | counter | /sys/class/net/\[iface\]/carrier_changes | Количество флапов интерфейса.                                                                                                                                                                                                                                                            |
| Speed       | gauge   | /sys/class/net/\[iface\]/speed           | Отображает последнее,  или текущее значение скорости интерфейса в Мбит/с.                                                                                                                                                                                                                |
| Duplex      | gauge   | /sys/class/net/\[iface\]/duplex          | Отображает последнее,  или текущее значение Duplex'a. Возможные значения: 1 - "full", 2 - "half", 3 - "unknown". Подробнее https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-class-net.                                                                                         |
| MTU         | gauge   | /sys/class/net/\[iface\]/mtu             | Отображает установленное значение MTU для интерфейса.                                                                                                                                                                                                                                    |
### Конфигурация
```json
{
  "collectors": {
    "netDev": {
      "enabled": true,
      "excludeWireless": true,
      "excludeByName": []
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор
* **"excludeWireless"**
  * **true** - исключает беспроводные интерфейсы из статистики
  * **false** - не исключает беспроводные интерфейсы из статистики
* **"excludeByName"** - список имен интерфейсов, которые нужно исключить из статистики

## netStat
### Описание
Коллектор netStat собирает статистику сетевой активности ОС из /proc/net/netstat.  
Количество и состав метрик может отличатся в зависимости от версии и настроек ядра Linux.  
Подробное описание метрик можно найти здесь: [https://www.kernel.org/doc/html/latest/networking/snmp_counter.html](https://www.kernel.org/doc/html/latest/networking/snmp_counter.html)
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
Коллектор netSNMP6 собирает все доступные метрики из /proc/net/snmp6.  
Данный коллектор собирает аналогичные коллектору netSNMP метрики для IPv6.
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

## wireless
### Описание
Коллектор wireless с помощью утилиты iwconfig собирает статистику для беспроводных интерфейсов.

| Имя метрики        | Тип     | Источник данных | Описание                                                                                                                                                  |
|--------------------|---------|-----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| Frequency          | gauge   | iwconfig        | Канал.                                                                                                                                                    |
| BitRate            | gauge   | iwconfig        | Скорость подключения.                                                                                                                                     |
| TxPower            | gauge   | iwconfig        | TX power в dBm.                                                                                                                                           |
| LinkQuality        | gauge   | iwconfig        | Качество подключения. Абстрактное значение, которое может базироваться на качестве принимаемого сигнала, частоте ошибок и/или других аппаратных метриках. |
| LinkQualityLimit   | gauge   | iwconfig        | Максимальное значение для LinkQuality.                                                                                                                    |
| SignalLevel        | gauge   | iwconfig        | Уровень (мощность) принимаемого сигнала.                                                                                                                  |
| RxInvalidNwid      | counter | iwconfig        | Количество пакетов, полученных с другим NWID или ESSID.                                                                                                   |
| RxInvalidCrypt     | counter | iwconfig        | Количество пакетов, которые оборудование не смогло расшифровать.                                                                                          |
| RxInvalidFrag      | counter | iwconfig        | Количество пакетов, для которых оборудование не могло должным образом повторно собрать фрагменты канального уровня (скорее всего, один отсутствовал).     |
| TxExcessiveRetries | counter | iwconfig        | Количество пакетов, которые не удалось доставить.                                                                                                         |
| InvalidMisc        | counter | iwconfig        | Количество пакетов, которые были потеряны по другим причинам.                                                                                             |
| MissedBeacon       | counter | iwconfig        | Количество пропущенных периодических wireless beacon от точки доступа.                                                                                    |

Подробнее: man iwconfig

### Конфигурация
```json
{
  "collectors": {
    "wireless": {
      "enabled": true,
      "excludeByName": [],
      "excludeByOperState": ["down"]
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор
* **"excludeByName"** - список имен интерфейсов, которые нужно исключить из статистики
* **"excludeByOperState"** - список состояний wireless-интерфейсов, для которых не нужно собирать метрики. Возможные состояния: "unknown", "notpresent", "down", "lowerlayerdown", "testing", "dormant", "up" (расшифровка здесь [https://tools.ietf.org/html/rfc2863#section-3.1.14](https://tools.ietf.org/html/rfc2863#section-3.1.14))

## entropy
### Описание
Коллектор entropy собирает следующие метрики:

| Имя метрики | Тип   | Источник данных                       | Описание                                                                                               |
|-------------|-------|---------------------------------------|--------------------------------------------------------------------------------------------------------|
| Available   | gauge | /proc/sys/kernel/random/entropy_avail | Количество доступной энтропии (количество бит, доступных в настоящее время для чтения из /dev/random). |

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
| Active      | gauge   | ngx_http_stub_status_module  | Текущее число активных клиентских соединений, включая Waiting-соединения.              |
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
Коллектор cmd позволяет выполнять произвольные shell-команды, или скрипты и записывать результат их выполнения в виде отдельных метрик. Таким образом возможно быстро и просто расширять функциональность xray-agent собственными метриками.
### Конфигурация
```json
{
  "collectors": {
    "cmd": {
      "enabled": true,
      "timeout": 1,
      "metrics": [
        {
          "names": ["metricName1", "metricName2", "metricNameN"],
          "delimiter": " ",
          "pipeline": [
            ["command1", "argument1ForCommand1", "argument2ForCommand1", "argumentNForCommand1"],
            ["command2", "argument1ForCommand2", "argument2ForCommand2", "argumentNForCommand2"],
            ["commandN", "argument1ForCommandN", "argument2ForCommandN", "argumentNForCommandN"]
          ],
          "attributes": [
            {
              "name": "AttributeName",
              "value": "AttributeValue"
            }
          ]
        }
      ]
    }
  }
}
```
* **"enabled"**
  * **true** - включить коллектор
  * **false** - отключить коллектор
* **"timeout"** - таймаут выполнения каждого pipeline'а в секундах. Минимальное значение 1, максимальное - 120
* **"metrics"** - конфигурация пользовательских команд
  * **"names"** - список названий метрик. Количество элементов данного списка должно равняться количеству значений, которые возвращает команда указанная в "pipeline". Что-бы игнорировать (пропустить) ненужные значения - укажите вместо имени "-".
  * **"delimiter"** - строка-разделитель
  * **"pipeline"** - список команд, которые нужно выполнить, что-бы получить метрики. В свою очередь каждая команда представляется в виде массива произвольной длины, в котором первый элемент - это исполняемый файл, а последующие элементы - аргументы.  
  Для наглядности несколько примеров:
    * Entropy  
      *shell*:
      ```shell
      cat /proc/sys/kernel/random/entropy_avail
      ```
      *эквивалентный xray cmd pipeline*:
      ```json
      "pipeline": [
        ["cat", "/proc/sys/kernel/random/entropy_avail"]
      ]
      ```
    * OOM killer  
      *shell*:
      ```shell
      cat /proc/vmstat | grep oom_kill
      ```
      *эквивалентный xray cmd pipeline*:
      ```json
      "pipeline": [
        ["cat", "/proc/vmstat"],
        ["grep", "oom_kill"]
      ]
      ```
    * ARP Limits  
      *shell*:
      ```shell
      sysctl net.ipv4.neigh.default.gc_thresh1 net.ipv4.neigh.default.gc_thresh2 net.ipv4.neigh.default.gc_thresh3 | awk '{print $3}' ORS=' '
      ```
      *эквивалентный xray cmd pipeline*:
      ```json
      "pipeline": [
        ["sysctl", "net.ipv4.neigh.default.gc_thresh1", "net.ipv4.neigh.default.gc_thresh2", "net.ipv4.neigh.default.gc_thresh3"],
        ["awk", "{print $3}", "ORS= "]
      ]
      ```
* **"attributes"** - дополнительные атрибуты, которые нужно добавить к метрикам

**Пример 1.**
Например, нужно сохранять в виде метрики с именем EntropyAvail вывод следующей команды:
```shel
cat /proc/sys/kernel/random/entropy_avail
```
В таком случае в конфигурационный файл xray-agent нужно будет добавить следующие строки:
```json
{
  "collectors": {
    "cmd": {
      "enabled": true,
      "metrics": [
        {
          "names": ["EntropyAvail"],
          "pipeline": [
            ["cat", "/proc/sys/kernel/random/entropy_avail"]
          ],
          "delimiter": " "
        }
      ]
    }
  }
}
```
**Пример 2.** Например, нужно сохранять в виде 3 метрик вывод следующей команды:
```shel
sysctl net.ipv4.neigh.default.gc_thresh1 net.ipv4.neigh.default.gc_thresh2 net.ipv4.neigh.default.gc_thresh3 | awk '{print $3}' ORS=' '
```
В таком случае в конфигурационный файл xray-agent нужно будет добавить следующие строки:
```json
{
  "collectors": {
    "cmd": {
      "metrics": [
        {
          "names": ["GcThresh1", "GcThresh2", "GcThresh3"],
          "pipeline": [
            ["sysctl", "net.ipv4.neigh.default.gc_thresh1", "net.ipv4.neigh.default.gc_thresh2", "net.ipv4.neigh.default.gc_thresh3"],
            ["awk", "{print $3}", "ORS= "]
          ],
          "delimiter": " ",
          "attributes": [
            {
              "name": "Set",
              "value": "IPv4NeighbourTable"
            }
          ]
        }
      ]
    }
  }
}
```