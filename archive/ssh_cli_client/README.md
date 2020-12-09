# ssh_cli_client
The idea behind this code is that you could quickly gather info on the fly as well as make minor changes if so desired.
When the code is run, it will ask you how many devices you want to connect to,  as well as what commands you want to run.

The example below is using devnet hosts but I've updated my /etc/hosts file to make it easier to test.

You will need to add login credential to main.go.

If you have multiple logins you could use if/else logic in the code to match hostname you're logging into with appropriate username and password.

In reality you will more than likely be using tacacs and have standard login. So the base code is written as such.

# Running code with show commands

The devices used in the code below each had different usernames and passwords. The code worked fine. Just needed to tweak the code using if/else logic
with login creds.

```
$ go run main.go 
Number of hosts: 2
Hostname: fastxe:22

cmds: sh run int loopback75, show ip int brief

Hostname: nxos:8181

cmds: sh run int loopback78, show ip int brief


Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.


csr1000v-1#term len 0
csr1000v-1#sh run int loopback75
Building configuration...

Current configuration : 68 bytes
!
interface Loopback75
 ip address 75.75.75.75 255.255.255.255
end

csr1000v-1# show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       unassigned      YES NVRAM  administratively down down    
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
Loopback70             70.0.0.1        YES manual up                    up      
Loopback75             75.75.75.75     YES manual up                    up      
Loopback200            172.16.200.1    YES other  up                    up      
csr1000v-1#
csr1000v-1#exit

stty: standard input: Inappropriate ioctl for device
gl_set_term_size: NULL arguments(s).

!Command: show running-config interface loopback78
!Running configuration last done at: Sun Dec  6 10:08:02 2020
!Time: Sun Dec  6 10:10:39 2020

version 9.3(3) Bios:version  

interface loopback78
  ip address 78.78.78.78/32


IP Interface Status for VRF "default"(1)
Interface            IP Address      Interface Status
Vlan100              172.16.100.1    protocol-down/link-down/admin-down 
Vlan101              172.16.101.1    protocol-down/link-down/admin-down 
Vlan102              172.16.102.1    protocol-down/link-down/admin-down 
Vlan103              172.16.103.1    protocol-down/link-down/admin-down 
Vlan104              172.16.104.1    protocol-down/link-down/admin-down 
Vlan105              172.16.105.1    protocol-down/link-down/admin-down 
Lo1                  172.16.0.1      protocol-up/link-up/admin-up       
Lo78                 78.78.78.78     protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       
Eth1/5               172.16.1.1      protocol-down/link-down/admin-down 
$ 
```


# Running code with configuration commands + validation commands.

```
$ go run main.go
Number of hosts: 2
Hostname: nxos:8181

cmds: config t, interface loopback78, description go test script, exit , exit, show interface description

Hostname: fastxe:22

cmds: config t, interface loopback75, description go test script, exit, exit, show interface description


stty: standard input: Inappropriate ioctl for device
gl_set_term_size: NULL arguments(s).

-------------------------------------------------------------------------------
Interface                Description                                            
-------------------------------------------------------------------------------
mgmt0                    DO NOT TOUCH CONFIG ON THIS INTERFACE

-------------------------------------------------------------------------------
Port          Type   Speed   Description
-------------------------------------------------------------------------------
Eth1/1        eth    10G     --
Eth1/2        eth    10G     --
Eth1/3        eth    10G     --

<cropped for brevity>

-------------------------------------------------------------------------------
Interface                Description                                            
-------------------------------------------------------------------------------
Lo1                      --
Lo30                     My Learning Lab Loopback
Lo78                     go test script
Lo98                     Configured using OpenConfig Model
Lo99                     Full intf config via NETCONF

-------------------------------------------------------------------------------
Interface                Description                                            
-------------------------------------------------------------------------------
Vlan1                    --
Vlan100                  mgmt svi - DEMO PLEASE DON'T TOUCH
Vlan101                  prod svi - DEMO PLEASE DON'T TOUCH
Vlan102                  dev svi - DEMO PLEASE DON'T TOUCH
Vlan103                  test svi - DEMO PLEASE DON'T TOUCH
Vlan104                  security svi - DEMO PLEASE DON'T TOUCH
Vlan105                  iot svi - DEMO PLEASE DON'T TOUCH

Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.



csr1000v-1#term len 0
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)# interface loopback75
csr1000v-1(config-if)# description go test script
csr1000v-1(config-if)# exit
csr1000v-1(config)# exit
csr1000v-1# show interface description
Interface                      Status         Protocol Description
Gi1                            up             up       MANAGEMENT INTERFACE - DON'T TOUCH ME
Gi2                            admin down     down     Network Interface
Gi3                            admin down     down     Network Interface
Lo70                           up             up       
Lo75                           up             up       go test script
Lo200                          up             up       Added by CBT Nuggets
csr1000v-1#
csr1000v-1#exit
$ 

```

