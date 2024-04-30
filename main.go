package main

import (
	"fmt"
	"proxy-go/proxy"
	"time"
)

func main() {
	pm := proxy.NewManager()
	proxys, done := pm.Run()

	// Выключить PM через 30 секунд
	go func() {
		time.Sleep(30 * time.Second)
		done <- struct{}{}
	}()

	// Чтение прокси
	// Он будет читать даже после остановки, если там что-то есть
	for v := range proxys {
		fmt.Println(v.Value)
	}
}
