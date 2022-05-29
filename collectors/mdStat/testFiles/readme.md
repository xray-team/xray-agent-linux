https://www.digitalocean.com/community/tutorials/how-to-manage-raid-arrays-with-mdadm-on-ubuntu-16-04

# 1 
disks:

    root@mdstat:~# lsblk 
    NAME   MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
    sda      8:0    0   16G  0 disk 
    └─sda1   8:1    0   16G  0 part /
    sdb      8:16   0   20G  0 disk 
    ├─sdb1   8:17   0    7G  0 part 
    └─sdb2   8:18   0   13G  0 part 
    sdc      8:32   0   20G  0 disk 
    ├─sdc1   8:33   0    7G  0 part 
    └─sdc2   8:34   0   13G  0 part 
    sdd      8:48   0   20G  0 disk 
    ├─sdd1   8:49   0    7G  0 part 
    └─sdd2   8:50   0   13G  0 part 
    
## 1.1 RAID1-create-missing
action:
    
    mdadm --create --verbose /dev/md0 --level=mirror --raid-devices=2 /dev/sdb1 missing
    mdadm --create --verbose /dev/md1 --level=mirror --raid-devices=2 /dev/sdb2 missing
    
mdstat:
<pre>
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md1 : active raid1 sdb2[0]
      13621248 blocks super 1.2 [2/1] [U_]
      
md0 : active raid1 sdb1[0]
      7334912 blocks super 1.2 [2/1] [U_]
      
unused devices: <none>
</pre>

disks:

    root@mdstat:~# lsblk
    NAME    MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda       8:0    0   16G  0 disk  
    └─sda1    8:1    0   16G  0 part  /
    sdb       8:16   0   20G  0 disk  
    ├─sdb1    8:17   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdb2    8:18   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1 
    sdc       8:32   0   20G  0 disk  
    ├─sdc1    8:33   0    7G  0 part  
    └─sdc2    8:34   0   13G  0 part  
    sdd       8:48   0   20G  0 disk    
    ├─sdd1    8:49   0    7G  0 part  
    └─sdd2    8:50   0   13G  0 part 

## 1.2 RAID1-resync-DELAYED
action:
    
	mdadm --manage /dev/md0 --add /dev/sdc1
	mdadm --manage /dev/md1 --add /dev/sdc2
    
mdstat:
<pre>
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md1 : active raid1 sdc2[2] sdb2[0]
      13621248 blocks super 1.2 [2/1] [U_]
      	resync=DELAYED
      
md0 : active raid1 sdc1[2] sdb1[0]
      7334912 blocks super 1.2 [2/1] [U_]
      [====>................]  recovery = 21.4% (1576832/7334912) finish=0.4min speed=225261K/sec
      
unused devices: <none>
</pre>

disks:
    
    NAME    MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda       8:0    0   16G  0 disk  
    └─sda1    8:1    0   16G  0 part  /
    sdb       8:16   0   20G  0 disk  
    ├─sdb1    8:17   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdb2    8:18   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1 
    sdc       8:32   0   20G  0 disk  
    ├─sdc1    8:33   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdc2    8:34   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1 
    sdd       8:48   0   20G  0 disk  
    ├─sdd1    8:49   0    7G  0 part  
    └─sdd2    8:50   0   13G  0 part
 
## 1.3 RAID1-ok

action: 

    wait for sync

mdstat:
<pre>
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md1 : active raid1 sdc2[2] sdb2[0]
      13621248 blocks super 1.2 [2/2] [UU]
      
md0 : active raid1 sdc1[2] sdb1[0]
      7334912 blocks super 1.2 [2/2] [UU]
      
unused devices: <none>
</pre>
    
## 1.4 RAID1-disk-fail
action:

    mdadm --manage /dev/md0 --fail /dev/sdb1
    mdadm --manage /dev/md1 --fail /dev/sdc2
    
mdstat:
<pre>
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md1 : active raid1 sdc2[2](F) sdb2[0]
      13621248 blocks super 1.2 [2/1] [U_]
      
md0 : active raid1 sdc1[2] sdb1[0](F)
      7334912 blocks super 1.2 [2/1] [_U]
      
unused devices: <none>
</pre>

disks:

    NAME    MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda       8:0    0   16G  0 disk  
    └─sda1    8:1    0   16G  0 part  /
    sdb       8:16   0   20G  0 disk  
    ├─sdb1    8:17   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdb2    8:18   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1 
    sdc       8:32   0   20G  0 disk  
    ├─sdc1    8:33   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdc2    8:34   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1 
    sdd       8:48   0   20G  0 disk  
    ├─sdd1    8:49   0    7G  0 part  
    └─sdd2    8:50   0   13G  0 part
    
## 1.5 RAID1-with-spare
action:

    mdadm --manage /dev/md0 --remove /dev/sdb1
    mdadm --manage /dev/md0 --add /dev/sdb1
    mdadm --manage /dev/md0 --add /dev/sdd1
    
    mdadm --manage /dev/md1 --remove /dev/sdc2
    mdadm --manage /dev/md1 --add /dev/sdc2
    mdadm --manage /dev/md1 --add /dev/sdd2
    
mdstat:
<pre>
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md1 : active raid1 sdd2[3](S) sdc2[2] sdb2[0]
      13621248 blocks super 1.2 [2/2] [UU]
      
md0 : active raid1 sdd1[4](S) sdb1[3] sdc1[2]
      7334912 blocks super 1.2 [2/2] [UU]
      
unused devices: <none>
</pre>

disks:

    NAME    MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda       8:0    0   16G  0 disk  
    └─sda1    8:1    0   16G  0 part  /
    sdb       8:16   0   20G  0 disk  
    ├─sdb1    8:17   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdb2    8:18   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1 
    sdc       8:32   0   20G  0 disk  
    ├─sdc1    8:33   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdc2    8:34   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1 
    sdd       8:48   0   20G  0 disk  
    ├─sdd1    8:49   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdd2    8:50   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1 
      
## 1.6 RAID1-disk-fail-with-spare

action:

    mdadm --manage /dev/md0 --fail /dev/sdb1
    mdadm --manage /dev/md1 --fail /dev/sdd2
    
mdstat:
<pre>
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md1 : active raid1 sdd2[3](F) sdc2[2] sdb2[0]
      13621248 blocks super 1.2 [2/2] [UU]
      
md0 : active raid1 sdd1[4] sdb1[3](F) sdc1[2]
      7334912 blocks super 1.2 [2/1] [_U]
      [==>..................]  recovery = 11.0% (810240/7334912) finish=0.5min speed=202560K/sec
      
unused devices: <none>
</pre>
disks:

    NAME    MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda       8:0    0   16G  0 disk  
    └─sda1    8:1    0   16G  0 part  /
    sdb       8:16   0   20G  0 disk  
    ├─sdb1    8:17   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdb2    8:18   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1 
    sdc       8:32   0   20G  0 disk  
    ├─sdc1    8:33   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdc2    8:34   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1 
    sdd       8:48   0   20G  0 disk  
    ├─sdd1    8:49   0    7G  0 part  
    │ └─md0   9:0    0    7G  0 raid1 
    └─sdd2    8:50   0   13G  0 part  
      └─md1   9:1    0   13G  0 raid1
      
## 1.7 RAID5 and RAID0
action:

    mdadm --stop /dev/md0 
    mdadm --create --verbose /dev/md0 --level=raid5 --raid-devices=3 /dev/sd{b,c,d}1
    
    mdadm --stop /dev/md1
    mdadm --create --verbose /dev/md1 --level=raid0 --raid-devices=2 /dev/sd{b,c}2
    
mdstat:
<pre>
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md1 : active raid0 sdc2[1] sdb2[0]
      27242496 blocks super 1.2 512k chunks
      
md0 : active raid5 sdd1[3] sdc1[1] sdb1[0]
      14669824 blocks super 1.2 level 5, 512k chunk, algorithm 2 [3/3] [UUU]
      
unused devices: <none>
</pre>
disks:

    NAME    MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda       8:0    0   16G  0 disk  
    └─sda1    8:1    0   16G  0 part  /
    sdb       8:16   0   20G  0 disk  
    ├─sdb1    8:17   0    7G  0 part  
    │ └─md0   9:0    0   14G  0 raid5 
    └─sdb2    8:18   0   13G  0 part  
      └─md1   9:1    0   26G  0 raid0 
    sdc       8:32   0   20G  0 disk  
    ├─sdc1    8:33   0    7G  0 part  
    │ └─md0   9:0    0   14G  0 raid5 
    └─sdc2    8:34   0   13G  0 part  
      └─md1   9:1    0   26G  0 raid0 
    sdd       8:48   0   20G  0 disk  
    ├─sdd1    8:49   0    7G  0 part  
    │ └─md0   9:0    0   14G  0 raid5 
    └─sdd2    8:50   0   13G  0 part 
    

# 2 
disks:

    root@mdstat:~# lsblk 
    NAME   MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
    sda      8:0    0   16G  0 disk 
    └─sda1   8:1    0   16G  0 part /
    sdb      8:16   0   10G  0 disk 
    └─sdb1   8:17   0   10G  0 part 
    sdc      8:32   0   10G  0 disk 
    └─sdc1   8:33   0   10G  0 part 
    sdd      8:48   0   10G  0 disk 
    └─sdd1   8:49   0   10G  0 part 
    sde      8:64   0   10G  0 disk 
    └─sde1   8:65   0   10G  0 part
    
## 2.1 RAID6
action:

    mdadm --create --verbose /dev/md0 --level=6 --raid-devices=4 /dev/sdb1 /dev/sdc1 /dev/sdd1 /dev/sde1
    
mdstat:
<pre>
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md0 : active raid6 sde1[3] sdd1[2] sdc1[1] sdb1[0]
      20951040 blocks super 1.2 level 6, 512k chunk, algorithm 2 [4/4] [UUUU]
      
unused devices: <none>
</pre>

disks:

    NAME    MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
    sda       8:0    0   16G  0 disk  
    └─sda1    8:1    0   16G  0 part  /
    sdb       8:16   0   10G  0 disk  
    └─sdb1    8:17   0   10G  0 part  
      └─md0   9:0    0   20G  0 raid6 
    sdc       8:32   0   10G  0 disk  
    └─sdc1    8:33   0   10G  0 part  
      └─md0   9:0    0   20G  0 raid6 
    sdd       8:48   0   10G  0 disk  
    └─sdd1    8:49   0   10G  0 part  
      └─md0   9:0    0   20G  0 raid6 
    sde       8:64   0   10G  0 disk  
    └─sde1    8:65   0   10G  0 part  
      └─md0   9:0    0   20G  0 raid6
      
## 2.2 RAID10
action:
    
     mdadm --stop /dev/md0
     mdadm --create --verbose /dev/md0 --level=10 --raid-devices=4 /dev/sdb1 /dev/sdc1 /dev/sdd1 /dev/sde1
     
mdstat:
<pre>
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md0 : active raid10 sde1[3] sdd1[2] sdc1[1] sdb1[0]
      20951040 blocks super 1.2 512K chunks 2 near-copies [4/4] [UUUU]
      
unused devices: <none>
</pre>

disks:

    NAME    MAJ:MIN RM  SIZE RO TYPE   MOUNTPOINT
    sda       8:0    0   16G  0 disk   
    └─sda1    8:1    0   16G  0 part   /
    sdb       8:16   0   10G  0 disk   
    └─sdb1    8:17   0   10G  0 part   
      └─md0   9:0    0   20G  0 raid10 
    sdc       8:32   0   10G  0 disk   
    └─sdc1    8:33   0   10G  0 part   
      └─md0   9:0    0   20G  0 raid10 
    sdd       8:48   0   10G  0 disk   
    └─sdd1    8:49   0   10G  0 part   
      └─md0   9:0    0   20G  0 raid10 
    sde       8:64   0   10G  0 disk   
    └─sde1    8:65   0   10G  0 part   
      └─md0   9:0    0   20G  0 raid10