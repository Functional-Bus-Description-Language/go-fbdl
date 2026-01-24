package util

import (
	"os"
	"syscall"
)

// Directory ID
type DirID struct {
	Dev uint64 // Device ID, identifies directory filesystem device
	Ino uint64 // Inode number
}

func GetDirID(dirPath string) (DirID, error) {
	info, err := os.Stat(dirPath)
	if err != nil {
		return DirID{0, 0}, err
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return DirID{0, 0}, err
	}

	return DirID{stat.Dev, stat.Ino}, nil
}
