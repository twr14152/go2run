# runscp

This package will allow you to scp files to a remote host.

# To Use:

Create a file main.go.

```
package main

import (
	"fmt"
	"github.com/twr14152/go2run/runscp"
)

func main() {
	// You can add as many hosts as needed
	fmt.Println("scp file to:sandbox-iosxe-latest-1.cisco.com")
	runscp.RunScp("username1", "password1", "sandbox-iosxe-latest-1.cisco.com:22", "testfile.txt")
	//
	fmt.Println("scp file to:ios-xe-mgmt.cisco.com")
	runscp.RunScp("username2", "password2", "ios-xe-mgmt.cisco.com:8181", "testfile.txt")

}

```
# To install the files needed for this repo
```
$go mod init <executible_filename> // this command is unnessary if go.mod file already exists
$go mod tidy //This will actually download the dependency files listed in the main.go import statements
```

# File to transfer:

In this example we created testfile.txt. Its in the same directory as main.go file.

```
This is a file transfer test



1
2
3
4
5

done!

```
```
$ ls -l
total 8
-rw-r--r-- 1 runner runner 245 Dec 21 12:35 main.go
-rw-r--r-- 1 runner runner  48 Dec 21 12:25 testfile.txt
$

```

# Run the script
```
$  go run main.go
scp file to:sandbox-iosxe-latest-1.cisco.com
success
scp file to:ios-xe-mgmt.cisco.com
success
```
# Verify the results

To keep this short only showing verification of one host.
```
$ ssh developer@sandbox-iosxe-latest-1.cisco.com
The authenticity of host 'sandbox-iosxe-latest-1.cisco.com (131.226.217.143)' can't be established.
RSA key fingerprint is SHA256:eun4Pkw8sSwnGENqUeO8iB8UqLrHI+vsT0nlEiW+ps8.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added 'sandbox-iosxe-latest-1.cisco.com,131.226.217.143' (RSA) to the list of known hosts.

Lab powered by Network Chuck and Keith Barker
Password: 

Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.



csr1000v-1#dir
Directory of bootflash:/

24      -rw-               48  Dec 21 2020 12:35:26 +00:00  testfile.txt
201601  drwx            20480  Dec 21 2020 09:03:00 +00:00  tracelogs
23      -rw-           114709  Dec 21 2020 06:37:14 +00:00  Line
80641   drwx             4096  Dec 19 2020 16:04:33 +00:00  .installer
137089  drwx             4096  Dec 19 2020 16:03:52 +00:00  license_evlog
22      -rw-              157  Dec 19 2020 16:03:51 +00:00  csrlxc-cfg.log
19      -rw-             2288  Dec 19 2020 16:03:49 +00:00  cvac.log
18      -rw-               30  Dec 19 2020 16:03:47 +00:00  throughput_monitor_params
15      -rw-             1216  Dec 19 2020 16:02:38 +00:00  mode_event_log
64513   drwx             4096   Sep 1 2020 14:51:38 +00:00  .dbpersist
274177  drwx             4096   Sep 1 2020 14:51:34 +00:00  onep
21      -rw-               16   Sep 1 2020 14:51:32 +00:00  ovf-env.xml.md5
20      -rw-                1   Sep 1 2020 14:51:32 +00:00  .cvac_version
104833  drwx             4096   Sep 1 2020 14:51:29 +00:00  pnp-info
145153  drwx             4096   Sep 1 2020 14:50:48 +00:00  virtual-instance
17      -rwx             1314   Sep 1 2020 14:50:21 +00:00  trustidrootx3_ca.ca
16      -rw-            20109   Sep 1 2020 14:50:21 +00:00  ios_core.p7b
193537  drwx             4096   Sep 1 2020 14:50:18 +00:00  gs_script
40321   drwx             4096   Sep 1 2020 14:50:16 +00:00  core
169345  drwx             4096   Sep 1 2020 14:50:12 +00:00  bootlog_history
161281  drwx             4096   Sep 1 2020 14:50:07 +00:00  .prst_sync
14      -rw-             1105   Sep 1 2020 14:49:08 +00:00  packages.conf
13      -rw-         48321761   Sep 1 2020 14:49:08 +00:00  csr1000v-rpboot.17.03.01a.SPA.pkg
12      -rw-        470611036   Sep 1 2020 14:49:08 +00:00  csr1000v-mono-universalk9.17.03.01a.SPA.pkg
8065    drwx             4096   Sep 1 2020 14:49:03 +00:00  .rollback_timer
11      drwx            16384   Sep 1 2020 14:48:15 +00:00  lost+found

6286540800 bytes total (5430255616 bytes free)
csr1000v-1#more testfile.txt
This is a file transfer test



1
2
3
4
5

done!
csr1000v-1#delete testfile.txt
Delete filename [testfile.txt]? 
Delete bootflash:/testfile.txt? [confirm]
csr1000v-1#
```
