# runscript
The goal of this package is to allow the user to use host and cmd files to change and validate device configuration.
You will simply need to import runscript into your main.go file or what ever you decide to call it. Then call the runscript.Connect() to connect to your devices. All you will need is the username password and the name of the hostfile you create with those devices with common login parameters. If you have multiple login parameters create multiple hostfiles grouping those with common parameters and make another function call.

# To Use:

sample code:
main.go
```
package main

import (
    "fmt"
    "github.com/twr14152/go2run/runscript"
)

func main() {
    fmt.Println("Connecting to Group1 hosts")
    runscript.Connect("USER1", "PASS1", "group1.txt")
    fmt.Println("Connecting to Group2 hosts")
    runscript.Connect("USER2", "PASS2", "group2.txt")
}


```
# To install the files from this repo
```
$go mod init <folder/filename> // if go.mod exists already this command is unneccessary
$go mod tidy //This will actually download the dependency files listed in the main.go import statements
```


You can call the hostfile what ever you want. It is neccesary to include the connecting port in the file.

hostfile.txt
- hostname:port
- ip_address:port

# This is key

runscript.go calls unique cmd files for each device using the following naming standard to determine the commands to apply to each device.

Commands files:

File name format:
- "file_"+"hostname"+":ssh_Port"+".cfg" 
- or -
- "file_"+"ip address"+":ssh_Port"+".cfg"

Eg.
- file_core_r1:22.cfg
- file_10.0.1.100:8181.cfg
 
By using the "file_" in front of the name you're able to use IP address as well has hostnames.

# In summary to use this package you will need to do the following:
```
- Import runscript into main.go
- Then use runscript.Connect("username", "password", "hostsfile.txt") to your devices 
- Create commands files for each device USING THE FORMAT PROVIDED.
- Create hostfile grouping those hosts that share common login parameters

```

Example provided:
--------------------

#Package files

```
$ ls -l
total 4672
-rw-r--r-- 1 runner runner     267 Feb  4 11:27 file_sandbox-iosxe-latest-1.cisco.com:22.cfg
-rw-r--r-- 1 runner runner     264 Feb  4 11:25 file_sandbox-nxos-1.cisco.com:22.cfg
-rw-r--r-- 1 runner runner     102 Feb  2 13:07 go.mod
-rw-r--r-- 1 runner runner    1280 Feb  2 13:07 go.sum
-rw-r--r-- 1 runner runner      36 Feb  4 11:22 group1.txt
-rw-r--r-- 1 runner runner      28 Feb  2 21:40 group2.txt
-rw-r--r-- 1 runner runner     310 Feb  4 11:33 main.go
$
```

# host_files

add as many hosts as you need per group. 
```
$ vi hostfile1.txt
hostname1:<port>
hostname2:<port>
...
```
demo hosts
```
$ cat group1.txt 
sandbox-iosxe-latest-1.cisco.com:22
$ cat group2.txt 
sandbox-nxos-1.cisco.com:22 
 
```
# cmds for host1
```
$ cat file_sandbox-iosxe-latest-1.cisco.com\:22.cfg 
term len 0

show interface description | inc Lo

config t
int loopback 70
description go2run_test_script_loop70
exit
exit


show interface description  | inc Lo
config t
int loopback70
no description
no int loopback 70
exit
show interface description | inc Lo

exit

$
```

# cmds for host2
```
$cat file_sandbox-nxos-1.cisco.com\:22.cfg 
show interface description | inc Lo

config 

interface loopback 72
description go2run_testscript
exit
exit


show run interface loopback 72
show interface description | inc Lo

config
interface loopback 72
no description
no interface loopback 72
exit


show run intface loopback 72
show interface description | inc Lo

exit
```

-------------------
# Results of the program

```
$ go run main.go
Connecting to Group1 hosts
[sandbox-iosxe-latest-1.cisco.com:22]


This is the config file named:file_sandbox-iosxe-latest-1.cisco.com:22.cfg





Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.


csr1000v-1#term len 0
csr1000v-1#
csr1000v-1#show interface description | inc Lo
Lo100                          up             up       
Lo123                          up             up       Configured by NETCONF
Lo1234                         up             up       Added with RESTCONF
Lo1235                         up             up       Added with RESTCONF
Lo1236                         up             up       Added with RESTCONF
Lo1237                         up             up       Added with goues
Lo1238                         up             up       Added with RESTCONF1
csr1000v-1#
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)#int loopback 70
csr1000v-1(config-if)#description go2run_test_script_loop70
csr1000v-1(config-if)#exit
csr1000v-1(config)#exit
csr1000v-1#
csr1000v-1#
csr1000v-1#show interface description  | inc Lo
Lo70                           up             up       go2run_test_script_loop70
Lo100                          up             up       
Lo123                          up             up       Configured by NETCONF
Lo1234                         up             up       Added with RESTCONF
Lo1235                         up             up       Added with RESTCONF
Lo1236                         up             up       Added with RESTCONF
Lo1237                         up             up       Added with goues
Lo1238                         up             up       Added with RESTCONF1
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)#int loopback70
csr1000v-1(config-if)#no description
csr1000v-1(config-if)#no int loopback 70
csr1000v-1(config)#exit
csr1000v-1#show interface description | inc Lo
Lo100                          up             up       
Lo123                          up             up       Configured by NETCONF
Lo1234                         up             up       Added with RESTCONF
Lo1235                         up             up       Added with RESTCONF
Lo1236                         up             up       Added with RESTCONF
Lo1237                         up             up       Added with goues
Lo1238                         up             up       Added with RESTCONF1
csr1000v-1#
csr1000v-1#exit
Connecting to Group2 hosts
[sandbox-nxos-1.cisco.com:22]


This is the config file named:file_sandbox-nxos-1.cisco.com:22.cfg





this is my exec banner
that contains a multiline
string

stty: standard input: Inappropriate ioctl for device
Lo0                      TEST NXAPI-REST
Lo1                      --
Lo2                      --
Lo3                      --
Lo4                      --
Lo5                      --
Lo6                      --
Lo7                      --
Lo10                     test
Lo12                     Sample Network Route Injection
Lo13                     Sample Network Route Injection
Lo14                     Sample Network Route Injection
Lo20                     My Learning Lab Loopback Home
Lo21                     MSN Learning Lab Loopback
Lo24                     --
Lo50                     test123
Lo55                     Full intf config via NETCONF
Lo99                     Full intf config via NETCONF
Lo100                    loopback int for routing
Lo101                    --
Lo123                    test
Lo201                    --
Lo300                    --
Lo400                    IPv6 ADDRESS
Lo432                    Configured using OpenConfig Model
Lo555                    Configured by Salt-Nornir using NETCONF
Lo666                    Loopback 666
Lo667                    Loopback 667
Lo737                    Full intf config via NETCONF
Lo778                    Full intf config via NETCONF
Lo980                    Full intf config via NETCONF
Lo987                    Configured using OpenConfig Model
Lo999                    Full intf config via NETCONF
Lo1000                   --

!Command: show running-config interface loopback72
!Running configuration last done at: Fri Feb  4 00:02:25 2022
!Time: Fri Feb  4 00:02:25 2022

version 9.3(3) Bios:version  

interface loopback72
  description go2run_testscript

Lo0                      TEST NXAPI-REST
Lo1                      --
Lo2                      --
Lo3                      --
Lo4                      --
Lo5                      --
Lo6                      --
Lo7                      --
Lo10                     test
Lo12                     Sample Network Route Injection
Lo13                     Sample Network Route Injection
Lo14                     Sample Network Route Injection
Lo20                     My Learning Lab Loopback Home
Lo21                     MSN Learning Lab Loopback
Lo24                     --
Lo50                     test123
Lo55                     Full intf config via NETCONF
Lo72                     go2run_testscript
Lo99                     Full intf config via NETCONF
Lo100                    loopback int for routing
Lo101                    --
Lo123                    test
Lo201                    --
Lo300                    --
Lo400                    IPv6 ADDRESS
Lo432                    Configured using OpenConfig Model
Lo555                    Configured by Salt-Nornir using NETCONF
Lo666                    Loopback 666
Lo667                    Loopback 667
Lo737                    Full intf config via NETCONF
Lo778                    Full intf config via NETCONF
Lo980                    Full intf config via NETCONF
Lo987                    Configured using OpenConfig Model
Lo999                    Full intf config via NETCONF
Lo1000                   --
Syntax error while parsing 'show run intface loopback 72'

Lo0                      TEST NXAPI-REST
Lo1                      --
Lo2                      --
Lo3                      --
Lo4                      --
Lo5                      --
Lo6                      --
Lo7                      --
Lo10                     test
Lo12                     Sample Network Route Injection
Lo13                     Sample Network Route Injection
Lo14                     Sample Network Route Injection
Lo20                     My Learning Lab Loopback Home
Lo21                     MSN Learning Lab Loopback
Lo24                     --
Lo50                     test123
Lo55                     Full intf config via NETCONF
Lo99                     Full intf config via NETCONF
Lo100                    loopback int for routing
Lo101                    --
Lo123                    test
Lo201                    --
Lo300                    --
Lo400                    IPv6 ADDRESS
Lo432                    Configured using OpenConfig Model
Lo555                    Configured by Salt-Nornir using NETCONF
Lo666                    Loopback 666
Lo667                    Loopback 667
Lo737                    Full intf config via NETCONF
Lo778                    Full intf config via NETCONF
Lo980                    Full intf config via NETCONF
Lo987                    Configured using OpenConfig Model
Lo999                    Full intf config via NETCONF
Lo1000                   --
$
$
```
