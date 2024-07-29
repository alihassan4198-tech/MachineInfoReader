# MachineInfoReader
Read the Linux and Mac Information and Upload it the  to server
This program will create CSV files and periodically upload them.

## To Run

### Step 1) 
Download Go 1.22 w.r.t your Machine [Downlod Here](https://go.dev/dl/)


### Step 2) 
#### for other OS ([Installation Instructions](https://go.dev/doc/install))
Insllation Instruction (Linux) 
```
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.<MACHINE-TYPE>.tar.gz
```


### Step 3) 
Add /usr/local/go/bin to the PATH environment variable.
You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
```
export PATH=$PATH:/usr/local/go/bin
```

### Step 4) 
Verify that you've installed Go by opening a command prompt and typing the following command:
```
go version
```


### Step 5)
Change current directory to "MachineInfoReader".

Build Binaries for your machine from source
```
source ./build-binaries.sh
```

### Step 6)
Install Binaries, Cron Job and Service fetching machine information.
```bash
sudo ./installation.sh
```



### Step ( Un-Installation )
```
sudo ./uninstallation.sh
```



### Note : Linux CSV Output Path = "/home/machineinfoserver"

### Note : Mac CSV Output Path = "/Users/Shared/machineinfoserver"

