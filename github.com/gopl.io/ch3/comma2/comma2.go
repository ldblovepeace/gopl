package comma2

import (
	"fmt"
	"bytes"
)

func Comma(str string) string{
	var buf bytes.Buffer
	for i, s := range str{
		if i%3 == 0 && i!=0{
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%c", s)
	}
	return buf.String()
}