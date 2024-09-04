package closure

import (
	"errors"
	"fmt"
)

// type TestFunctionPointType1 *func() int
// type TestFunctionPointType2 *func(int) int
type Qda struct {
	p1    int
	p2    string
	state bool
}

func Initialization(a int, s string, q *Qda) (err error) {
	if len(s) == 0 {
		err = errors.New("-1")
		defer func() {
			q.state = false
		}()
		return
	}
	if q.state {
		err = errors.New("-1")
		return
	} else {
		q.p1 = a
		q.p2 = s
		defer func() {
			q.state = true
		}()
		err = nil
		return
	}
}
func Display(q *Qda) (err error) {
	if !q.state {
		err = errors.New("-1")
		return
	} else {
		err = nil
		fmt.Printf("value of p1 is %d -- value of p2 is %s\n", q.p1, q.p2)
		return
	}

}

func Reset(q *Qda) (err error) {
	if !q.state {
		err = errors.New("-1")
		return
	} else {
		q.p1 = 0
		q.p2 = ""
		err = nil
		defer func() {
			q.state = false
		}()
		return
	}
}

func Delete(q *Qda) (err error) {
	q = nil
	err = nil
	return
}
