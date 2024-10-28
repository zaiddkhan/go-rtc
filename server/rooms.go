package server

import (
	"log"
	"math/rand"
	"sync"
	"time"
	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn

}

type RoomMap struct {
	Mutex sync.RWMutex
	Map map[string][]Participant
}

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

func (r *RoomMap ) Get(roomId string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	return r.Map[roomId]
}

func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune,9)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	r.Map[roomID] = []Participant{}
	return roomID
}

func (r *RoomMap) InsertIntoRoom(roomId string,host bool,conn *websocket.Conn){
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	p := Participant{
		Host: host,
		Conn: conn,
	}
	log.Print("Inserting into room")
	r.Map[roomId] = append(r.Map[roomId], p)
}


func (r *RoomMap) DeleteRoom(roomId string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	delete(r.Map,roomId)
}