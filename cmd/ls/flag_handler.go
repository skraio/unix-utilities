package ls

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"sort"
	"strconv"
	"syscall"
)

func longFormat(file fs.FileInfo) (FileAttributes, error) {
	fileMode := file.Mode()

	var nlink string
	if stat, ok := file.Sys().(*syscall.Stat_t); ok {
		nlink = strconv.Itoa(int(stat.Nlink))
	}

	var uid string
	if stat, ok := file.Sys().(*syscall.Stat_t); ok {
		id := stat.Uid

		u, err := user.LookupId(fmt.Sprintf("%d", id))
		if err != nil {
            return FileAttributes{}, err
		}

		uid = u.Username
	}

	var fileSize string
	if pFlags.Readable {
		fileSize = humanReadableSize(file.Size())
	} else {
		fileSize = strconv.FormatInt(file.Size(), 10)
	}

	modTime := file.ModTime().Format("Jan _2 15:04")

	return FileAttributes{fileMode, nlink, uid, fileSize, modTime}, nil
}

func humanReadableSize(size int64) string {
	var fileSize string
	if size >= 1024*1024 {
		fileSize = strconv.FormatInt(size/1024/1024, 10) + "M"
	} else {
		fileSize = strconv.FormatInt(size, 10)
	}
	return fileSize
}

func sortByModTime(content []os.FileInfo) {
	sort.Slice(content, func(i, j int) bool {
		return content[j].ModTime().Before(content[i].ModTime())
	})
}

func reverseOrder(content []os.FileInfo) {
	sort.Slice(content, func(i, j int) bool {
		return i > j
	})
}
