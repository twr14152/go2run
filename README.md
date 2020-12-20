# go2run

This repo holds packages that can be used in your code to easily manage network devices. The code was written in such away that it allows the user to determine the degree to which they influence their devices. At the time of publishing only Cisco Devnet devices IOS-XE and NXOS have been tested.

# runcli 
- Package that allow users to login to groups of devices and interactively issuing both show and config commands
- The groupings are based on common login credentials
- Good for troubleshooting, gathering info quickly, and lab environments
- See the README in runcli directory for instuctions

# runscript 
- The goal of this package is to allow the user to use host and command files to change and validate device configurations
- The amount of Go code the user has to write is minimal for the script to work
- See the README in the runscript directory for instructions



