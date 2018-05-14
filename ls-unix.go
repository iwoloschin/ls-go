// +build !windows

package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"syscall"
)

func getOwnerAndGroup(fileInfo *os.FileInfo) (string, string) {
	stat_t := (*fileInfo).Sys().(*syscall.Stat_t)
	uid := fmt.Sprint(stat_t.Uid)
	gid := fmt.Sprint(stat_t.Gid)
	owner, err := user.LookupId(uid)
	check(err)
	group, err := user.LookupGroupId(gid)
	if err != nil && strings.Contains(err.Error(), "unknown groupid") {
		return owner.Username, gid
	} else {
		check(err)
	}
	return owner.Username, group.Name
}
