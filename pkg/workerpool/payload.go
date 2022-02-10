package workerpool

import (
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

// 真正干活的函数，占用内存，耗费时间
func (p *Payload) UploadToS3() error {
	// log.Printf("uploading %v\n", p.N)
	time.Sleep(2 * time.Second)
	// fmt.Printf("uploaded %v\n", p.N)
	return nil
}
