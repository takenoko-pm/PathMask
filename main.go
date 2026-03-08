package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"

	"github.com/atotto/clipboard"
)

func checkClipboardCmd() {
	if runtime.GOOS != "linux" {
		return
	}
	// xclipかxselのどちらかがあるか確認
	if _, err := exec.LookPath("xclip"); err == nil {
		return
	}
	if _, err := exec.LookPath("xsel"); err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, "Linuxでクリップボード操作するには xclip または xsel が必要です")
	os.Exit(1)
}

func main() {
	//Linuxのxclipかxsel確認
	checkClipboardCmd()
	// クリップボードを読み込む
	data, err := clipboard.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
		return
	}
	text := string(data)
	// ユーザー名部分を検出する正規表現
	var re *regexp.Regexp
	if runtime.GOOS == "windows" {
		// Windowsの場合: C:\Users\tarou\ → C:\Users\<USER>\
		re = regexp.MustCompile(`Users\\[^\\]+`)
	} else {
		// Mac, Linuxの場合: /home/tarou/ → home/<USER>/
		re = regexp.MustCompile(`/home/[^/]+`)
	}
	// 置換
	var replaced string
	if runtime.GOOS == "windows" {
		replaced = re.ReplaceAllString(text, "Users\\<USER>")
	} else {
		replaced = re.ReplaceAllString(text, "/home/<USER>")
	}
	// 出力
	fmt.Print(replaced)
	// クリップボードに戻す
	if err := clipboard.WriteAll(replaced); err != nil {
		fmt.Fprintln(os.Stderr, "クリップボード書き込みエラー:", err)
	}
}
