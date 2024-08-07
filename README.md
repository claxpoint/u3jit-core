# u3jit-core
A way to be free

Our SubProjects from u3jit-core:
# NetBaan: Sucker5
# socks5-vpn
This is just a combination of [dante](https://github.com/vimagick/dockerfiles/tree/master/dante) and [dockvpn](https://github.com/umputun/dockvpn)
Docker compose socks5 and vpn server

# Configuration
There are default ports in a file 'docker-compose' - 1088/tcp, 1194/udp, 443/tcp for OpenVPN, 8088/tcp, 9443/tcp/udp (actualy I don't know udp or tcp and I opened them all) for MTProto proxy - which must be opened

# How to run
* First of all, you need to open ports which defined in a file 'docker-compose'
* Install docker, docker-compose
* git clone this repository
* for configuration openvpn follow [this](https://github.com/kylemanna/docker-openvpn/blob/master/docs/docker-compose.md) instraction
* ```docker-compose up -d```

For checking socks5 ```curl -x socks5h://suck-rkn:telegram@127.0.0.1:1088 https://www.youtube.com```

# socks proxy
In [docker-compse](docker-compose.yml#L17) define `SOCKS_USERNAME`, `SOCKS_PASSWORD` for SOCKS5 proxy.
`sockd.conf` have all configuration for dante-server (SOCKS5)


# mtproxy (MTProto Proxy)
You should generate some random 32 hex strings ([for example here](https://www.browserling.com/tools/random-hex)) and define them in [config_mtproxy_server.py](config_mtproxy_server.py#L4).
After docker container up you could read secret keys from ```docker-compose logs mtproxy```


# NetBaan: TGP2VPN
A new way to use telegram proxies to bypass internet filtering - with U3jit FilterBreak System
and find a line like ```tg: YYY.....``` YYY is your key for a mtproxy
