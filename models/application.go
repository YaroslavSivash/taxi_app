package models

import (
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// объект одной заявки со счётчиком больше одной заявкиб и имя заявки
type Application struct {
	Name  string
	Count int
}

// содержит информацию об активных и отменнёных заявках в системе
type Applications struct {
	mu                 *sync.Mutex
	activeApplications []*Application
	allApplications    []*Application
}

func NewApplications() *Applications {
	m := &Applications{mu: new(sync.Mutex)}

	for i := 0; i < 50; i++ {
		m.activeApplications = append(m.activeApplications, &Application{Name: generateApplicationName()})
	}
	go m.updateApplications() //запуск горутины которая обновляет пул заявок каждые 200 мсек
	return m
}

func (a *Applications) updateApplications() {
	for {
		time.Sleep(200 * time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 49
		index := rand.Intn(max-min+1) + min
		a.mu.Lock()
		a.allApplications = append(a.allApplications, a.activeApplications[index])
		a.activeApplications[index] = &Application{Name: generateApplicationName()}
		a.mu.Unlock()
	}
}
func generateApplicationName() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz")
	length := 2
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // Например "ExcbsVQs"
	return str
}

func (a *Applications) GetApp() string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 49
	index := rand.Intn(max-min+1) + min
	log.Println(index)
	a.mu.Lock()
	a.activeApplications[index].Count++
	str := a.activeApplications[index].Name
	a.mu.Unlock()
	return str
}

func (a *Applications) GetAllApps() []*Application {
	res := []*Application{}
	a.mu.Lock()
	for _, value := range a.activeApplications {
		if value.Count > 0 {
			res = append(res, value)
		}
	}
	for _, value := range a.allApplications {
		if value.Count > 0 {
			res = append(res, value)
		}
	}
	a.mu.Unlock()
	return res
}
