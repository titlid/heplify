package protos

import (
	"log"

	"github.com/google/gopacket"
	"github.com/sipcapture/heplify/ownlayers"
)

func NewRTP(raw []byte) string {
	rtpl := gopacket.NewPacket(raw, ownlayers.LayerTypeRTP, gopacket.DecodeOptions{Lazy: true, NoCopy: true})
	rtp, ok := rtpl.Layers()[0].(*ownlayers.RTP)
	if !ok {
		//return nil
		return "this is not a RTP packet!"
	}
	log.Println("rtp payload type", rtp.PayloadType)
	return rtp.String()
}
