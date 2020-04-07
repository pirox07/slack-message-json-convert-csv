# slack-message-json-convert-csv

## ツールの概要
Slack のデータのエクスポート機能により出力した各チャンネルのメッセージデータ( JSON 形式)を、 CSV 形式に変換するツールです。

以下の情報を CSV ファイルに出力します。

- `thread_ts` ... スレッド ID ( スレッド形式ではないメッセージの場合は空白 )
- `ts` ... タイムスタンプ ( JST )
- `user_name` ... ユーザ名( real_name + display_name )
- `text` ... 投稿メッセージ

動作確認は以下の環境で行っています。
- macOS Catalina 10.15.3
- Amazon WorkSpaces / Standard with Windows 10 バンドル
    - OS エディション: Windows Server 2016 Datacenter
    - OS バージョン: 1607
    - x64-based


## 開発環境
- macOS Catalina 10.15.3
- go 1.12.4

## インストール

[Go のインストール](https://golang.org/doc/install)後、 `tz` パッケージを追加でインストールします。
```
go get 4d63.com/tz
```

## ビルド
`Makefile` が配置されているディレクトリで `make` コマンドを実行し、 macOS 、 Windows 用の実行ファイルを生成します。

```
make
rm -rf ./bin/darwin64/slack-message-json-convert-csv
rm -rf ./bin/windows64/slack-message-json-convert-csv.exe
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin64/slack-message-json-convert-csv ./main.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows64/slack-message-json-convert-csv.exe ./main.go
```

## 使い方

### Slack からデータをエクスポート
Slack の標準機能を使って[ワークスペースのデータをエクスポート](https://slack.com/intl/ja-jp/help/articles/201658943-%E3%83%AF%E3%83%BC%E3%82%AF%E3%82%B9%E3%83%9A%E3%83%BC%E3%82%B9%E3%81%AE%E3%83%87%E3%83%BC%E3%82%BF%E3%82%92%E3%82%A8%E3%82%AF%E3%82%B9%E3%83%9D%E3%83%BC%E3%83%88%E3%81%99%E3%82%8B)してください。  

エクスポートしたデータの構造は以下のとおりです。( 2019/03/29 現在)

```
<Workspace Name> Slack export <term>
├ user.josn
├ channels.json
├ integration_logs.json
├ channel-01
│  ├ yyyy-mm-dd.json
│  └ yyyy-mm-dd.json
└ channel-02
   └ yyyy-mm-dd.json

※ フリー / スタンダードプランの場合
```

各ファイルの詳細については以下を参照してください。

- [Slack からエクスポートしたデータの読み方](https://slack.com/intl/ja-jp/help/articles/220556107-Slack-%E3%81%8B%E3%82%89%E3%82%A8%E3%82%AF%E3%82%B9%E3%83%9D%E3%83%BC%E3%83%88%E3%81%97%E3%81%9F%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E8%AA%AD%E3%81%BF%E6%96%B9i)

### 変換ツールの実行

任意のチャンネルのディレクトリ直下( `yyyy-mm-dd.json` と同一の階層)に実行ファイルを保存・実行すると、カレントディレクトリに `SlackMessages.csv` を作成します。