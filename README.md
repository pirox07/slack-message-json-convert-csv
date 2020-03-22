# parseSlackMessage

## ツールの概要

Slack の**データのエクスポート機能**により各チャンネルのメッセージを JSON 形式でエクスポートしたデータを、 **CSV 形式に変換**するためのツールです。

CSV ファイルには以下の情報をフィルタリングして出力します。  
- thread_ts ... スレッドの ID 
- ts ... タイムスタンプ ( JST )
- user_name ... ユーザ名( real_name + display_name )
- text ... 投稿メッセージ

本ツールでは Mac OS での利用を想定しています。

## 使い方

### データのエクスポート
Slack の標準機能を使って[ワークスペースのデータをエクスポート](https://slack.com/intl/ja-jp/help/articles/201658943-%E3%83%AF%E3%83%BC%E3%82%AF%E3%82%B9%E3%83%9A%E3%83%BC%E3%82%B9%E3%81%AE%E3%83%87%E3%83%BC%E3%82%BF%E3%82%92%E3%82%A8%E3%82%AF%E3%82%B9%E3%83%9D%E3%83%BC%E3%83%88%E3%81%99%E3%82%8B)してください。  

エクスポートデータは以下のような構造となっています。  

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
本ツールでは `user.json` (ワークスペースに参加しているメンバーの情報)及び各チャンネルに投稿されたメッセージが保存されている `yyyy-mm-dd.json` をインプットとして使用します。



### 変換ツールの実行

本リポジトリを Clone もしくは ダウンロードしてください。


`./tool` 内の [`parseSlackMessage`](./tool/parseSlackMessage) を、 CSV 形式に変換したいチャンネルのディレクトリ直下に保存・実行すると、
同一ディレクトリ内に `SlackMessages.csv` が作成されます。


## 参考

- [Slack からエクスポートしたデータの読み方](https://slack.com/intl/ja-jp/help/articles/220556107-Slack-%E3%81%8B%E3%82%89%E3%82%A8%E3%82%AF%E3%82%B9%E3%83%9D%E3%83%BC%E3%83%88%E3%81%97%E3%81%9F%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E8%AA%AD%E3%81%BF%E6%96%B9)