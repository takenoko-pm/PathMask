# PathMask CLI

コピーしたパスのユーザー名を匿名化してクリップボードに戻す小さな CLI ツールです。
LLMなどにコマンドの結果などを渡すときにご使用ください。

## ダウンロード
1. [Releases ページ](https://github.com/takenoko-pm/PathMask/releases/latest) にアクセスします。

2. お使いの OS に合わせた ZIP ファイルをダウンロードしてください。
   - Windows: `pathmask-win-amd64.zip`
   - Mac (M1/M2/M3): `pathmask-mac-arm64.zip`
   - Mac (Intel): `pathmask-mac-amd64.zip`
   - Linux: `pathmask-linux-amd64.zip`

3. ZIP を展開し、中にある実行ファイルにパスを通すか、任意の場所で実行してください。

※ Linux は xclip または xsel が必要です

## 使い方

1. ターミナル / PowerShell で展開先に移動  
2. コマンド実行
```powershell
# Windows
.\pathmask.exe

# Linux / Mac
./pathmask
```

クリップボードにテキストをコピーした状態で使用すると、ユーザー名を[USER]に変換した結果を出力し、クリップボードの中身も編集します。

例
```
PS C:\Users\tarou> where.exe python
C:\Users\tarou\AppData\Local\Programs\Python\Python313\python.exe
C:\Users\tarou\AppData\Local\Microsoft\WindowsApps\python.exe
```
↓
```
PS C:\Users\[USER]> where.exe python
C:\Users\[USER]\AppData\Local\Programs\Python\Python313\python.exe
C:\Users\[USER]\AppData\Local\Microsoft\WindowsApps\python.exe
```
