package item

import (
	"strconv"
	"time"
)

func GererateId() string {
	return "todo_" + strconv.FormatInt(time.Now().Unix(), 10)
}
