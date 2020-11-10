SSH
===

Install, Enable, Start
```bash
sudo apt-get install openssh-server
sudo systemctl enable ssh
sudo systemctl start ssh
```

Connect
```bash
ssh user@server-name
```

SSH Server
----------
```bash
sudo apt-get install openssh-server 
```    

**Back-Up**

First, make a backup of your sshd\_config file by copying it to your home directory, or by making a read-only copy in /etc/ssh
```bash
sudo cp /etc/ssh/sshd_config /etc/ssh/sshd_config.factory-defaults
sudo chmod a-w /etc/ssh/sshd_config.factory-defaults
```    

Revert
```bash
sudo gedit /etc/ssh/sshd_config
```    

Disable password authentication in `sshd_config` replace by the second
```
#PasswordAuthentication yes
PasswordAuthentication no
``` 

**Generating Keys** Client-side. Create your public and private SSH keys
```bash
mkdir ~/.ssh
chmod 700 ~/.ssh
ssh-keygen -t rsa -b 4096
```    

[more on ubuntu.com](https://help.ubuntu.com/community/SSH/OpenSSH/Keys "Ubuntu.com")

ufw - Firewall ports
--------------------
```bash
sudo ufw status
sudo ufw status verbose
    

sudo ufw enable
sudo ufw allow 22/tcp comment 'Open port ssh tcp port 22'
```    

Open incoming SSH but deny connections from an IP address that has attempted to initiate 6 or more connections in the last 30 seconds.
```bash
$ sudo ufw limit ssh
```    

Credits
-------

[Open ssh 22/tcp using ufw](https://www.cyberciti.biz/faq/ufw-allow-incoming-ssh-connections-from-a-specific-ip-address-subnet-on-ubuntu-debian/ "Cyberciti.biz")

