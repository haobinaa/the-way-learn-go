package code_block

import (
	"encoding/binary"
	"io"
	"strings"
)

type Point struct {
	Longitude     string
	Latitude      string
	Distance      string
	ElevationGain string
	ElevationLoss string
}
type Reader struct {
	r   io.Reader
	err error
}

func (r *Reader) read(data interface{}) {
	// 将错误包在结构体内部， 屏蔽过程中的 error
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

// MaskErr 屏蔽过程中的错误， 避免过多 if err != nil 写法
func MaskErr() (*Point, error) {
	var p Point
	r := Reader{r: strings.NewReader("Read will return these bytes")}
	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)
	if r.err != nil {
		return nil, r.err
	}
	return &p, nil
}
