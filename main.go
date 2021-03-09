package main

import (
	"bytes"
	"compress/flate"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"log"
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
	//http.Handle("/metrics", promhttp.Handler())
	//err := http.ListenAndServe(":2112", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
func connectOkex() {

	u := url.URL{
		Scheme: "wss",
		//Scheme: "ws",
		//Host:   "okexcomrealtest.bafang.com:8443",
		//Host:   "okexcomreal.bafang.com:8443",
		//Host:   "okcoin-push-service.test-a-com.svc.test.local:10442",
		Host: "real.okex.com:8443",
		//Host: "awspush.okex.com:8443",
		//Host: "www.baidu.com:443",
		//Host: "real.coinall.ltd:8443",
		//Host: "real.okcoin.com:8443",
		//Host: "okcoincomreal.bafang.com:8443",
		//Host:   "127.0.0.1:10442",

		Path: "ws/v3",
	}
	log.Printf("connecting to %s", u.String())
	websocket.DefaultDialer.EnableCompression = true
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	log.Printf("connected: %s", u.String())
	//c.WriteMessage(websocket.TextMessage, []byte{0,0,127,127,255,255,0,255})
	//c.WriteMessage(websocket.TextMessage, []byte("hello"))
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["_debug/info"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"jwt","args":"eyJhbGciOiJIUzUxMiJ9.eyJqdGkiOiJvazExMDE2MDk4MzU2NDE4NzE4MjUxNDAzODUyQzdBRTM0T2hzQiIsInVpZCI6Im1qTHZvalVoUHBZTUhGOUhldmw0dmc9PSIsInN0YSI6MCwibWlkIjoibWpMdm9qVWhQcFlNSEY5SGV2bDR2Zz09IiwiaWF0IjoxNjA5ODM1NjQxLCJleHAiOjE2MTA0NDA0NDEsImJpZCI6MCwiZG9tIjoid3d3Lm9rY29pbi5jb20iLCJpc3MiOiJva2NvaW4iLCJzdWIiOiJDMUFEMzBFNEZDQkU3RDE4RjBFQTVDM0ExMUM3NjA1MCJ9.uOs8xAsYobeR1l9TiCeGHxdMsh6WKcWIFsRrPOf4igXgriaREYcXrQk5KobW8qcdu62kSydNkDpyiuS4YDmyNw"}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["swap/mark_price:BTC-USDT-SWAP","index/ticker:BTC-USDT","swap/optimized_depth:BTC-USDT-SWAP","swap/price_range:BTC-USDT-SWAP","swap/funding_rate:BTC-USDT-SWAP","perpetual/funding_fee_task:BTC-USDT-SWAP","swap/position:BTC-USDT-SWAP","swap/account:BTC-USDT-SWAP","swap/order:BTC-USDT-SWAP","swap/ticker:BTC-USDT-SWAP","swap/ticker:LTC-USDT-SWAP","swap/ticker:ETH-USDT-SWAP","swap/ticker:ETC-USDT-SWAP","swap/ticker:XRP-USDT-SWAP","swap/ticker:EOS-USDT-SWAP","swap/ticker:BCH-USDT-SWAP","swap/ticker:BSV-USDT-SWAP","swap/ticker:TRX-USDT-SWAP","swap/ticker:ADA-USDT-SWAP","swap/ticker:ALGO-USDT-SWAP","swap/ticker:ATOM-USDT-SWAP","swap/ticker:COMP-USDT-SWAP","swap/ticker:DASH-USDT-SWAP","swap/ticker:DOGE-USDT-SWAP","swap/ticker:DOT-USDT-SWAP","swap/ticker:FIL-USDT-SWAP","swap/ticker:IOST-USDT-SWAP","swap/ticker:IOTA-USDT-SWAP","swap/ticker:LINK-USDT-SWAP","swap/ticker:KNC-USDT-SWAP","swap/ticker:NEO-USDT-SWAP","swap/ticker:ONT-USDT-SWAP","swap/ticker:QTUM-USDT-SWAP","swap/ticker:THETA-USDT-SWAP","swap/ticker:XLM-USDT-SWAP","swap/ticker:XMR-USDT-SWAP","swap/ticker:XTZ-USDT-SWAP","swap/ticker:ZEC-USDT-SWAP","swap/candle60s:BTC-USDT-SWAP","index/candle60s:BTC-USDT"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["spot/depth:TESTA-TESTB"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["spot/optimized_depth:BTC-USD"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["spot/depth:TRX-USDT","spot/depth:BTC-USDT"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["swap/dep:*"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["spot/ticker:BTC-USDT"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["option/optimized_summary:BTC-USD-210121"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["option/optimized_depth:BTC-USD-210121-35000-C"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["option/trades:BTC-USD"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth:BTC-USD-210326"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["spot/order:BTC-USDT"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth5:BTC-USD-200327"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth5:BTC-USD-200626"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["futures/depth:BTC-USD-200327"]}`))
	c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["option/depth5:BTC-USD-210301-45000-C"]}`))
	//c.WriteMessage(websocket.TextMessage, []byte(`{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}{"op":"subscribe","args":["option/oomunit:TBTC-USD-200103-4000-C","option/candle60s:TBTC-USD-200103-4000-C","option/optimized_depth:TBTC-USD-200103-4000-C"]}`))

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			//go connectOkex()
			return
		}
		switch messageType {
		case websocket.TextMessage:
			// no need uncompressed
			log.Printf("recv uncompress: %s", message)
		case websocket.BinaryMessage:
			// uncompressed
			//log.Println("lenth: %s", len(message))
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
					//if sub.Milliseconds() > 10 {
					log.Printf("%d", sub.Milliseconds())
					//}
					promRecorder.Observe(sub.Seconds())

				}
			}
		case websocket.CloseMessage:
			c.Close()
			//go connectOkex()
		}

	}

}

func GzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()

	return ioutil.ReadAll(reader)

}
