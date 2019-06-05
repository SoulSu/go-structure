package timer

import (
	"testing"
	"time"
)

func TestTaskLink_AddTask(t *testing.T) {
	tl := NewTaskLink()
	n := time.Now()
	tl.AddTask(&Task{exAt: n.AddDate(0, 0, 1)})
	t.Log(tl.String())
	tl.AddTask(&Task{exAt: n.AddDate(0, 0, 2)})
	t.Log(tl.String())
	tl.AddTask(&Task{exAt: n.AddDate(0, 0, 3)})
	t.Log(tl.String())
	tl.AddTask(&Task{exAt: n.AddDate(0, 0, 4)})
	t.Log(tl.String())
	tl.AddTask(&Task{exAt: n.AddDate(0, 0, 5)})
	t.Log(tl.String())
	tl.AddTask(&Task{exAt: n.AddDate(0, 0, -1)})
	t.Log(tl.String())
	tl.AddTask(&Task{exAt: n.AddDate(0, 0, -9)})
	t.Log(tl.String())
	tl.AddTask(&Task{exAt: n.AddDate(0, 0, -5)})
	t.Log(tl.String())
	//tl.AddTask(&Task{exAt: n.AddDate(0, 0, 9)})
	//t.Log(tl.String())

}
