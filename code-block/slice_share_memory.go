package code_block

import (
	"bytes"
	"fmt"
)

func ShareMemory() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	// 会造成 dir1 和 dir2 内存共享， dir1 append 右界共享 dir2 左界
	//dir1 := path[:sepIndex]
	// FullSliceExpression   y := x[start, end, capacity]
	// 限制数组的右界， 防止 append 内存共享
	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB
	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB
}
