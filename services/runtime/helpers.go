package runtime

import (
	"crypto/md5"
	"fmt"
	"io"
)

const (
	vethNameLen = 5
)

func getName(id string) string {
	h := md5.New()
	io.WriteString(h, id)
	res := fmt.Sprintf("%x", h.Sum(nil))
	return res[0:vethNameLen]
}
