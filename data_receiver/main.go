package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UpLiftL1f3/tollCalc/types"
	"github.com/gorilla/websocket"
)

func main() {
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
	fmt.Println("data receiver working fine")
}

type DataReceiver struct {
	msgCh chan types.OBUData
	conn  *websocket.Conn
	prod  DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	var (
		p          DataProducer
		err        error
		kafkaTopic = "obuData"
	)

	p, err = NewKafkaProducer(kafkaTopic)
	if err != nil {
		log.Fatal(err)
	}
	p = NewLogMiddleware(p)

	return &DataReceiver{
		msgCh: make(chan types.OBUData, 128),
		prod:  p,
	}, nil
}

func (dr *DataReceiver) produceData(data types.OBUData) error {
	return dr.prod.ProduceData(data)
}

// -> Web Socket
func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	dr.conn = conn

	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("new OBU client connected!")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error: ", err)
			continue
		}

		// fmt.Println("received message", data)
		// fmt.Printf("received OBU data from [%d] :: <lat %.2f, long %.2f> \n", data.OBUID, data.Lat, data.Long)
		if err := dr.produceData(data); err != nil {
			log.Println("kafka produce errror: ", err)
		}

		//! once channel is full it will end the for loop and stop Receiving
		// dr.msgCh <- data
	}
}
