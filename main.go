package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/atotto/clipboard"
)

func main() {
	text := ""
	isClip := false
	// 標準入力があるか確認
	if file, _ := os.Stdin.Stat(); (file.Mode() & os.ModeCharDevice) == 0 {
		// 標準入力から読み込む
		data, err := io.ReadAll(os.Stdin)
		if err == nil {
			text = string(data)
		}
	}
	// 標準入力が空ならクリップボードを読む
	if text == "" {
		data, err := clipboard.ReadAll()
		if err != nil {
			fmt.Fprintln(os.Stderr, "クリップボード読み込みエラー:", err)
			return
		}
		text = data
		isClip = true
	}
	if text == "" {
		return // どちらも空なら終了
	}
	// 置換
	// ユーザー名部分を検出する正規表現
	re := regexp.MustCompile(`(?i)([/\\])(Users|home)([/\\])[^/\r\n\\>]+`)
	replaced := re.ReplaceAllString(text, "${1}${2}${3}[USER]")
	// 出力
	fmt.Print(replaced)
	// クリップボードから読んでいたらクリップボードに戻す
	if isClip {
		if err := clipboard.WriteAll(replaced); err != nil {
			fmt.Fprintln(os.Stderr, "クリップボード書き込みエラー:", err)
		}
	}
}
