package main

import (
    "fmt"
    "log"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)

func main() {
    device := "eth0" // change this to your network interface name
    snapshotLen := int32(1024)
    promiscuous := false
    timeout := pcap.BlockForever

    handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
    if err != nil {
        log.Fatal(err)
    }
    defer handle.Close()

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

    fmt.Println("Starting packet capture on", device)

    for packet := range packetSource.Packets() {
        fmt.Println(packet)
    }
}
