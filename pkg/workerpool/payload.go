package workerpool

import (
	"log"
	"time"
)

type PayloadCollection struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

type Payload struct {
	N int `json:"n"`
}

func (p *Payload) UploadToS3() error {
	log.Printf("uploading %v\n", p.N)
	time.Sleep(2 * time.Second)
	// fmt.Printf("uploaded %v\n", p.N)
	return nil
}
