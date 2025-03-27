package config

import (
	"errors"
	"fmt"

	peer "github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

// DefaultBootstrapAddresses are the hardcoded bootstrap addresses
// for IPFS. they are nodes run by the IPFS team. docs on these later.
// As with all p2p networks, bootstrap is an important security concern.
//
// NOTE: This is here -- and not inside cmd/btfs/init.go -- because of an
// import dependency issue. TODO: move this into a config/default/ package.
var DefaultBootstrapAddresses = []string{
	"/ip4/63.176.242.235/tcp/4001/p2p/16Uiu2HAmVeJwSMkeaEXEZdDAtxM6mngAALjTPwq4w2suehMVPwA5",
	"/ip4/3.66.98.120/tcp/4001/p2p/16Uiu2HAmVSpShqGg8c7dEuG8qSWZjisx1rxNFwgAAi47HKHHXFr4",
	"/ip4/54.179.164.197/tcp/4001/p2p/16Uiu2HAmVZP7ueF6jkSsMvqPEZgGGEMURtn5MZiz1KGP7FQWvHU8",
	"/ip4/18.138.163.50/tcp/4001/p2p/16Uiu2HAm2RX2aMwcHMsEvLzqQ76Jm5bcK9Ut869pipZ9UPSuG9zB",
	"/ip4/15.184.198.54/tcp/4001/p2p/16Uiu2HAmU6iej57dbFD1qEcG2UtSBc1KUXgSAph23AcagDsLBhxW",
	"/ip4/15.185.79.232/tcp/4001/p2p/16Uiu2HAmQ2Gjyjevt1MhZm7zmQiBjXSdoupLTc6dLLotwBJ7jian",
	"/ip4/3.7.21.138/tcp/4001/p2p/16Uiu2HAm7QD77kxSKf1GTM3YkrYp8vkhUwS2ySJPht9jALeaHaft",
	"/ip4/43.204.199.237/tcp/4001/p2p/16Uiu2HAm3tpaz9zgqB4i2FEwX7dwTJzv88Krpdy3kRecXZos3WdM",
	"/ip4/35.155.192.241/tcp/4001/p2p/16Uiu2HAm29iAxcKRPNRBVMYCz455uck5o7KmdPJ9GQ5BKvpxxca9",
	"/ip4/35.83.203.96/tcp/4001/p2p/16Uiu2HAmNnKCdkBKdoPo4sXSLhDgXvPmCi7NCjo8cfcP5RRb4mKL",
	"/ip4/35.72.132.60/tcp/4001/p2p/16Uiu2HAmSwJux2LgfMjQn8CzcG8jufKHbkRm9fQAqEASumU2R38h",
	"/ip4/52.198.239.158/tcp/4001/p2p/16Uiu2HAmQ6YjTL2LCxiYjpRXuukchFjDwm5p7HRjj6nLGcrjwEsL",
	"/ip4/13.57.144.203/tcp/4001/p2p/16Uiu2HAm5SdP8So2MpknrEAi6avMVXokEuujynKkXRPnxaTxtn5W",
	"/ip4/54.176.58.20/tcp/4001/p2p/16Uiu2HAmPLc4QAuN6gmSUZKYWF1Kh22TFxTANFQYrKHmhh31hikk",
	"/ip4/3.130.97.111/tcp/4001/p2p/16Uiu2HAm5Kd8T7GphFK2kUwfQXEPa6nKwYtkzZKr2R2VoB3PBhNF",
	"/ip4/3.131.15.203/tcp/4001/p2p/16Uiu2HAm4afRu4ny2rcseF34JYNut1TLoenfufAR8SDiDHPZz6Zz",
	"/ip4/3.133.9.176/tcp/4001/p2p/16Uiu2HAmGXv3CDUK8FpV1rBCpGtGsV3rxGshcgBiHTDxYHrbic2Q",
	"/ip4/3.17.210.15/tcp/4001/p2p/16Uiu2HAmSiBrUTxeYj9FXxeHYcboTz8NMyPgDkfjB5yAag3opKur",
	"/ip4/13.40.252.137/tcp/4001/p2p/16Uiu2HAmKM8d2mZ4yNHJ1GZj3FDVD1BXE8Qg7TUBMexkbMGWWbnR",
	"/ip4/13.43.125.230/tcp/4001/p2p/16Uiu2HAkz6Uz3fv2vHPEmn5He89emRZmiToJg7fZaG7MwUob7qr1",
}
var DefaultTestnetBootstrapAddresses = []string{
	"/ip4/18.224.174.215/tcp/45301/p2p/16Uiu2HAmFFwNdgSoLhfgJUPEfPEVodppRxaeZBVpAvrH5s3qSkWo",
	"/ip4/18.224.174.215/tcp/34237/p2p/16Uiu2HAmDigS3SDx6g9Sp6MUfdFHvDwS8Zw8E14V6bLhCAHA3jjB",
	"/ip4/18.224.174.215/tcp/43097/p2p/16Uiu2HAm7HQoEbQe1fYt4LtnG6z5TqwTrrqUv5xsnt4nukskWmAi",
	"/ip4/18.224.174.215/tcp/38955/p2p/16Uiu2HAm5WrYvkJwaRP7ZAroWCfjaUxKkNssqcSmEmKJ8vXVYp1o",
	"/ip4/54.151.185.243/tcp/36707/p2p/16Uiu2HAmDis3wAorW46YyNmXNk963VAAHwZ1phjHXj5yduyawAUy",
	"/ip4/54.151.185.243/tcp/42741/p2p/16Uiu2HAmSfqLCyqH5qQQF8zpzPMQvWiQunhWpYtSxwGw5QR2jhgU",
	"/ip4/54.151.185.243/tcp/37403/p2p/16Uiu2HAmBHwyRUETsGqjYpgPRpnMC9y39tcVYH6vKxZidCBcBeFG",
	"/ip4/54.151.185.243/tcp/37739/p2p/16Uiu2HAm2oKy37KvYmiv1nnRWZwUoLPZumNKFxPzhM1t8F3KxADu",
	"/ip4/18.158.67.141/tcp/40155/p2p/16Uiu2HAmTMEqndByECXuxk1Rg8szxMqwS3tUFFWhAUduFzwfwmfK",
	"/ip4/18.158.67.141/tcp/44569/p2p/16Uiu2HAmL4QNi68nSNbedUWp1A1cRR3z3NuJqQYmAYoj19ht6iNv",
	"/ip4/18.158.67.141/tcp/39703/p2p/16Uiu2HAkzF6JMx4EL2C4cLoCLyQH8t1sgyttQxPfQtNt5FZhvpxs",
	"/ip4/18.158.67.141/tcp/46713/p2p/16Uiu2HAm85HXJA7xmgNxxTVdFRuRCGstvrY8nW6KqfTtkuZrZg64",
	"/ip4/18.163.235.175/tcp/36335/p2p/16Uiu2HAm8wVUsVsqksBfxy6yzHpVv5gELQnpU7Q2uhDyXFwr9bfV",
	"/ip4/18.163.235.175/tcp/44029/p2p/16Uiu2HAmBvnQU5FWgEcfY1jaAK2Q9iQBy6FwQdDUtyT7mo8HU1Yu",
	"/ip4/18.163.235.175/tcp/40191/p2p/16Uiu2HAkurshicwtTrqbrL3yv9xR7hogPvreUHJP3W8n9W5XMibz",
}

// ErrInvalidPeerAddr signals an address is not a valid peer address.
var ErrInvalidPeerAddr = errors.New("invalid peer address")

func (c *Config) BootstrapPeers() ([]peer.AddrInfo, error) {
	return ParseBootstrapPeers(c.Bootstrap)
}

// DefaultBootstrapPeers returns the (parsed) set of default bootstrap peers.
// if it fails, it returns a meaningful error for the user.
// This is here (and not inside cmd/btfs/init) because of module dependency problems.
func DefaultBootstrapPeers() ([]peer.AddrInfo, error) {
	ps, err := ParseBootstrapPeers(DefaultBootstrapAddresses)
	if err != nil {
		return nil, fmt.Errorf(`Failed to parse hardcoded bootstrap peers: %s
This is a problem with the BTFS codebase.
Please report it to https://github.com/bittorrent/go-btfs/issues.`, err)
	}
	return ps, nil
}
func DefaultTestnetBootstrapPeers() ([]peer.AddrInfo, error) {
	ps, err := ParseBootstrapPeers(DefaultTestnetBootstrapAddresses)
	if err != nil {
		return nil, fmt.Errorf(`Failed to parse hardcoded testnet bootstrap peers: %s
This is a problem with the BTFS codebase.
Please report it to https://github.com/bittorrent/go-btfs/issues.`, err)
	}
	return ps, nil
}

func (c *Config) SetBootstrapPeers(bps []peer.AddrInfo) {
	c.Bootstrap = BootstrapPeerStrings(bps)
}

// ParseBootstrapPeer parses a bootstrap list into a list of AddrInfos.
func ParseBootstrapPeers(addrs []string) ([]peer.AddrInfo, error) {
	maddrs := make([]ma.Multiaddr, len(addrs))
	for i, addr := range addrs {
		var err error
		maddrs[i], err = ma.NewMultiaddr(addr)
		if err != nil {
			return nil, err
		}
	}
	return peer.AddrInfosFromP2pAddrs(maddrs...)
}

// BootstrapPeerStrings formats a list of AddrInfos as a bootstrap peer list
// suitable for serialization.
func BootstrapPeerStrings(bps []peer.AddrInfo) []string {
	bpss := make([]string, 0, len(bps))
	for _, pi := range bps {
		addrs, err := peer.AddrInfoToP2pAddrs(&pi)
		if err != nil {
			// programmer error.
			panic(err)
		}
		for _, addr := range addrs {
			bpss = append(bpss, addr.String())
		}
	}
	return bpss
}
