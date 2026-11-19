//go:build unix

package acfs

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

// fileIdentity returns the owner and device and inode numbers of info. Dev's
// type varies by platform, so it is widened through int64.
func fileIdentity(info os.FileInfo) (uid, gid int, device, inode uint64) {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return 0, 0, 0, 0
	}
	return int(stat.Uid), int(stat.Gid), uint64(int64(stat.Dev)), stat.Ino
}

func preserveFileOwnership(target *os.File, original os.FileInfo) error {
	stat, ok := original.Sys().(*syscall.Stat_t)
	if !ok {
		return errors.New("original file ownership is unavailable")
	}
	current, err := target.Stat()
	if err != nil {
		return fmt.Errorf("inspect staged file ownership: %w", err)
	}
	uid, gid, _, _ := fileIdentity(current)
	if uid == int(stat.Uid) && gid == int(stat.Gid) {
		return nil
	}
	if err = target.Chown(int(stat.Uid), int(stat.Gid)); err != nil {
		return fmt.Errorf("preserve original file ownership: %w", err)
	}
	return nil
}
