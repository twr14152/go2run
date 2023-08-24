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
        fmt.Printf("\n\n#######################\n\n")
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

x.x.x.x:22

```

# Running code with show commands

Remember when you give the host device to add the port you're connecting on.

```
go run main.go 
Connecting to ios-xe devices:
Number of hosts: 1
Hostname: iosxe:22

cmds: show ip route, show ip int brief


Welcome to the DevNet Sandbox for CSR1000v and IOS XE

The following programmability features are already enabled:
  - NETCONF
  - RESTCONF

Thanks for stopping by.



csr1000v-1#term len 0
csr1000v-1#show ip route
Codes: L - local, C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route, H - NHRP, l - LISP
       a - application route
       + - replicated route, % - next hop override, p - overrides from PfR

Gateway of last resort is 10.10.20.254 to network 0.0.0.0

S*    0.0.0.0/0 [1/0] via 10.10.20.254, GigabitEthernet1
      10.0.0.0/8 is variably subnetted, 7 subnets, 2 masks
C        10.0.0.0/24 is directly connected, GigabitEthernet2
L        10.0.0.100/32 is directly connected, GigabitEthernet2
C        10.2.1.1/32 is directly connected, Loopback987
C        10.10.20.0/24 is directly connected, GigabitEthernet1
L        10.10.20.48/32 is directly connected, GigabitEthernet1
C        10.255.255.0/24 is directly connected, GigabitEthernet2
L        10.255.255.100/32 is directly connected, GigabitEthernet2
      88.0.0.0/8 is variably subnetted, 2 subnets, 2 masks
C        88.88.88.0/24 is directly connected, Loopback88
L        88.88.88.88/32 is directly connected, Loopback88
      99.0.0.0/8 is variably subnetted, 2 subnets, 2 masks
C        99.99.99.0/24 is directly connected, Loopback99
L        99.99.99.100/32 is directly connected, Loopback99
      192.168.100.0/32 is subnetted, 1 subnets
C        192.168.100.1 is directly connected, Loopback203
csr1000v-1# show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       10.0.0.100      YES other  up                    up      
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
GigabitEthernet3.1     unassigned      YES unset  deleted               down    
GigabitEthernet3.2     unassigned      YES unset  deleted               down    
Loopback10             unassigned      YES unset  up                    up      
Loopback31             unassigned      YES unset  up                    up      
Loopback88             88.88.88.88     YES other  up                    up      
Loopback99             99.99.99.100    YES other  up                    up      
Loopback203            192.168.100.1   YES other  up                    up      
Loopback987            10.2.1.1        YES other  up                    up      
csr1000v-1#
csr1000v-1#exit


#######################

Connecting to nxos device:
Number of hosts: 1    
Hostname: nxos:22

cmds: show ip route, show ip int brief


stty: standard input: Inappropriate ioctl for device
gl_set_term_size: NULL arguments(s).
IP Route Table for VRF "default"
'*' denotes best ucast next-hop
'**' denotes best mcast next-hop
'[x/y]' denotes [preference/metric]
'%<string>' in via output denotes VRF <string>

10.98.98.0/24, ubest/mbest: 1/0, attached
    *via 10.98.98.1, Lo98, [0/0], 3d17h, direct
10.98.98.1/32, ubest/mbest: 1/0, attached
    *via 10.98.98.1, Lo98, [0/0], 3d17h, local
10.99.99.0/24, ubest/mbest: 1/0, attached
    *via 10.99.99.1, Lo99, [0/0], 3d17h, direct
10.99.99.1/32, ubest/mbest: 1/0, attached
    *via 10.99.99.1, Lo99, [0/0], 3d17h, local
10.131.22.0/24, ubest/mbest: 1/0, attached
    *via 10.131.22.192, Vlan406, [0/0], 1d00h, direct
10.131.22.192/32, ubest/mbest: 1/0, attached
    *via 10.131.22.192, Vlan406, [0/0], 1d00h, local
110.110.110.111/32, ubest/mbest: 2/0, attached
    *via 110.110.110.111, Lo110, [0/0], 1d09h, local
    *via 110.110.110.111, Lo110, [0/0], 1d09h, direct
172.16.10.0/24, ubest/mbest: 1/0, attached
    *via 172.16.10.4, Lo1, [0/0], 13:56:54, direct
172.16.10.4/32, ubest/mbest: 1/0, attached
    *via 172.16.10.4, Lo1, [0/0], 13:56:54, local


IP Interface Status for VRF "default"(1)
Interface            IP Address      Interface Status
Vlan100              172.16.100.1    protocol-down/link-down/admin-down 
Vlan406              10.131.22.192   protocol-up/link-up/admin-up       
Lo1                  172.16.10.4     protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       
Lo110                110.110.110.111 protocol-up/link-up/admin-up       
Eth1/5               172.16.1.1      protocol-down/link-down/admin-down 
pi@RaspPi4:
```


# Running code with configuration commands + validation commands.

Using the same test runcli.go file we will add loopback75 to fastxe csr and loopback76 to nxos device.

```
go run main.go 
Connecting to ios-xe devices:
Number of hosts: 1
Hostname: iosxe:22

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
GigabitEthernet2       10.0.0.100      YES other  up                    up      
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
GigabitEthernet3.1     unassigned      YES unset  deleted               down    
GigabitEthernet3.2     unassigned      YES unset  deleted               down    
Loopback10             unassigned      YES unset  up                    up      
Loopback31             unassigned      YES unset  up                    up      
Loopback88             88.88.88.88     YES other  up                    up      
Loopback99             99.99.99.100    YES other  up                    up      
Loopback110            110.110.110.110 YES manual up                    up      
Loopback203            192.168.100.1   YES other  up                    up      
Loopback987            10.2.1.1        YES other  up                    up      
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


#######################

Connecting to nxos device:
Number of hosts: 1
Hostname: nxos:22

cmds: config t, interface loopback 110, description runcli_script_nxos, ip address 110.110.110.111/32, exit, exit, show ip int brief, show run interface loopback 110, exit


stty: standard input: Inappropriate ioctl for device
gl_set_term_size: NULL arguments(s).

IP Interface Status for VRF "default"(1)
Interface            IP Address      Interface Status
Vlan100              172.16.100.1    protocol-down/link-down/admin-down 
Vlan406              10.131.22.192   protocol-up/link-up/admin-up       
Lo1                  172.16.10.4     protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       
Lo110                110.110.110.111 protocol-up/link-up/admin-up       
Eth1/5               172.16.1.1      protocol-down/link-down/admin-down 

!Command: show running-config interface loopback110
!Running configuration last done at: Tue Aug 22 09:36:58 2023
!Time: Tue Aug 22 09:36:58 2023

version 9.3(3) Bios:version  

interface loopback110
  description runcli_script_nxos
  ip address 110.110.110.111/32


```

