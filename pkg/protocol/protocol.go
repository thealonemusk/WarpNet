package protocol

import (
	p2pprotocol "github.com/libp2p/go-libp2p/core/protocol"
)

const (
	WarpNet         Protocol = "/WarpNet/0.1"
	ServiceProtocol Protocol = "/WarpNet/service/0.1"
	FileProtocol    Protocol = "/WarpNet/file/0.1"
	EgressProtocol  Protocol = "/WarpNet/egress/0.1"
)

const (
	FilesLedgerKey    = "files"
	MachinesLedgerKey = "machines"
	ServicesLedgerKey = "services"
	UsersLedgerKey    = "users"
	HealthCheckKey    = "healthcheck"
	DNSKey            = "dns"
	EgressService     = "egress"
	TrustZoneKey      = "trustzone"
	TrustZoneAuthKey  = "trustzoneAuth"
)

type Protocol string

func (p Protocol) ID() p2pprotocol.ID {
	return p2pprotocol.ID(string(p))
}
