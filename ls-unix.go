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
	if err != nil && strings.Contains(err.Error(), "unknown userid") {
		owner = new(user.User)
		owner.Username = uid
	} else {
		check(err)
	}
	group, err := user.LookupGroupId(gid)
	if err != nil && strings.Contains(err.Error(), "unknown groupid") {
		group = new(user.Group)
		group.Name = gid
	} else {
		check(err)
	}
	return owner.Username, group.Name
}
