package timer

import (
	"bytes"
	"time"
)

// 执行任务
type Task struct {
	Next *Task

	// 触发时间
	exAt time.Time
}

func NewTask(exAt time.Time) *Task {
	return &Task{exAt: exAt}
}

// 任务列表
type TaskLink struct {
	head *Task
}

func NewTaskLink() *TaskLink {
	return &TaskLink{}
}

// 添加新任务到链表中
func (tl *TaskLink) AddTask(newTask *Task) {
	if tl.head == nil {
		tl.head = newTask
		return
	}

	if tl.head.exAt.After(newTask.exAt) {
		newTask.Next = tl.head
		tl.head = newTask
		return
	}

	prev, next := tl.head, tl.head
	for next != nil {
		if next.exAt.After(newTask.exAt) {
			break
		}
		prev = next
		next = next.Next
	}
	if next == nil {
		prev.Next = newTask
		return
	}

	newTask.Next = prev.Next
	prev.Next = newTask
}

// 取出 t 时间之前的任务
func (tl *TaskLink) PopAt(t time.Time) []*Task {
	var tasks []*Task
	next := tl.head
	for next != nil {
		if next.exAt.Before(t) || next.exAt.Equal(t) {
			tasks = append(tasks, next)
			next = next.Next
		} else {
			break
		}
	}
	tl.head = next
	return tasks
}

func (tl *TaskLink) String() string {
	bf := bytes.NewBufferString("")
	head := tl.head
	for head != nil {
		bf.WriteString(head.exAt.String())
		bf.WriteString(" ----> ")
		head = head.Next
	}
	return bf.String()
}
