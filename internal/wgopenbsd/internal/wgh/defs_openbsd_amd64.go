//+build openbsd,amd64

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs.go

package wgh

const (
	SIOCGIFGMEMB = 0xc028698a

	SizeofIfgreq = 0x10
)

type Ifgroupreq struct {
	Name   [16]byte
	Len    uint32
	Pad1   [4]byte
	Groups *Ifgreq
	Pad2   [8]byte
}

type Ifgreq struct {
	Ifgrqu [16]byte
}

type Timespec struct {
	Sec  int64
	Nsec int64
}

type WGDataIO struct {
	Name [16]byte
	Size uint64
	Mem  *byte
}

type WGInterfaceIO struct {
	Flags   uint8
	Peers   *WGPeerIO
	Port    uint16
	Rtable  int32
	Public  [32]byte
	Private [32]byte
}

type WGPeerIO struct {
	Flags            int32
	Next             *WGPeerIO
	Aips             *WGAIPIO
	Protocol_version int32
	Public           [32]byte
	Psk              [32]byte
	Pka              uint16
	Pad_cgo_0        [2]byte
	Endpoint         [28]byte
	Txbytes          uint64
	Rxbytes          uint64
	Last_handshake   Timespec
}

type WGAIPIO struct {
	Flags int32
	Next  *WGAIPIO
	Data  WGAIPData
}

type WGAIPData struct {
	Af   uint8
	Cidr int32
	Addr [16]byte
}

const (
	SIOCGWG = 0xc02069d3

	WG_INTERFACE_HAS_PUBLIC    = 0x1
	WG_INTERFACE_HAS_PRIVATE   = 0x2
	WG_INTERFACE_HAS_PORT      = 0x4
	WG_INTERFACE_HAS_RTABLE    = 0x8
	WG_INTERFACE_REPLACE_PEERS = 0x10

	SizeofWGInterfaceIO = 0x58
)
