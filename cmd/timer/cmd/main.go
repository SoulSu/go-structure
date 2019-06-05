package main

import (
	"log"
	"math/rand"
	"time"
	"timer"
)

const (
	// 时间轮的长度
	wheelSlot = 60
	// 时间轮执行间隔
	wheelTick = time.Minute
)

var task = make([]*timer.TaskLink, 0, wheelSlot)
var pos = int(0)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for i := 1; i <= wheelSlot; i++ {
		task = append(task, timer.NewTaskLink())
	}
	go addTask()
	loop()
}

func addTask() {
	for {
		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
		exAt := time.Now().AddDate(0, 0, rand.Intn(10))
		i := (pos + int(exAt.Unix()%wheelSlot)) % wheelSlot
		task[i].AddTask(timer.NewTask(exAt))
	}
}

func loop() {
	sleep := wheelTick / 180
	for {
		time.Sleep(sleep)
		log.Println(pos, task[pos%wheelSlot])
		pos++
		exAt := time.Now().AddDate(0, 0, rand.Intn(10))
		task[pos%wheelSlot].PopAt(exAt)
	}
}
