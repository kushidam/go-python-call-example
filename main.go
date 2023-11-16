package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("python3", "./PythonScript/echoPython.py")

	// 実行して標準出力を取得
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	// 標準出力を表示
	fmt.Println(string(output))

}
