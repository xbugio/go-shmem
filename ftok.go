package shmem

import (
	"os"
	"syscall"
)

//Ftok C版本ftok的go实现，根据一个已存在文件和一个int值生成另一个int值作为id
func Ftok(pathname string, projID int) (int, error) {
	info, err := os.Stat(pathname)
	if err != nil {
		return 0, err
	}

	stat := info.Sys().(*syscall.Stat_t)
	// 8bit
	a := uint8(projID)
	// 8bit
	b := uint8(stat.Dev)
	// 16bit
	c := uint16(stat.Ino)

	return int(a)<<24 + int(b)<<16 + int(c), nil
}
