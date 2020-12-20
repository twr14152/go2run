# go2run

This repo holds packages that can be used in your code to easily manage network devices.
At the time of publishing this Cisco Devnet devices IOS-XE and NXOS have been tested.

# runscript 
- Package that allows the user to call host and command files for scripting network devices
- Simply import the package into your code and customize for your environment

# runcli 
- Package that allow users to login to groups of devices issuing show/config commands
- Good for troubleshooting, gathering info quickly, and lab environments
- Import package into your code and your off to the races

# Unfortunately I've been having some issues with Go.mod with both of these go use go get <package.name> to install locally

