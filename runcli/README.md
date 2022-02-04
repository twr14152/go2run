# runcli

The idea behind this code is that you could quickly gather info from network devices on the fly as well as make minor changes if so desired.
When you import this package into your code it will ask you how many devices you want to connect to, as well as what commands you want to run.
The code will work with show commands as well as configuration commands. Good for gathering info and for minor changes especially in a lab environment.


In this example we have one ios-xe and one nx-os device. The login parameters for the ios-xe and the nxos is different. The app will then prompt you to enter the commands you want. In your code all you will need to do is import "runcli" and add your login credentials to runcli.RunCli() for each group. The app will then prompt you for the commands to run.

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


The one thing you will need to do when you get prompted is to provide the hostname or ip address with the port you're connecting on.

For example:
```
host1:22

host2:8181

```

# Running code with show commands

Remember when you give the host device to add the port your connecting on.

```
 go run main.go 
Connecting to ios-xe devices: 
Number of hosts: 1
Hostname: 131.226.217.143:22

cmds: show ip int brief  


Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.



csr1000v-1#term len 0
csr1000v-1#show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       unassigned      YES NVRAM  administratively down down    
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
Loopback105            192.168.1.1     YES manual up                    up      
Loopback106            192.168.1.2     YES manual up                    up      
csr1000v-1#
csr1000v-1#exit
Connecting to nxos devices: 
Number of hosts: 1
Hostname: 64.103.37.14:8181

cmds: show ip int brief


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
 
```


# Running code with configuration commands + validation commands.

Using the same testruncli.go file we will add loopback75 to fastxe csr and loopback76 to nxos device.

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

