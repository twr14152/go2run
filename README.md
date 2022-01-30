# go2run

This repo holds packages that can be used in your code to easily manage network devices. The code was written in such away that it allows the user to determine the degree to which they influence their devices. At the time of publishing only Cisco Devnet devices IOS-XE and NXOS have been tested. These packages do not play well with go.mod. As a work-a-round to that issue delete go.mod and issue go get on the repo desired and run your code. 

I will revisit the go.mod issues as time permits.

# runcli 
- Package that allow users to login to groups of devices and interactively issuing both show and config commands
- The groupings are based on common login credentials
- Good for troubleshooting, gathering info quickly, and lab environments

# runscript 
- This package will allow the user make changes and validate device configurations using host and command files
- The amount of Go code the user has to write is minimal for the script to work


# runscp
- This package will allow users to scp file over to a remote device

