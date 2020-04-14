package shmem

import "unsafe"

// CopyToMem 将[]byte数据拷贝到内存地址中
func CopyToMem(dst uintptr, src []byte, n uint) {
	var i uint
	for i = 0; i < n; i++ {
		*(*byte)(unsafe.Pointer(dst + uintptr(i))) = src[i]
	}
}

// CopyFromMem 将内存地址中的数据拷贝到[]byte中
func CopyFromMem(dst []byte, src uintptr, n uint) {
	var i uint
	for i = 0; i < n; i++ {
		dst[i] = *(*byte)(unsafe.Pointer(src + uintptr(i)))
	}
}
