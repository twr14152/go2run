# runcli
The idea behind this code is that you could quickly gather info on the fly as well as make minor changes if so desired.
When you run the code it will ask you how many devices you want to connect to, as well as what commands you want to run.
The code will work with show commands as well as configuration commands. Good for minor changes especially in a lab environment.

The example below is using Cisco Devnet hosts but I've updated my /etc/hosts file to make it easier to test.
The code below was tested against ios-xe and nxos hosts. The one thing you will need to do is when you get prompted to provide
the hostname of your device you need to add the port they are connecting with. For example:
```
host1:22
host2:8181

```
Once you've downloaded package you need to make sure your GOROOT knows where to find it.

On my raspberry PI I needed to copy the file over to /home/pi/go/src/ it may be something different for you.
```
cp -r runcli/ /home/pi/go/src/
```

In your code all you will need to add is login credentials to the file you create.

Sample code:

Create file testruncli.go add the following:
```
package main

import (
	"fmt"
	"runcli"
)

func main() {
	fmt.Println("Connecting to ios-xe devices:")
	runcli.RunCli("username1", "password1")
	fmt.Println("\n\n\nConnecting to nxos device:")
	runcli.RunCli("username2", "password2")
}
```
# Running code with show commands

The devices used in the code below each had different usernames and passwords. The code worked fine. Just needed to tweak the code using if/else logic
with login creds.

```
pi@raspberrypi:~/Code_folder/go_folder/go2run $ go run testruncli.go 
Connecting to ios-xe devices:
Number of hosts: 2
Hostname: fastxe:22

cmds: show ip int brief

Hostname: slowxe:8181

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
Loopback0              100.64.0.1      YES other  up                    up      
Loopback300            unassigned      YES unset  up                    up      
Loopback444            unassigned      YES unset  up                    up      
VirtualPortGroup0      192.168.35.1    YES manual up                    up      
csr1000v-1#
csr1000v-1#exit

Welcome to the DevNet Sandbox for CSR1000v and IOS XE

The following programmability features are already enabled:
  - NETCONF
  - RESTCONF

Thanks for stopping by.



csr1000v#term len 0
csr1000v#show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       unassigned      YES NVRAM  administratively down down    
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
csr1000v#
csr1000v#exit



Connecting to nxos device:
Number of hosts: 1
Hostname: nxos:8181

cmds: show ip int brief


stty: standard input: Inappropriate ioctl for device
gl_set_term_size: NULL arguments(s).

IP Interface Status for VRF "default"(1)
Interface            IP Address      Interface Status
Vlan100              172.16.100.1    protocol-up/link-up/admin-up       
Vlan101              172.16.101.1    protocol-down/link-down/admin-down 
Vlan102              172.16.102.1    protocol-down/link-down/admin-down 
Vlan103              172.16.103.1    protocol-down/link-down/admin-down 
Vlan104              172.16.104.1    protocol-down/link-down/admin-down 
Vlan105              172.16.105.1    protocol-down/link-down/admin-down 
Lo1                  172.16.0.1      protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       
Eth1/5               172.16.1.1      protocol-down/link-down/admin-down 
pi@raspberrypi:~/Code_folder/go_folder/go2run $ $ 
```


# Running code with configuration commands + validation commands.

In this example we will add loopback75 to fastxe csr and loopback76 to nxos device.

```
pi@raspberrypi:~/Code_folder/go_folder/go2run $ go run testruncli.go 
Connecting to ios-xe devices:
Number of hosts: 1
Hostname: fastxe:22

cmds: config t, interface loopback75, ip address 75.0.0.1 255.255.255.255, description go script test, exit, exit, show ip int brief, show run int loopback75


Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.



csr1000v-1#term len 0
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)# interface loopback75
csr1000v-1(config-if)# ip address 75.0.0.1 255.255.255.255
csr1000v-1(config-if)# description go script test
csr1000v-1(config-if)# exit
csr1000v-1(config)# exit
csr1000v-1# show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       unassigned      YES NVRAM  administratively down down    
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
Loopback0              100.64.0.1      YES other  up                    up      
Loopback75             75.0.0.1        YES manual up                    up      
Loopback300            unassigned      YES unset  up                    up      
Loopback444            unassigned      YES unset  up                    up      
VirtualPortGroup0      192.168.35.1    YES manual up                    up      
csr1000v-1# show run int loopback75
Building configuration...

Current configuration : 93 bytes
!
interface Loopback75
 description go script test
 ip address 75.0.0.1 255.255.255.255
end

csr1000v-1#
csr1000v-1#exit



Connecting to nxos device:
Number of hosts: 1        
Hostname: nxos:8181

cmds: config t, interface loopback76, ip address 76.0.0.1/32, description go script test, exit, exit, show ip int brief, show run int loopback76

stty: standard input: Inappropriate ioctl for device

gl_set_term_size: NULL arguments(s).

IP Interface Status for VRF "default"(1)
Interface            IP Address      Interface Status
Vlan100              172.16.100.1    protocol-up/link-up/admin-up       
Vlan101              172.16.101.1    protocol-down/link-down/admin-down 
Vlan102              172.16.102.1    protocol-down/link-down/admin-down 
Vlan103              172.16.103.1    protocol-down/link-down/admin-down 
Vlan104              172.16.104.1    protocol-down/link-down/admin-down 
Vlan105              172.16.105.1    protocol-down/link-down/admin-down 
Lo1                  172.16.0.1      protocol-up/link-up/admin-up       
Lo76                 76.0.0.1        protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       
Eth1/5               172.16.1.1      protocol-down/link-down/admin-down 

!Command: show running-config interface loopback76
!Running configuration last done at: Wed Dec  9 01:08:25 2020
!Time: Wed Dec  9 01:08:25 2020

version 9.3(3) Bios:version  

interface loopback76
  description go script test
  ip address 76.0.0.1/32

pi@raspberrypi:~/Code_folder/go_folder/go2run $ 
 

```

