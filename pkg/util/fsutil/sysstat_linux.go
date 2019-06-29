package fsutil

import (
	"syscall"
)

func SysStatInfo(raw *syscall.Stat_t) SysStat {
	return SysStat{
		Uid:   raw.Uid,
		Gid:   raw.Gid,
		Atime: raw.Atim,
		Mtime: raw.Mtim,
		Ctime: raw.Ctim,
	}
}
