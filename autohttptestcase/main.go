// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

// This binary provides sample code for using the gopacket TCP assembler and TCP
// stream reader.  It reads packets off the wire and reconstructs HTTP requests
// it sees, logging them.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"encoding/json"
	"github.com/google/gopacket"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"github.com/tidwall/pretty"
	"io/ioutil"
	"strings"
)

var iface = flag.String("i", `\Device\NPF_{E894FE23-2ED6-4941-B639-D796B08FEC3D}`, "Interface to get packets from")
var fname = flag.String("r", "", "Filename to read from, overrides -i")
var snaplen = flag.Int("s", 1600, "SnapLen for pcap packet capture")
var filter = flag.String("f", ":80", "BPF filter for pcap")
var logAllPackets = flag.Bool("v", false, "Logs every packet in great detail")

// Build a simple HTTP request parser using tcpassembly.StreamFactory and tcpassembly.Stream interfaces

// httpStreamFactory implements tcpassembly.StreamFactory
type httpStreamFactory struct{}

// httpStream will handle the actual decoding of http requests.
type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hstream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}

	src, dst := transport.Endpoints()
	if fmt.Sprintf("%v", src) == g_strserverport {
		go hstream.runResponse() // Important... we must guarantee that data from the reader stream is read.
	} else if fmt.Sprintf("%v", dst) == g_strserverport {
		go hstream.runRequest() // Important... we must guarantee that data from the reader stream is read.
	} else {
		log.Panic(net, transport, "not good")
	}
	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &hstream.r
}

func (h *httpStream) runResponse() {
	buf := bufio.NewReader(&h.r)
	defer tcpreader.DiscardBytesToEOF(buf)
	for {
		resp, err := http.ReadResponse(buf, nil)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			// We must read until we see an EOF... very important!
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
			return
		} else {
			bdiscard := true
			if bdiscard {
				//bodyBytes :=
				tcpreader.DiscardBytesToEOF(resp.Body)
			} else {
				//s, _ :=
				ioutil.ReadAll(resp.Body)
				// add this to resp
			}
			//
			resp.Body.Close()
			printResponse(resp, h, 0)
			// log.Println("Received response from stream", h.net, h.transport, ":", resp, "with", bodyBytes, "bytes in response body")
		}
	}
}

var g_discardurllist []string = []string{".html", ".js", ".css", ".png", ".jpg", ".jpeg", ".bmp", ".ico", "htm"}

func DiscardUrl(strurl string) bool {
	for _, val := range g_discardurllist {
		if strings.HasSuffix(strurl, val) {
			return true
		}
	}
	return false
}

func (h *httpStream) runRequest() {

	buf := bufio.NewReader(&h.r)
	defer tcpreader.DiscardBytesToEOF(buf)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			// We must read until we see an EOF... very important!
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
		} else {
			if DiscardUrl(req.RequestURI) {
				tcpreader.DiscardBytesToEOF(req.Body)
				//tell response to discard
			} else {
				ioutil.ReadAll(req.Body)
				// added request to the map
			}

			req.Body.Close()
			printRequest(req, h, 0)
			// log.Println("Received request from stream", h.net, h.transport, ":", req, "with", bodyBytes, "bytes in request body")
		}
	}
}

func printHeader(h http.Header) {
	for k, v := range h {
		fmt.Println(k, v)
	}
}

func printRequest(req *http.Request, h *httpStream, bodyBytes int) {

	fmt.Println("\n\r\n\r")
	fmt.Println("request:", h.net, h.transport)
	fmt.Println("\n\r")
	fmt.Println(req.Method, req.RequestURI, req.Proto)
	printHeader(req.Header)

}

func printResponse(resp *http.Response, h *httpStream, bodyBytes int) {
	fmt.Println("\n\r\n\r datalen: ", bodyBytes)
	fmt.Println("response:", h.net, h.transport)
	fmt.Println("\n\r")
	fmt.Println(resp.Proto, resp.Status)
	printHeader(resp.Header)
}

var g_strserverport string

func main() {
	fmt.Println(string(pretty.Pretty([]byte(` {"response": {"status": null,"method": null,"body": null}}`))))

	data, err2 := json.MarshalIndent(NewTestCase("testcase01", CaseDetail{}), "", "    ")
	fmt.Println(err2)
	fmt.Println(string(data))
	return
	defer util.Run()()
	var handle *pcap.Handle
	var err error

	// Set up pcap packet capture
	if *fname != "" {
		log.Printf("Reading from pcap dump %q", *fname)
		handle, err = pcap.OpenOffline(*fname)
	} else {
		log.Printf("Starting capture on interface %q", *iface)
		handle, err = pcap.OpenLive(*iface, int32(*snaplen), true, pcap.BlockForever)
	}
	if err != nil {
		log.Fatal(err)
	}

	ipportlist := strings.Split(*filter, ":")
	g_strserverport = ipportlist[1]
	var strfilter = ""
	if len(ipportlist[0]) > 0 {
		strfilter = fmt.Sprintf("host %s tcp and port %s", ipportlist[0], ipportlist[1])
	} else {
		strfilter = fmt.Sprintf("tcp and port %s", ipportlist[1])
	}

	if err := handle.SetBPFFilter(strfilter); err != nil {
		log.Fatal(err)
	}

	// Set up assembly
	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	log.Println("reading in packets")
	// Read in packets, pass to assembler.
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	ticker := time.Tick(time.Minute)
	for {
		select {
		case packet := <-packets:
			// A nil packet indicates the end of a pcap file.
			if packet == nil {
				return
			}
			if *logAllPackets {
				log.Println(packet)
			}
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				log.Println("Unusable packet")
				continue
			}
			tcp := packet.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)
		case <-ticker:
			// Every minute, flush connections that haven't seen activity in the past 2 minutes.
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		}
	}
}

//http://httpd.apache.org/
