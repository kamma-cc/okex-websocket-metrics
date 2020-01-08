package main

import (
	"bytes"
	"compress/flate"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var (
	promRecorder = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tbt_durations_histogram_seconds",
		Help:    "RPC latency distributions.",
		Buckets: prometheus.LinearBuckets(0, 0.001, 50),
	})
)

func main() {
	prometheus.MustRegister(promRecorder)
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
		//Scheme: "ws",
		//Host:   "okexcomreal.bafang.com:8443",
		Host: "okexcomrealtest.bafang.com:10442",
		//Host:   "real.okex.com:8443",
		//Host:   "127.0.0.1:10442",
		Path: "ws/v3",
	}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	log.Printf("connected: %s", u.String())
	//c.WriteMessage(websocket.TextMessage, []byte{0,0,127,127,255,255,0,255})
	//c.WriteMessage(websocket.TextMessage, []byte("hello"))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"jwt","args":"eyJhbGciOiJIUzUxMiJ9.eyJqdGkiOiJleDExMDE1NzUwOTExMjY5NjQ0OTg0NDQ3Qjg2QjQxQzBDZmN2YyIsInVpZCI6ImNCK2ZtTWs2eWVpdGlhR1dQdzBQR2c9PSIsInN0YSI6MCwibWlkIjowLCJpYXQiOjE1NzUwOTExMjYsImV4cCI6MTU3NTY5NTkyNiwiYmlkIjowLCJkb20iOiJva2V4YmV0YS5iYWZhbmcuY29tIiwiaXNzIjoib2tjb2luIiwic3ViIjoiMzAwMjJGNEE5NzU3N0VFRjE0NUE5QTk4REM5RTk2MEUifQ.99aPWVeuUNr69ipNfoPYPW_2kVjBAtUXQ65eHi0A-dRQF5RUBpsHNFRZMBQjHEuN7CnukiEdM0fFJ3EC7N3WIw"}`))
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["_debug/info"]}`))
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["swap/depth_l2_tbt:EOS-USD-SWAP"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth_l2_tbt:BTC-USD-200327"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["option/depth:TBTC-USD-200117-6000-C"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}`))
	//c.WriteMessage(websocket.PingMessage, []byte(`fwaeawfwef`))

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			go connectOkex()
			return
		}
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
				log.Printf("recv: %s:%s", time.Now().Unix(), text)
				r, _ := regexp.Compile(`"timestamp":"(.{24})"`)
				loc := r.FindStringSubmatch(string(text))
				if len(loc) > 1 {
					parse, _ := time.Parse("2006-01-02T15:04:05.999999999Z", loc[1])
					//log.Println(parse)
					sub := time.Now().Sub(parse)
					if sub.Milliseconds() > 3 {
						log.Printf("%d", sub.Milliseconds())
					}
					promRecorder.Observe(sub.Seconds())

				}
			}
		case websocket.CloseMessage:
			c.Close()
			go connectOkex()
		}

	}

}

func GzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()

	return ioutil.ReadAll(reader)

}
