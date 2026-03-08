# PathMask CLI

コピーしたパスのユーザー名を匿名化してクリップボードに戻す小さな CLI ツールです。
LLMなどにコマンドの結果などを渡すときにご使用ください
---

## ダウンロード

- Windows 64bit: `release/pathmask-win-amd64.zip`
- Linux 64bit: `release/pathmask-linux-amd64.zip`
- Mac 64bit: `release/pathmask-mac-amd64.zip` または `release/pathmask-mac-arm64.zip`

※ Linux は xclip または xsel が必要です

---

## 使い方

1. ZIP を展開  
2. ターミナル / PowerShell で展開先に移動  
3. コマンド実行

クリップボードにテキストをコピーした状態で使用すると、ユーザー名を<USER>に変換した結果を出力し、クリップボードの中身も編集します。

```powershell
# Windows
.\pathmask.exe

# Linux / Mac
./pathmask
```
