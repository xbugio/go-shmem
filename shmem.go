package shmem

import (
	"syscall"
	"unsafe"
)

// IPCPerm C版本struct ipc_perm的包装
type IPCPerm struct {
	Key  uint32 /* Key supplied to shmget(2) */
	UID  uint32 /* Effective UID of owner */
	GID  uint32 /* Effective GID of owner */
	CUID uint32 /* Effective UID of creator */
	CGID uint32 /* Effective GID of creator */
	Mode uint16 /* Permissions + SHM_DEST and SHM_LOCKED flags */
	Seq  uint16 /* Sequence number */
}

// ShmIDDs C版本struct shmid_ds的包装
type ShmIDDs struct {
	ShmPerm   IPCPerm /* Ownership and permissions */
	ShmSegsz  uint32  /* Size of segment (bytes) */
	ShmAtime  int32   /* Last attach time */
	ShmDtime  int32   /* Last detach time */
	ShmCtime  int32   /* Last change time */
	ShmCpid   int32   /* PID of creator */
	ShmLpid   int32   /* PID of last shmat(2)/shmdt(2) */
	ShmNattch uint32  /* No. of current attaches */
}

const (
	/* common mode bits */
	IPC_R = 000400 /* read permission */
	IPC_W = 000200 /* write/alter permission */
	IPC_M = 010000 /* permission to change control info */

	/* SVID required constants (same values as system 5) */
	IPC_CREAT  = 001000 /* create entry if key does not exist */
	IPC_EXCL   = 002000 /* fail if key exists */
	IPC_NOWAIT = 004000 /* error if request must wait */

	IPC_PRIVATE = 0 /* private key */

	IPC_RMID = 0 /* remove identifier */
	IPC_SET  = 1 /* set options */
	IPC_STAT = 2 /* get options */
)

// Shmget C版本shmget包装，创建/获取共享内存. 返回共享内存id
func Shmget(key int, size uint, shmflg int) (int, error) {
	ret, _, errNo := syscall.Syscall(syscall.SYS_SHMGET, uintptr(key), uintptr(size), uintptr(shmflg))
	return int(ret), errnoErr(errNo)
}

// Shmat C版本shmat包装，将共享内存挂载，使进程可以进行操作. 返回共享内存挂载地址
func Shmat(shmid int, shmaddr uintptr, shmflg int) (uintptr, error) {
	ret, _, errNo := syscall.Syscall(syscall.SYS_SHMAT, uintptr(shmid), shmaddr, uintptr(shmflg))
	return ret, errnoErr(errNo)
}

// Shmdt C版本shmdt包装，将共享内存卸载，进程将不可再进行操作
func Shmdt(shmaddr uintptr) error {
	_, _, errNo := syscall.Syscall(syscall.SYS_SHMDT, shmaddr, 0, 0)
	return errnoErr(errNo)
}

// Shmctl C版本shmctl包装，用于控制管理共享内存，标记删除，修改权限等
func Shmctl(shmid, cmd int, buf *ShmIDDs) error {
	_, _, errNo := syscall.Syscall(syscall.SYS_SHMCTL, uintptr(shmid), uintptr(cmd), uintptr(unsafe.Pointer(buf)))
	return errnoErr(errNo)
}

func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case syscall.EAGAIN:
		return syscall.EAGAIN
	case syscall.EINVAL:
		return syscall.EINVAL
	case syscall.ENOENT:
		return syscall.ENOENT
	}
	return e
}
