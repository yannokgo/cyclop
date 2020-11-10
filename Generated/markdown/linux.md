
### Linux - General


Username `uname -a -v -r`

ip       `ip addr show eht0` or `ip a`

disk free humain readable `df -ah`

directory size  `du -sh dir`

Ports  `netstat -tulpn`

CPU usage of the nginx process `ps aux | grep nginx` or `top` or `htop`



### Microphone record and loopback


```bash
arecord -f cd - | aplay -

arecord -f cd - | tee output.wav | aplay -
```

Initialize a fresh install


```bash
sudo apt-get update && sudo apt-get upgrade
sudo apt install timeshift
sudo timeshift --create
sudo snap install
```

