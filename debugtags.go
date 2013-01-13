package debugtags

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

var spaces = "                                                                  "

type Tracer struct {
	Enabled bool
	level   int
}

func (t *Tracer) In(name string, items ...interface{}) {
	if t.Enabled {
		fmt.Printf("%s+%s %v\n", spaces[:t.level], name, items)
	}
	t.level++
}

func (t *Tracer) Out(name string, items ...interface{}) {
	if e := recover(); e != nil {
		panic(e)
	}
	t.level--
	if t.Enabled {
		fmt.Printf("%s-%s %v\n", spaces[:t.level], name, items)
	}
}

func (t *Tracer) Println(items ...interface{}) {
	if t.Enabled {
		fmt.Print(spaces[:t.level])
		fmt.Println(items...)
	}
}
func (t *Tracer) Printf(format string, items ...interface{}) {
	if t.Enabled {
		str := fmt.Sprintf(format, items...)
		str = strings.TrimSpace(str)
		fmt.Print(spaces[:t.level])
		fmt.Println(str)
	}
}

func (t *Tracer) JSON(item interface{}) {
	if t.Enabled {
		var b1 bytes.Buffer
		enc := json.NewEncoder(&b1)
		enc.Encode(item)
		var b2 bytes.Buffer
		json.Indent(&b2, b1.Bytes(), spaces[:t.level], " ")
		fmt.Println(b2.String())
	}
}
