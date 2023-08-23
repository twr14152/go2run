# runcli

The idea behind this code is that you could quickly gather info from network devices on the fly as well as make minor changes if so desired.
When you import this package into your code it will ask you how many devices you want to connect to, as well as what commands you want to run.
The code will work with show commands as well as configuration commands. Good for gathering info and for minor changes especially in a lab environment.


In this example we have one ios-xe and one nx-os device. The login parameters for the ios-xe and the nxos is different. The app will then prompt you to enter the commands you want. In your code all you will need to do is import "github.com/twr14152/go2run/runcli" and add your login credentials to runcli.RunCli() for each group. The app will then prompt you for the commands to run.

Sample code:

Created file main.go and added the following:
```
package main

import (
	"fmt"
	"github.com/twr14152/go2run/runcli"
)

func main() {
	fmt.Println("Connecting to ios-xe devices:")
	runcli.RunCli("username1", "password1")
	fmt.Println("\n\n\nConnecting to nxos device:")
	runcli.RunCli("username2", "password2")
}
```
# To install the files used in this repo:
```
$go mod init <folder/filename> // if go.mod exists already this command is unneccessary
$go mod tidy  // This will actually pull the dependency files from the repo based off import statements in main.go

```


The one thing you will need to do when you get prompted is to provide the hostname or ip address with the PORT you are connecting on.

For example:
```
host1:22

host2:8181


```

# Running code with show commands

Remember when you give the host device to add the port your connecting on.

```
 go run main.go
Connecting to ios-xe devices:
Number of hosts: 1
Hostname: sandbox-iosxe-recomm-1.cisco.com:22

cmds: show ip int brief, show version | inc Ver


Welcome to the DevNet Sandbox for CSR1000v and IOS XE

The following programmability features are already enabled:
  - NETCONF
  - RESTCONF

Thanks for stopping by.


csr1000v-1#term len 0
csr1000v-1#show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet1.10    unassigned      YES other  deleted               down    
GigabitEthernet2       172.16.1.2      YES manual up                    up      
GigabitEthernet3       200.200.210.199 YES manual up                    up      
Loopback20             10.1.1.1        YES other  up                    up      
Loopback100            192.168.100.100 YES manual up                    up      
Loopback222            192.0.2.222     YES manual up                    up      
Loopback223            192.0.2.223     YES manual up                    up      
Loopback409            172.16.30.122   YES other  up                    up      
Loopback12345          9.9.9.9         YES manual up                    up      
Loopback1234567        unassigned      YES unset  up                    up      
Loopback2110999        178.18.90.1     YES other  up                    up      
nve1                   unassigned      YES unset  down                  down    
csr1000v-1# show version | inc Ver
Cisco IOS XE Software, Version 16.09.03
Cisco IOS Software [Fuji], Virtual XE Software (X86_64_LINUX_IOSD-UNIVERSALK9-M), Version 16.9.3, RELEASE SOFTWARE (fc2)
licensed under the GNU General Public License ("GPL") Version 2.0.  The
software code licensed under GPL Version 2.0 is free software that comes
GPL code under the terms of GPL Version 2.0.  For more details, see the
csr1000v-1#
csr1000v-1#exit
Connecting to nxos device:
Number of hosts: 1
Hostname: sbx-nxos-mgmt.cisco.com:22

cmds: show ip int brief, show int status | inc connected


stty: standard input: Inappropriate ioctl for device
gl_set_term_size: NULL arguments(s).

IP Interface Status for VRF "default"(1)
Interface            IP Address      Interface Status
Vlan100              172.16.100.1    protocol-down/link-down/admin-down 
Vlan101              172.16.101.1    protocol-down/link-down/admin-down 
Vlan102              172.16.102.1    protocol-down/link-down/admin-down 
Vlan103              172.16.103.1    protocol-down/link-down/admin-down 
Vlan104              172.16.104.1    protocol-down/link-down/admin-down 
Vlan105              172.16.105.1    protocol-down/link-down/admin-down 
Lo1                  172.16.0.1      protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       
Eth1/5               172.16.1.1      protocol-down/link-down/admin-down 
mgmt0         DO NOT TOUCH CONFI connected routed    full    1000    --         
Eth1/2        --                 connected trunk     full    1000    10g        
Eth1/3        --                 connected 1         full    1000    10g        
Eth1/4        --                 connected 1         full    1000    10g        
Po11          --                 connected trunk     full    1000    --         
Lo0           --                 connected routed    auto    auto    --         
Lo1           --                 connected routed    auto    auto    --         
Lo30          My Learning Lab Lo connected routed    auto    auto    --         
Lo98          Configured using O connected routed    auto    auto    --         
Lo99          Full intf config v connected routed    auto    auto    --         
Lo100         Full intf config v connected routed    auto    auto    --         
Lo122         --                 connected routed    auto    auto    --         
Lo200         MASSIVO            connected routed    auto    auto    --         
Vlan402       --                 connected routed    auto    auto    --
$ 
```


# Running code with configuration commands + validation commands.

Using the same test runcli.go file we will add loopback75 to fastxe csr and loopback76 to nxos device.

```
 go run main.go 
Connecting to ios-xe devices: 
Number of hosts: 1
Hostname: 131.226.217.143:22

cmds: config t, interface loopback 110, description runcli_script_iosxe, ip address 110.110.110.110 255.255.255.255, exit, exit, show ip interface brief, show run interface loopback 110


Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.



csr1000v-1#term len 0
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)# interface loopback 110
csr1000v-1(config-if)# description runcli_script_iosxe
csr1000v-1(config-if)# ip address 110.110.110.110 255.255.255.255
csr1000v-1(config-if)# exit
csr1000v-1(config)# exit
csr1000v-1# show ip interface brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       unassigned      YES NVRAM  administratively down down    
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
Loopback105            192.168.1.1     YES manual up                    up      
Loopback106            192.168.1.2     YES manual up                    up      
Loopback110            110.110.110.110 YES manual up                    up      
csr1000v-1# show run interface loopback 110
Building configuration...

Current configuration : 106 bytes
!
interface Loopback110
 description runcli_script_iosxe
 ip address 110.110.110.110 255.255.255.255
end

csr1000v-1#
csr1000v-1#exit
Connecting to nxos devices: 
Number of hosts: 1
Hostname: 64.103.37.14:8181

cmds: config t, interface loopback 110, description runcli_script_nxos, ip address 110.110.110.111/32, exit, exit, show ip int brief, show run interface loopback 110, exit


stty: standard input: Inappropriate ioctl for device
gl_set_term_size: NULL arguments(s).

IP Interface Status for VRF "default"(1)
Interface            IP Address      Interface Status
Vlan100              172.16.100.1    protocol-down/link-down/admin-down 
Vlan101              172.16.101.1    protocol-down/link-down/admin-down 
Vlan102              172.16.102.1    protocol-down/link-down/admin-down 
Vlan103              172.16.103.1    protocol-down/link-down/admin-down 
Vlan104              172.16.104.1    protocol-down/link-down/admin-down 
Vlan105              172.16.105.1    protocol-down/link-down/admin-down 
Lo1                  172.16.0.1      protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       
Lo110                110.110.110.111 protocol-up/link-up/admin-up       
Eth1/5               172.16.1.1      protocol-down/link-down/admin-down 

!Command: show running-config interface loopback110
!Running configuration last done at: Sun Dec 20 04:16:21 2020
!Time: Sun Dec 20 04:16:22 2020

version 9.3(3) Bios:version  

interface loopback110
  description runcli_script_nxos
  ip address 110.110.110.111/32

 

```

