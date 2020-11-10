# Networks

## Netplan

#### Example
```yaml
# This is the network config written by 'yannokgo'
network:
  renderer: networkd
  ethernets:
    ens33:
      addresses: [192.168.19.31/24]
      gateway4: 192.168.19.2
      nameservers:
        addresses: [8.8.8.8, 1.1.1.1]
        search: [labo.local]
      dhcp4: false
      dhcp6: false
  version: 2
```

#### usage
```bash
netplan generate
netplan try
netplan apply
```

```
test
```


