<p align="center">
	 <br>
    Fully Decentralized. Immutable. Portable. Easy to use Statically compiled VPN and a reverse proxy over p2p.<br>
    <b>VPN</b> -  <b>Reverse Proxy</b> - <b>Send files securely over p2p</b> -  <b>Blockchain</b>
</p>


WarpNet uses libp2p to build private decentralized networks that can be accessed via shared secrets.

It can:

- **Create a VPN** :  Secure VPN between p2p peers
  - Automatically assign IPs to nodes
  - Embedded tiny DNS server to resolve internal/external IPs
  - Create trusted zones to prevent network access if token is leaked

- **Act as a reverse Proxy** : Share a tcp service like you would do with `ngrok`. WarpNet let expose TCP services to the p2p network nodes without establishing a VPN connection: creates reverse proxy and tunnels traffic into the p2p network.

- **Send files via p2p** : Send files over p2p between nodes without establishing a VPN connection.

- **Be used as a library**: Plug a distributed p2p ledger easily in your golang code!

See the [documentation](https://thealonemusk.github.io/WarpNet).

# :camera: Screenshots

Dashboard (Dark mode)            |  Dashboard (Light mode)
:-------------------------:|:-------------------------:
![Screenshot 2021-10-31 at 00-12-16 WarpNet - Machines index](https://user-images.githubusercontent.com/2420543/163020448-8e9238c1-3b6d-435d-9b25-7729d8779ebd.png) | ![Screenshot 2021-10-31 at 23-03-26 WarpNet - Machines index](https://user-images.githubusercontent.com/2420543/163020460-e18c07d7-8426-4992-aab3-0b2fd90279ae.png)

DNS            |  Machine index
:-------------------------:|:-------------------------:
![Screenshot 2021-10-31 at 23-03-44 WarpNet - Services index](https://user-images.githubusercontent.com/2420543/163020465-3d481da4-4912-445e-afc0-2614966dcadf.png) | ![Screenshot 2021-10-31 at 23-03-59 WarpNet - Files index](https://user-images.githubusercontent.com/2420543/163020462-7821a622-8c13-4971-8abe-9c5b6b491ae8.png)

Services            |  Blockchain index
:-------------------------:|:-------------------------:
![Screenshot 2021-10-31 at 23-04-12 WarpNet - Users connected](https://user-images.githubusercontent.com/2420543/163021285-3c5a980d-2562-4c10-b266-7e99f19d8a87.png) | ![Screenshot 2021-10-31 at 23-04-20 WarpNet - Blockchain index](https://user-images.githubusercontent.com/2420543/163020457-77ef6e50-40a6-4e3b-83c4-a81db729bd7d.png)


# :new: GUI

A Desktop GUI application (alpha) for Linux is available [here](https://github.com/thealonemusk/WarpNet-gui)

Dashboard            |  Connections index
:-------------------------:|:-------------------------:
![WarpNet-gui-2](https://user-images.githubusercontent.com/2420543/147854909-a223a7c1-5caa-4e90-b0ac-0ae04dc0949d.png) | ![WarpNet-3](https://user-images.githubusercontent.com/2420543/147854904-09d96991-8752-421a-a301-8f0bdd9d5542.png)
![WarpNet-gui](https://user-images.githubusercontent.com/2420543/147854907-1e4a4715-3181-4dc2-8bc0-d052b3bf46d3.png) | 

# Kubernetes 

Check out [c3os](https://github.com/thealonemusk/c3os) for seeing WarpNet in action with Kubernetes!

# :running: Installation

Download the precompiled static release in the [releases page](https://github.com/thealonemusk/WarpNet/releases). You can either install it in your system or just run it.

# :computer: Usage

WarpNet works by generating tokens (or a configuration file) that can be shared between different machines, hosts or peers to access to a decentralized secured network between them.

Every token is unique and identifies the network,  no central server setup, or specifying hosts ip is required.

To generate a config run:

```bash
# Generate a new config file and use it later as WarpNetCONFIG
$ WarpNet -g > config.yaml
```

OR to generate a portable token:

```bash
$ WarpNetTOKEN=$(WarpNet -g -b)
```

Note, tokens are config merely encoded in base64, so this is equivalent:

```bash
$ WarpNetTOKEN=$(WarpNet -g | tee config.yaml | base64 -w0)
```

All WarpNet commands implies that you either specify a `WarpNetTOKEN` (or `--token` as parameter) or a `WarpNetCONFIG` as this is the way for `WarpNet` to establish a network between the nodes. 

The configuration file is the network definition and allows you to connect over to your peers securely.

**Warning** Exposing this file or passing-it by is equivalent to give full control to the network.

## :satellite: As a VPN

To start the VPN, simply run `WarpNet` without any argument.

An example of running WarpNet on multiple hosts:

```bash
# on Node A
$ WarpNetTOKEN=.. WarpNet --address 10.1.0.11/24
# on Node B
$ WarpNetTOKEN=.. WarpNet --address 10.1.0.12/24
# on Node C ...
$ WarpNetTOKEN=.. WarpNet --address 10.1.0.13/24
...
```

... and that's it! the `--address` is a _virtual_ unique IP for each node, and it is actually the ip where the node will be reachable to from the vpn. You can assign IPs freely to the nodes of the network, while you can override the default `WarpNet0` interface with `IFACE` (or `--interface`)

*Note*: It might take up time to build the connection between nodes. Wait at least 5 mins, it depends on the network behind the hosts.


# :question: Is it for me?

WarpNet makes VPN decentralization a first strong requirement. 

Its main use is for edge and low-end devices and especially for development.

The decentralized approach has few cons:

- The underlying network is chatty. It uses a Gossip protocol for synchronizing the routing table and p2p. Every blockchain message is broadcasted to all peers, while the traffic is to the host only.
- Might be not suited for low latency workload.

Keep that in mind before using it for your prod networks!

But it has a strong pro: it just works everywhere libp2p works!

# :question: Why? 

First of all it's my first experiment with libp2p. Second, I always wanted a more "open" `ngrok` alternative, but I always prefer to have "less infra" as possible to maintain. That's why building something like this on top of `libp2p` makes sense.

# :warning: Warning!

I'm not a security expert, and this software didn't went through a full security audit, so don't use and rely on it for sensible traffic and not even for production environment! I did this mostly for fun while I was experimenting with libp2p. 

## Example use case: network-decentralized [k3s](https://github.com/k3s-io/k3s) test cluster

Let's see a practical example, you are developing something for kubernetes and you want to try a multi-node setup, but you have machines available that are only behind NAT (pity!) and you would really like to leverage HW.

If you are not really interested in network performance (again, that's for development purposes only!) then you could use `WarpNet` + [k3s](https://github.com/k3s-io/k3s) in this way:

1) Generate WarpNet config: `WarpNet -g > vpn.yaml`
2) Start the vpn:

   on node A: `sudo IFACE=WarpNet0 ADDRESS=10.1.0.3/24 WarpNetCONFIG=vpn.yml WarpNet`
   
   on node B: `sudo IFACE=WarpNet0 ADDRESS=10.1.0.4/24 WarpNetCONFIG=vpm.yml WarpNet`
3) Start k3s:
 
   on node A: `k3s server --flannel-iface=WarpNet0`
   
   on node B: `K3S_URL=https://10.1.0.3:6443 K3S_TOKEN=xx k3s agent --flannel-iface=WarpNet0 --node-ip 10.1.0.4`

We have used flannel here, but other CNI should work as well.


# :notebook: As a library

WarpNet can be used as a library. It is very portable and offers a functional interface.

To join a node in a network from a token, without starting the vpn:

```golang

import (
    node "github.com/thealonemusk/WarpNet/pkg/node"
)

e := node.New(
    node.Logger(l),
    node.LogLevel(log.LevelInfo),
    node.MaxMessageSize(2 << 20),
    node.FromBase64( mDNSEnabled, DHTEnabled, token ),
    // ....
  )

e.Start(ctx)

```

or to start a VPN:

```golang

import (
    vpn "github.com/thealonemusk/WarpNet/pkg/vpn"
    node "github.com/thealonemusk/WarpNet/pkg/node"
)

opts, err := vpn.Register(vpnOpts...)
if err != nil {
	return err
}

e := WarpNet.New(append(o, opts...)...)

e.Start(ctx)
```

# 🧑‍💻 Projects using WarpNet

- [Kairos](https://github.com/kairos-io/kairos) - creates Kubernetes clusters with K3s automatically using WarpNet networks


# 🐜 Contribution

You can improve this project by contributing in following ways:

- report bugs
- fix issues
- request features
- asking questions (just open an issue)

and any other way if not mentioned here.

# :notebook: Credits

- The awesome [libp2p](https://github.com/libp2p) library
- [https://github.com/songgao/water](https://github.com/songgao/water) for tun/tap devices in golang
- [Room example](https://github.com/libp2p/go-libp2p/tree/master/examples/chat-with-rendezvous) (shamelessly parts are copied by)
- Logo originally made by [Uniconlabs](https://www.flaticon.com/authors/uniconlabs) from [www.flaticon.com](https://www.flaticon.com/), modified by me

# :notebook: Troubleshooting

If during bootstrap you see messages like:

```
WarpNet[3679]:             * [/ip4/104.131.131.82/tcp/4001] failed to negotiate stream multiplexer: context deadline exceeded     
```

or

```
WarpNet[9971]: 2024/12/16 20:56:34 failed to sufficiently increase receive buffer size (was: 208 kiB, wanted: 2048 kiB, got: 416 kiB). See https://github.com/lucas-clemente/quic-go/wiki/UDP-Receive-Buffer-Size for details.
```

or generally experiencing poor network performance, it is recommended to increase the maximum buffer size by running:

```
sysctl -w net.core.rmem_max=2500000
```

# :notebook: TODO

- [x] VPN
- [x] Send and receive files via p2p
- [x] Expose remote/local services via p2p tunnelling
- [x] Store arbitrary data on the blockchain
- [x] Allow to persist blockchain on disk

