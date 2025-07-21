# go2run - Network Automation library

This repo holds packages that can be used in your code to easily manage network devices. The code was written in such away that it allows the user to determine the degree to which they influence their devices. The code is written so you don't have to write a bunch of boiler plate code. Test platforms include Cisco Devnet devices IOS, IOS-XE and NXOS and Arista cEOS devices on containerlab. The library as of now uses the ssh protocol. Hoping to expand that to http but time will tell.

The ssh cipher have been modified so that legacy devices can be connected to as well newer devices. This was done by adding aes128-cbc to the list of allowed ciphers. It should be noted that this is considered a vulnerability. If thats a problem for you don't use this code. That said there are a lot of devices out there that cannot have the ciphers upgraded, and they are still being used. You may as well do your best to manage them until they can be replaced. Or if this is a lab, you now should have access to your gear via ssh.

Testing environment for this has been Cisco Devnet always on network devices as well as older Cisco72xx-K9 images used in GNS3 and Arista cEOS-lab images.
Planning to expand to other vendor platforms at some point.


# runcli 
- Package that allow users to login to groups of devices and interactively issuing both show and config commands
- The groupings are based on common login credentials
- Good for troubleshooting, gathering info quickly, and lab environments

# runscript 
- This package will allow the user make changes and validate device configurations using host and command files
- The amount of Go code the user has to write is minimal for the script to work

# runscp
- This package will allow users to scp file over to a remote device



