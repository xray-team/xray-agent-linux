# CMD
## Description
The CMD collector allows you to execute shell commands or scripts and record the results as separate metrics.
Thus, it is possible to quickly and easily extend the functionality of the xray-agent with your own metrics.
## Configuration
```json
{
  "collectors": {
    "CMD": {
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
    * **true** - enable collector
    * **false** - disable collector
* **"timeout"** - timeout for pipeline in seconds. Min: 1, max: 120
* **"metrics"** - command configuration
    * **"names"** - list of metric names. The number of elements in this list must be equal to the number of values returned by the command specified in the "pipeline". To ignore (skip) unnecessary values, specify "-" instead of name.
    * **"delimiter"** - delimiter string
    * **"pipeline"** - list of commands  
      Examples:
        * Entropy  
          *shell*:
          ```shell
          cat /proc/sys/kernel/random/entropy_avail
          ```
          *equivalent xray cmd pipeline*:
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
          *equivalent xray cmd pipeline*:
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
          *equivalent xray cmd pipeline*:
          ```json
          "pipeline": [
            ["sysctl", "net.ipv4.neigh.default.gc_thresh1", "net.ipv4.neigh.default.gc_thresh2", "net.ipv4.neigh.default.gc_thresh3"],
            ["awk", "{print $3}", "ORS= "]
          ]
          ```
* **"attributes"** - additional attributes

**Example 1.**
For example, if you want to save the output of the following command as a metric named EntropyAvail:
```shel
cat /proc/sys/kernel/random/entropy_avail
```
add to xray-agent config:
```json
{
  "collectors": {
    "CMD": {
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
**Example 2.** For example, if you need to save the output of the following command as 3 metrics:
```shel
sysctl net.ipv4.neigh.default.gc_thresh1 net.ipv4.neigh.default.gc_thresh2 net.ipv4.neigh.default.gc_thresh3 | awk '{print $3}' ORS=' '
```
add to xray-agent config:
```json
{
  "collectors": {
    "CMD": {
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