package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"

	"github.com/atotto/clipboard"
)

func main() {
	// クリップボードを読み込む
	data, err := clipboard.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
		return
	}
	if data == "" {
		return // クリップボードが空なら何もしない
	}
	text := string(data)
	// 置換
	var re *regexp.Regexp // ユーザー名部分を検出する正規表現
	var replaced string
	switch runtime.GOOS {
	case "windows":
		// Windowsの場合: C:\Users\tarou\ → C:\Users\[USER]\
		re = regexp.MustCompile(`Users\\[^\r\n\\>]+`)
		replaced = re.ReplaceAllString(text, "Users\\[USER]")
	case "darwin":
		// Macの場合: /Users/tarou/ → Users/[USER]/
		re = regexp.MustCompile(`/Users/[^\r\n/>]+`)
		replaced = re.ReplaceAllString(text, "/Users/[USER]")
	case "linux":
		// Linuxの場合: /home/tarou/ → home/[USER]/
		re = regexp.MustCompile(`/home/[^\r\n/>]+`)
		replaced = re.ReplaceAllString(text, "/home/[USER]")
	default:
		return
	}
	// 出力
	fmt.Print(replaced)
	// クリップボードに戻す
	if err := clipboard.WriteAll(replaced); err != nil {
		fmt.Fprintln(os.Stderr, "クリップボード書き込みエラー:", err)
	}
}
