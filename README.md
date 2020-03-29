# parseSlackMessage

## ツールの概要

Slack の**データのエクスポート機能**により出力した各チャンネルのメッセージデータ( JSON 形式)を、 **CSV 形式に変換**するツールです。

以下の情報をフィルタリング・加工して　CSV ファイルに出力します。  
`thred_ts` は親スレッドの一意性を示す目的であるため、 UnixTime 形式から変換せずそのままの値を出力しています。

- `thread_ts` ... スレッドの識別子( スレッドとなっていないメッセージe-ta ｍの場合は値なし )
- `ts` ... タイムスタンプ ( JST )
- `user_name` ... ユーザ名( real_name + display_name )
- `text` ... 投稿メッセージ


各フィールドに関する情報は以下を参照してください。

- [Slack からエクスポートしたデータの読み方](https://slack.com/intl/ja-jp/help/articles/220556107-Slack-%E3%81%8B%E3%82%89%E3%82%A8%E3%82%AF%E3%82%B9%E3%83%9D%E3%83%BC%E3%83%88%E3%81%97%E3%81%9F%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E8%AA%AD%E3%81%BF%E6%96%B9i)
- [slack api - Retrieving messages #Threaded messages](https://api.slack.com/messaging/retrieving#threaded_messages)


動作確認は以下の環境で行っています。
- macOS Catalina 10.15.3
- Amazon WorkSpaces / Standard with Windows 10 バンドル
    - OS エディション: Windows Server 2016 Datacenter
    - OS バージョン: 1607
    - x64-based


## 使い方

### Slack からデータをエクスポート
Slack の標準機能を使って[ワークスペースのデータをエクスポート](https://slack.com/intl/ja-jp/help/articles/201658943-%E3%83%AF%E3%83%BC%E3%82%AF%E3%82%B9%E3%83%9A%E3%83%BC%E3%82%B9%E3%81%AE%E3%83%87%E3%83%BC%E3%82%BF%E3%82%92%E3%82%A8%E3%82%AF%E3%82%B9%E3%83%9D%E3%83%BC%E3%83%88%E3%81%99%E3%82%8B)してください。  

2020/03/29 時点で、以下のディレクトリ構成でエクスポートされます。

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

※ Free プランの場合
```
本ツールでは `user.json` (ワークスペースに参加しているメンバーの情報)および各チャンネルに投稿されたメッセージが保存されている `yyyy-mm-dd.json` をインプットとして使用ますが、上記のディレクトリ構成でこれらのファイルが配置されることを想定しています。



### 変換ツールの実行

本リポジトリを Clone もしくは ダウンロード後、実行環境に応じた実行ファイル( `./bin` 配下 )を使用してください。
- Mac : [parseSlackMessage](./bin/darwin64/parseSlackMessage)
- Windows : [parseSlackMessage.exe](/bin/windows64/parseSlackMessage.exe)


任意のチャンネルのディレクトリ直下( yyyy-mm-dd.json と同一の階層)に実行ファイルを保存・実行すると、カレントディレクトリに `SlackMessages.csv` を作成します。