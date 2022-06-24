package torrent

import (
	"fmt"
	"strconv"
)

type Announcement struct {
	PeerId     string
	Ip         string
	Port       uint16
	Uploaded   uint64
	Downloaded uint64
	Left       uint64
	Event      PeerStatus
}

func BuildAnnouncement(
	peerId string,
	ip string,
	portStr string,
	uploadedStr string,
	downloadedStr string,
	leftStr string,
	eventStr string,
) (Announcement, error) {

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return Announcement{}, fmt.Errorf("could not parse port")
	}

	uploaded, err := strconv.Atoi(uploadedStr)
	if err != nil {
		return Announcement{}, fmt.Errorf("could not parse uploaded amount")
	}

	downloaded, err := strconv.Atoi(downloadedStr)
	if err != nil {
		return Announcement{}, fmt.Errorf("could not parse downloaded amount")
	}

	left, err := strconv.Atoi(leftStr)
	if err != nil {
		return Announcement{}, fmt.Errorf("could not parse left amount")
	}

	var event PeerStatus

	switch eventStr {
	case "started":
		event = PeerStatusStarted
	case "completed":
		event = PeerStatusCompleted
	case "stopped":
		event = PeerStatusStopped
	case "empty":
		event = PeerStatusEmpty
	case "":
		event = PeerStatusEmpty
	default:
		return Announcement{}, fmt.Errorf("unrecognized event %q", eventStr)
	}

	announcement := Announcement{
		PeerId:     peerId,
		Ip:         ip,
		Port:       uint16(port),
		Uploaded:   uint64(uploaded),
		Downloaded: uint64(downloaded),
		Left:       uint64(left),
		Event:      event,
	}

	return announcement, nil
}
