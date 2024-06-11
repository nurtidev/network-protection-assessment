package interceptor

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	packetCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "packet_count",
			Help: "Total number of packets",
		},
		[]string{"protocol"},
	)
)

func init() {
	prometheus.MustRegister(packetCount)
}

func StartCapture(device string) error {
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("failed to open device: %v", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		processPacket(packet)
	}

	return nil
}

func processPacket(packet gopacket.Packet) {
	fmt.Println("Packet captured:", packet)

	if packet.NetworkLayer() != nil {
		switch packet.NetworkLayer().LayerType() {
		case layers.LayerTypeIPv4:
			packetCount.WithLabelValues("IPv4").Inc()
		case layers.LayerTypeIPv6:
			packetCount.WithLabelValues("IPv6").Inc()
		}
	}

	// Process the packet and extract metrics
}
