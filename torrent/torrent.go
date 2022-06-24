package torrent

import (
	"bytes"
	"fmt"
)

type Torrent struct {
	Infohash string
	Peers    []Peer
}

type Peer struct {
	Identifier      string
	IpAddress       string
	Port            uint16
	ClaimedTransfer TransferMetrics
	ActualTransfer  TransferMetrics
	Remaining       uint64
	Status          PeerStatus
}

type TransferMetrics struct {
	Upload   uint64
	Download uint64
}

type PeerStatus string

const (
	PeerStatusStarted   = "started"
	PeerStatusCompleted = "completed"
	PeerStatusStopped   = "stopped"
	PeerStatusEmpty     = "empty"
)

// Database of all torrents
var Torrents map[string]*Torrent

func LookupTorrent(infohash string) (*Torrent, bool) {
	torrent, exists := Torrents[infohash]
	return torrent, exists
}

func (torrent *Torrent) CompactEncodedPeers(infohash string) []byte {
	var result bytes.Buffer

	result.WriteString("l")

	for _, peer := range torrent.Peers {
		result.WriteString(
			fmt.Sprintf(
				"d2:ip%d:%s4:porti%dee",
				len(peer.IpAddress),
				peer.IpAddress,
				peer.Port,
			),
		)
	}

	result.WriteString("e")

	return result.Bytes()
}
