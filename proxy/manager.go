package proxy

import (
	"time"
)

type Proxy struct {
	Value      string
	activateAt time.Time
}

type Manager struct {
	Proxys []*Proxy
	next   uint
}

// Добавить получение прокси из конфигурационного файла
func NewManager() *Manager {
	return &Manager{
		next: 0,
		Proxys: []*Proxy{
			{Value: "1"},
			{Value: "2"},
			{Value: "3"},
			{Value: "4"},
			{Value: "5"},
		},
	}
}

func (m *Manager) Next() {
	m.next++
	if int(m.next) == len(m.Proxys) {
		m.next = 0
	}
}

func (m *Manager) Run() (<-chan *Proxy, chan<- struct{}) {
	// Канал, куда будут отправляться прокси
	ch := make(chan *Proxy, len(m.Proxys))
	// Канал, чтобы завершить работу PM
	done := make(chan struct{})

	go func() {
		defer close(ch)

		for {
			select {
			case <-done:
				return
			default:
				// Получение прокси
				p := m.Proxys[m.next]
				// Переключение прокси
				m.Next()
				// Если время сейчас, меньше чем время активации пропуск
				if time.Now().Before(p.activateAt) {
					continue
				}

				// Можно добавить случайное время
				p.activateAt = time.Now().Add(7 * time.Second)

				// Отправка прокси
				ch <- p
			}
		}
	}()

	return ch, done
}
