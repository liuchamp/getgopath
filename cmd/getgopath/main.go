package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

// getGoModulePath 获取当前模块路径
func getGoModulePath(dir string) (string, error) {
	cmd := exec.Command("go", "list", "-m")
	cmd.Dir = dir // 设置当前工作目录
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	modulePath := strings.TrimSpace(string(output))
	if modulePath == "command-line-arguments" {
		return "", fmt.Errorf("this dir not in go project %s", dir)
	}
	return modulePath, nil
}

// getGoModPath 获取当前文件夹的 go.mod 路径
func getGoModPath(dir string) (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.GoMod}}")
	cmd.Dir = dir // 设置当前工作目录
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	goModPath := strings.TrimSpace(string(output))
	if goModPath == "" {
		return "", fmt.Errorf("this dir not in go project %s", dir)
	}
	goModPath = strings.ReplaceAll(goModPath, "go.mod", "")
	return goModPath, nil
}

func main() {
	dir, _ := os.Getwd()
	// fmt.Println("dir is :", dir)
	modulePath, err := getGoModulePath(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Println("Go Module Path:", modulePath)

	goModPath, err := getGoModPath(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Println("go.mod Path:", goModPath)
	rePaht := strings.TrimPrefix(dir, goModPath)
	goDirPath := path.Join(modulePath, rePaht)
	fmt.Println(goDirPath)
}
