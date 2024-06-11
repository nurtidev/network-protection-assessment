package interceptor

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os/exec"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/nurtidev/network-protection-assessment/internal/exporter"
)

func StartCapture(device string) error {
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("Failed to open device: %v", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		startTime := time.Now()
		processPacket(packet)
		duration := time.Since(startTime).Seconds()
		recordMetrics(packet, duration)
	}

	return nil
}

func processPacket(packet gopacket.Packet) {
	fmt.Println("Packet captured:", packet)

	if packet.NetworkLayer() != nil {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			// Проверка подозрительной активности
			if isSuspicious(ip) {
				logSuspiciousActivity(ip)
				blockIP(ip.SrcIP.String())
			}
		}
	}

	// Случайная генерация подозрительной активности
	if rand.Float32() < 0.1 { // 10% вероятность
		srcIP := packet.NetworkLayer().NetworkFlow().Src().String()
		logSuspiciousActivity(&layers.IPv4{SrcIP: net.ParseIP(srcIP)})
	}
}

func isSuspicious(ip *layers.IPv4) bool {
	// Пример простого правила обнаружения: проверка черного списка IP
	blacklistedIPs := []string{"192.168.1.100", "10.0.0.1"}
	for _, blkIP := range blacklistedIPs {
		if ip.SrcIP.String() == blkIP {
			return true
		}
	}
	// Дополнительные правила обнаружения
	return false
}

func logSuspiciousActivity(ip *layers.IPv4) {
	if ip == nil || ip.SrcIP == nil {
		return
	}
	fmt.Printf("Suspicious activity detected from IP: %s\n", ip.SrcIP)
	// Логирование события и экспорт в Prometheus
	exporter.SuspiciousActivity.WithLabelValues(ip.SrcIP.String()).Inc()
}

func blockIP(ip string) {
	if ip == "" {
		return
	}
	// Пример простого блока IP: добавление правила фаервола
	cmd := exec.Command("iptables", "-A", "INPUT", "-s", ip, "-j", "DROP")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to block IP: %v", err)
	}
	fmt.Printf("Blocked IP: %s\n", ip)
}

func recordMetrics(packet gopacket.Packet, duration float64) {
	if packet.NetworkLayer() != nil {
		switch packet.NetworkLayer().LayerType() {
		case layers.LayerTypeIPv4:
			exporter.PacketCount.WithLabelValues("IPv4").Inc()
			exporter.DataVolume.WithLabelValues("IPv4").Add(float64(len(packet.Data())))
			exporter.PacketDuration.WithLabelValues("IPv4").Observe(duration)
		case layers.LayerTypeIPv6:
			exporter.PacketCount.WithLabelValues("IPv6").Inc()
			exporter.DataVolume.WithLabelValues("IPv6").Add(float64(len(packet.Data())))
			exporter.PacketDuration.WithLabelValues("IPv6").Observe(duration)
		}
	}
}
