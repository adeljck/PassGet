package utils

import (
	"os"
	"os/user"
	"path/filepath"
)

func CheckIsInSlice(Items []string, Item string) bool {
	for _, item := range Items {
		if item == Item {
			return true
		}
	}
	return false
}
func GetSliceLastOne(Items []string) (Item string) {
	return Items[len(Items)-1]
}
func ListFilesAndDirs(directory string) ([]string, []string, error) {
	var files []string
	var dirs []string

	// 使用 filepath.Walk 遍历目录
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 忽略根目录本身
		if path == directory {
			return nil
		}

		// 判断是文件还是文件夹
		if info.IsDir() {
			dirs = append(dirs, path) // 如果是目录，添加到 dirs 列表
		} else {
			files = append(files, path) // 如果是文件，添加到 files 列表
		}
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return files, dirs, nil
}
func CheckIsAdmin() bool {
	_, err := user.Current()
	return err == nil
}
