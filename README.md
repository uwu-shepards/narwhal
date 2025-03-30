# narwhal (unicorn testnet) 🦄

![image](https://github.com/user-attachments/assets/8c0a6698-4926-49f2-8bbe-eb8c470e608f)

### Official RPCs
```
https://rpc.uwupunks.com
https://narwhal.uwupunks.com/
```

### Peers
```
6b37a66a808b986a4cb33c0249166010bf0b32bf@p2p.uwupunks.com:26656
346557c007aeb7fc28ce8a53ac8322faf8d087f7@kaijunicorn.uwupunks.com:26656
```

Never forget what they took from us

▄︻デ ══‐一 ❤️

▄︻デ ══‐一 ❤️

▄︻デ ══‐一 ❤️

Censorship is the enemy.

### Install

`curl https://get.ignite.com/uwupunks/narwhal! | sudo bash`

`wget https://github.com/CosmWasm/wasmvm/releases/download/v2.2.1/libwasmvm.x86_64.so`

```
sudo mv libwasmvm.x86_64.so /usr/local/lib/
sudo chmod 755 /usr/local/lib/libwasmvm.x86_64.so
sudo ldconfig
ldconfig -p | grep libwasmvm
```

`narwhal init uwupunks`

```
nano ~/.narwhal/config/app.toml
minimum-gas-prices = "0narwhal"
```

`nano ~/.narwhal/config/config.toml `
```
seeds = "6b37a66a808b986a4cb33c0249166010bf0b32bf@p2p.uwupunks.com:26656"
persistent_peers = "6b37a66a808b986a4cb33c0249166010bf0b32bf@p2p.uwupunks.com:26656,346557c007aeb7fc28ce8a53ac8322faf8d087f7@kaijunicorn.uwupunks.com:26656" 
```

```
cd ~/.narwhal/config
curl -s https://rpc.uwupunks.com/genesis | jq -r '.result.genesis' > genesis.json
```

# Start the chain
`narwhal start`


# Running as a service

You may wish to run the chain as a `systemctl` service so that it runs on boot and logs can be accessed via `journalctl`. Below is a unit file you can use:

`/etc/systemd/system/narwhal.service`
```
[Unit]
Description=Narwhal Full Node Service
After=network-online.target
Wants=network-online.target

[Service]
User=<your user name>
Group=<your group>
ExecStart=/usr/local/bin/narwhal start
Restart=always
RestartSec=3
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
```
