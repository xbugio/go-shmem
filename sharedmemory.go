package shmem

import (
	"unsafe"
)

// SharedMemory 共享内存
type SharedMemory struct {
	shmid int
	size  uint
	addr  uintptr
}

// Open 打开获取一段共享内存
func Open(key int, size uint, mode int, perm int) (*SharedMemory, error) {
	s := &SharedMemory{}

	shmid, err := Shmget(key, size, mode|perm)
	if err != nil {
		return nil, err
	}

	s.shmid = shmid
	s.size = size
	return s, nil
}

// Attach 挂载共享内存
func (s *SharedMemory) Attach(mode int) error {
	addr, err := Shmat(s.shmid, 0, mode)
	if err != nil {
		return err
	}
	s.addr = addr
	return nil
}

// Detach 卸载共享内存
func (s *SharedMemory) Detach() error {
	return Shmdt(s.addr)
}

// Close 标记删除共享内存
func (s *SharedMemory) Close() error {
	return Shmctl(s.shmid, IPC_RMID, nil)
}

// Pointer 获取共享内存操作指针
func (s *SharedMemory) Pointer() unsafe.Pointer {
	return unsafe.Pointer(s.addr)
}

// Addr 获取共享内存地址
func (s *SharedMemory) Addr() uintptr {
	return s.addr
}

// Size 获取共享内存空间大小
func (s *SharedMemory) Size() uint {
	return s.size
}

// Set 设置共享内存数据
func (s *SharedMemory) Set(data []byte) {
	size := uint(len(data))
	if size > s.size {
		size = s.size
	}
	CopyToMem(s.addr, data, size)
}

// Get 获取共享内存数据
func (s *SharedMemory) Get(buff []byte) {
	size := uint(len(buff))
	if size > s.size {
		size = s.size
	}
	CopyFromMem(buff, s.addr, size)
}
