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
	"/ip4/63.176.242.235/tcp/4001/p2p/16Uiu2HAmT9HSazxmnS4ucPY3Zpq2B5NT3wbiJi9ETurkVGGpxa57",
	"/ip4/3.66.98.120/tcp/4001/p2p/16Uiu2HAmVSpShqGg8c7dEuG8qSWZjisx1rxNFwgAAi47HKHHXFr4",
	"/ip4/54.179.164.197/tcp/4001/p2p/16Uiu2HAmVZP7ueF6jkSsMvqPEZgGGEMURtn5MZiz1KGP7FQWvHU8",
	"/ip4/18.138.163.50/tcp/4001/p2p/16Uiu2HAm2RX2aMwcHMsEvLzqQ76Jm5bcK9Ut869pipZ9UPSuG9zB",
	"/ip4/18.140.91.213/tcp/4001/p2p/16Uiu2HAmTmxfiGRkKSYjcnXz6mK4mHZAC5kPHdLYZaVcEcpuMG7q",
	"/ip4/52.220.30.16/tcp/4001/p2p/16Uiu2HAmRaW1vufkhRG6uH4Fm7hARkRWf3SKzyvWmL8mTunE8o7U",
	"/ip4/157.175.142.101/tcp/4001/p2p/16Uiu2HAmR2xYRccvbMwp3kY719AbKFuwXXZN5UdqEZ7WspcYKbhe",
	"/ip4/157.241.81.109/tcp/4001/p2p/16Uiu2HAkx39tHJ32ijBAzBmnPwAYqXyeoJF8Q8TFhw9hfxjvVGAh",
	"/ip4/15.184.198.54/tcp/4001/p2p/16Uiu2HAmU6iej57dbFD1qEcG2UtSBc1KUXgSAph23AcagDsLBhxW",
	"/ip4/15.185.79.232/tcp/4001/p2p/16Uiu2HAmQ2Gjyjevt1MhZm7zmQiBjXSdoupLTc6dLLotwBJ7jian",
	"/ip4/16.24.14.84/tcp/4001/p2p/16Uiu2HAmM96uUH53Ab9JBWfuwUBXJvGMbfVbsBXiGZGqStP93DTS",
	"/ip4/16.24.16.4/tcp/4001/p2p/16Uiu2HAmJ6vEtzmmC6nM6SJwHA9NCPwTRWy7K5WT2UFXDqzJFGSf",
	"/ip4/3.76.64.148/tcp/4001/p2p/16Uiu2HAmFc3snGkwK76yMYMAkHWhq6GD29w7m8Sa7kUciUK5xovu",
	"/ip4/3.78.178.244/tcp/4001/p2p/16Uiu2HAmHeUHakzYG1YWfWoSriVwKhSHYz88rL3USmgeRpqtWqMw",
	"/ip4/3.7.21.138/tcp/4001/p2p/16Uiu2HAm7QD77kxSKf1GTM3YkrYp8vkhUwS2ySJPht9jALeaHaft",
	"/ip4/43.204.199.237/tcp/4001/p2p/16Uiu2HAm3tpaz9zgqB4i2FEwX7dwTJzv88Krpdy3kRecXZos3WdM",
	"/ip4/35.155.192.241/tcp/4001/p2p/16Uiu2HAm29iAxcKRPNRBVMYCz455uck5o7KmdPJ9GQ5BKvpxxca9",
	"/ip4/35.83.203.96/tcp/4001/p2p/16Uiu2HAmNnKCdkBKdoPo4sXSLhDgXvPmCi7NCjo8cfcP5RRb4mKL",
	"/ip4/54.69.57.58/tcp/4001/p2p/16Uiu2HAmNDZWZtyRNZMLQ88933SFcVp2gtb99aQVbADXcCFcjFn9",
	"/ip4/35.164.151.55/tcp/4001/p2p/16Uiu2HAmMgufksaU9aaenq2bNtGnG5QokCS1xdzJwUS6yRtakhbs",
	"/ip4/35.72.132.60/tcp/4001/p2p/16Uiu2HAmSwJux2LgfMjQn8CzcG8jufKHbkRm9fQAqEASumU2R38h",
	"/ip4/52.198.239.158/tcp/4001/p2p/16Uiu2HAmQ6YjTL2LCxiYjpRXuukchFjDwm5p7HRjj6nLGcrjwEsL",
	"/ip4/35.73.84.3/tcp/4001/p2p/16Uiu2HAmJ4LQKqAthsP2qcioxMhw2SHSLy21NJTRMGnXcydBufS9",
	"/ip4/13.114.52.45/tcp/4001/p2p/16Uiu2HAkvC1xFzgrZzjjDLmXXuZALVV8HBR8qXJhvEvruJnBmBBm",
	"/ip4/15.184.102.203/tcp/4001/p2p/16Uiu2HAkuiXFRL281t7hqUf7y2MiacfmSiSbnsT5ocj7wZ6pX3Za",
	"/ip4/15.184.3.137/tcp/4001/p2p/16Uiu2HAmJuQmX7oU6hLidy9sxCrWtqjFmiLvuJ7Pw1y3JeJ2k7u1",
	"/ip4/157.175.35.148/tcp/4001/p2p/16Uiu2HAm85nCJNLQ7MLzdCCsZM2im93voZWVp27qanSXUpmPucus",
	"/ip4/16.24.43.94/tcp/4001/p2p/16Uiu2HAm6uAmFB643jWWqUVutHH1hmAKkrycP15hjEte7qvdspTA",
	"/ip4/13.57.144.203/tcp/4001/p2p/16Uiu2HAm5SdP8So2MpknrEAi6avMVXokEuujynKkXRPnxaTxtn5W",
	"/ip4/54.176.58.20/tcp/4001/p2p/16Uiu2HAmPLc4QAuN6gmSUZKYWF1Kh22TFxTANFQYrKHmhh31hikk",
	"/ip4/54.193.237.237/tcp/4001/p2p/16Uiu2HAm814QDPiRztAcyD21NY3ZbeLgKfVWJccFhZ47iDAER9mk",
	"/ip4/54.67.56.29/tcp/4001/p2p/16Uiu2HAmPXJWseUiwwXqjSJX3FkTKcaxb5ixCDDqxsrPCKpCyzLy",
	"/ip4/3.130.97.111/tcp/4001/p2p/16Uiu2HAm5Kd8T7GphFK2kUwfQXEPa6nKwYtkzZKr2R2VoB3PBhNF",
	"/ip4/3.131.15.203/tcp/4001/p2p/16Uiu2HAm4afRu4ny2rcseF34JYNut1TLoenfufAR8SDiDHPZz6Zz",
	"/ip4/3.133.9.176/tcp/4001/p2p/16Uiu2HAmGXv3CDUK8FpV1rBCpGtGsV3rxGshcgBiHTDxYHrbic2Q",
	"/ip4/3.17.210.15/tcp/4001/p2p/16Uiu2HAmSiBrUTxeYj9FXxeHYcboTz8NMyPgDkfjB5yAag3opKur",
	"/ip4/13.40.252.137/tcp/4001/p2p/16Uiu2HAmKM8d2mZ4yNHJ1GZj3FDVD1BXE8Qg7TUBMexkbMGWWbnR",
	"/ip4/13.43.125.230/tcp/4001/p2p/16Uiu2HAkz6Uz3fv2vHPEmn5He89emRZmiToJg7fZaG7MwUob7qr1",
	"/ip4/18.135.186.218/tcp/4001/p2p/16Uiu2HAmBvfpwrE3YBv27kS3E4Zs7wWHj5mni11hScs6QHRdDX8J",
	"/ip4/18.135.47.215/tcp/4001/p2p/16Uiu2HAkvMReWtg6dZ67CyWViYTBe3Zm5hF5LmfstDquK72nfuXq",
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
