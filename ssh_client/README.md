
The goal of this repo is to allow the user to create configuration/ 
validation files that will go out to the appropriate device and issue
the appropriate commands. Device authentication distinction is done by
using if-else logic to make sure the appropiate device are logged into.
The file naming convention is used to determined how the device is configured.

The main file main.go uses a hostfile to determine which devices to log into.

hostfile.txt
- hostname:port#
- hostname:port#

Main.go calls unique cmd files for each device using the following naming 
standard to determine the commands to apply to each device.

File name format:  "hostname" + ":ssh Port" + ".cfg"

- hostname:port.cfg
- hostname:port.cfg

- eg. CORE_R1:20.cfg

--------------------

#Package files

```
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_client $ ls -l
total 3840
-rw-r--r-- 1 pi pi       0 Dec  5 14:18 README.md
-rw-r--r-- 1 pi pi      91 Dec  5 14:20 go.mod
-rw-r--r-- 1 pi pi     832 Dec  5 14:20 go.sum
-rw-r--r-- 1 pi pi      65 Dec  5 14:30 host_file.txt
-rw-r--r-- 1 pi pi    1948 Dec  5 14:42 main.go
-rw-r--r-- 1 pi pi     345 Dec  5 14:50 sandbox-iosxe-latest-1.cisco.com:22.cfg
-rw-r--r-- 1 pi pi     367 Dec  5 14:52 sbx-nxos-mgmt.cisco.com:8181.cfg
-rwxr-xr-x 1 pi pi 3903487 Dec  5 14:45 ssh_client
```

#host_file.txt
```
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_client $ cat host_file.txt 
sbx-nxos-mgmt.cisco.com:8181
sandbox-iosxe-latest-1.cisco.com:22
```

# cmds for host1
```
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_client $ cat sandbox-iosxe-latest-1.cisco.com\:22.cfg 
show run int loopback76
!! REMOVE INTERFACE
config t
 interface loopback76
 no ip address
 no interface loopback 76
 exit
!! VALIDATE
show ip int brief

!! CONFIGURE INTERFACE
config t
interface Loopback76
 description scripted with Go
 ip address 1.1.1.76 255.255.255.255
exit
exit
!
!! VALIDATE
show run int loopback 76
show ip int brief
exit
```

# cmds for host2

```
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_client $ cat sbx-nxos-mgmt.cisco.com\:8181.cfg 
!
! REMOVE INTERFACE
!
config t
interface eth1/24
 no ip address
 no description
 switchport 
 shut
 exit
exit
!
! VALIDATE
!
show ip int brief
!
! CONFIG INTERFACE
!
config t
interface Eth1/24
  no switchport
  description scripted interface - Go Rules...
  ip address 24.24.24.1/30
  no shut
  exit
exit
!
! VALIDATE
!
show run int eth1/24
show ip int brief
exit

````
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_client $ 


-------------------
# Results of the program

```
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_client $ ./ssh_client  
[sbx-nxos-mgmt.cisco.com:8181 sandbox-iosxe-latest-1.cisco.com:22]


This is the config file named:sbx-nxos-mgmt.cisco.com:8181.cfg




stty: standard input: Inappropriate ioctl for device


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

!Command: show running-config interface Ethernet1/24
!Running configuration last done at: Sat Dec  5 07:45:03 2020
!Time: Sat Dec  5 07:45:04 2020

version 9.3(3) Bios:version  

interface Ethernet1/24
  description scripted interface - Go Rules...
  no switchport
  ip address 24.24.24.1/30
  no shutdown


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
Eth1/24              24.24.24.1      protocol-down/link-down/admin-up   


This is the config file named:sandbox-iosxe-latest-1.cisco.com:22.cfg





Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.



csr1000v-1#show run int loopback76
Building configuration...

Current configuration : 95 bytes
!
interface Loopback76
 description scripted with Go
 ip address 1.1.1.76 255.255.255.255
end

csr1000v-1#!! REMOVE INTERFACE
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)# interface loopback76
csr1000v-1(config-if)# no ip address
csr1000v-1(config-if)# no interface loopback 76
csr1000v-1(config)# exit
csr1000v-1#!! VALIDATE
csr1000v-1#show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       10.255.255.1    YES other  up                    up      
GigabitEthernet3       unassigned      YES NVRAM  down                  down    
Loopback0              unassigned      YES unset  up                    up      
Loopback1              2.2.2.2         YES manual up                    up      
Loopback2              unassigned      YES unset  up                    up      
Loopback10             192.168.10.10   YES manual up                    up      
Loopback1337           192.168.100.1   YES other  up                    up      
Loopback1500           unassigned      YES unset  up                    up      
Port-channel1          unassigned      YES unset  down                  down    
Port-channel1.101      unassigned      YES unset  down                  down    
VirtualPortGroup0      172.31.0.1      YES manual up                    up      
csr1000v-1#
csr1000v-1#!! CONFIGURE INTERFACE
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)#interface Loopback76
csr1000v-1(config-if)# description scripted with Go
csr1000v-1(config-if)# ip address 1.1.1.76 255.255.255.255
csr1000v-1(config-if)#exit
csr1000v-1(config)#exit
csr1000v-1#!
csr1000v-1#!! VALIDATE
csr1000v-1#show run int loopback 76
Building configuration...

Current configuration : 95 bytes
!
interface Loopback76
 description scripted with Go
 ip address 1.1.1.76 255.255.255.255
end

csr1000v-1#show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       10.255.255.1    YES other  up                    up      
GigabitEthernet3       unassigned      YES NVRAM  down                  down    
Loopback0              unassigned      YES unset  up                    up      
Loopback1              2.2.2.2         YES manual up                    up      
Loopback2              unassigned      YES unset  up                    up      
Loopback10             192.168.10.10   YES manual up                    up      
Loopback76             1.1.1.76        YES manual up                    up      
Loopback1337           192.168.100.1   YES other  up                    up      
Loopback1500           unassigned      YES unset  up                    up      
Port-channel1          unassigned      YES unset  down                  down    
Port-channel1.101      unassigned      YES unset  down                  down    
VirtualPortGroup0      172.31.0.1      YES manual up                    up      
csr1000v-1#exit
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_client $ 

```



