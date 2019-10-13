package main

import (
	"bytes"
	"compress/flate"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

func main() {
	go runProm()
	go connectOkex()
	select {}
}

func runProm() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func connectOkex() {

	u := url.URL{
		Scheme: "wss",
		Host:   "okexcomreal.bafang.com:8443",
		Path:   "ws/v3",
	}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	log.Printf("connected: %s", u.String())
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth_l2_tbt:BTC-USD-191227"]}`))
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth_l2_tbt:ETH-USD-191227"]}`))
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth_l2_tbt:LTC-USD-191227"]}`))
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth_l2_tbt:EOS-USD-191227"]}`))
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth_l2_tbt:XRP-USD-191227"]}`))
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth_l2_tbt:BCH-USD-191227"]}`))
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth_l2_tbt:TRX-USD-191227"]}`))

	for {
		messageType, message, err := c.ReadMessage()
		switch messageType {
		case websocket.TextMessage:
			// no need uncompressed
			log.Printf("recv uncompress: %s", message)
		case websocket.BinaryMessage:
			// uncompressed
			text, err := GzipDecode(message)
			if err != nil {
				log.Println("err", err)
			} else {
				//log.Printf("recv: %s:%s", time.Now().Unix(), text)
				r, _ := regexp.Compile(`"timestamp":"(.{24})"`)
				loc := r.FindStringSubmatch(string(text))
				if len(loc) > 1 {
					parse, _ := time.Parse("2006-01-02T15:04:05.999999999Z", loc[1])
					//log.Println(parse)
					sub := time.Now().Sub(parse)
					log.Println(sub)
					if sub.Seconds() > 1 {
						log.Println(sub)
						log.Println(time.Now())
						log.Println(string(text))
					}

				}
			}
		}
		if err != nil {
			log.Println("read:", err)
			return
		}
	}

}

func GzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()

	return ioutil.ReadAll(reader)

}
