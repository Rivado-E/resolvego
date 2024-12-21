package main

import (
	"log"
	"net"
	"os"

	lib "github.com/rivado-e/resolego/lib"
)

func dnsServer() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6969")
	if err != nil {
		log.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Println("Failed to bind to address:", err)
		return
	}

	log.Println("DNS server started at ", udpAddr.String())
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error receiving data:", err)
			break
		}

		message := buf[:size]
		response := []byte{}
			
		// TODO: implement message compression and decompression 
		// see this rfc ref (https://www.rfc-editor.org/rfc/rfc1035#section-4.1.4)
		if receivedHeader, receivedQuestions, _, err := lib.ParseDNSMessage(message); err != nil {
			log.Fatal(err) 
		} else {
			headerFlags := lib.DecodeDNSFlags(receivedHeader.Flags)
			headerFlags.QR = 1
			headerFlags.AA = 0
			headerFlags.TC = 0
			headerFlags.RA = 0
			headerFlags.Z = 0
			headerFlags.RCODE = 4

			receivedHeader.ANCount = receivedHeader.QDCount
			receivedHeader.Flags = lib.EncodeDNSFlags(headerFlags)
			answers := []lib.DNSRecord{}

			for i, question := range receivedQuestions {

				receivedQuestions[i].QType = 1
				receivedQuestions[i].QClass = 1

				for _, arg := range os.Args {
					log.Println("Command line args", arg)
				}

				ips, _ := lib.Forwad(question.QName, os.Args[2])

				for _, ip := range ips {
					log.Println(ip)
				}
				log.Println("Logged it before filtering")

				ips = lib.FilterIpV4(ips)

				data := []byte{}
				if len(ips) != 0 {
					data, err = lib.IPAddressStringToBytes(ips[0])
					if err != nil {
						log.Println("This is the ip: ", ips[0])
						log.Fatal("Error trying to convert ip to bytes: ", err, ips[0])
					}
				}

				record := lib.DNSRecord{
					Name:     question.QName,
					Type:     receivedQuestions[i].QType,
					Class:    receivedQuestions[i].QType,
					TTL:      60,
					RDLength: uint16(len(data)),
					RData: data,
				}
				answers = append(answers, record)
			}

			response = lib.EncodeDNSMessage(receivedHeader, receivedQuestions, answers)
		}

		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			log.Println("Failed to send response: ", err)
		}
	}
}

func main() {
	dnsServer()
}

