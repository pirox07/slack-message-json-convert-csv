# parseSlackMessage

## ツールの概要

Slack に投稿されたメッセージを JSON 形式でエクスポート後、 CSV 形式に変換するための簡易ツールです。  
CSV ファイルには以下の情報をフィルタリングして出力します。  
- thread_ts ... スレッド ID （同一スレッドへの投稿は、同一 thread_ts が割り当てられる）
- ts ... タイムスタンプ ( JST )
- real_name ... ユーザ名
- text ... 投稿メッセージ

※ Mac での実行を想定しています。

## 使い方

[ワークスペースのデータをエクスポート](https://slack.com/intl/ja-jp/help/articles/201658943-%E3%83%AF%E3%83%BC%E3%82%AF%E3%82%B9%E3%83%9A%E3%83%BC%E3%82%B9%E3%81%AE%E3%83%87%E3%83%BC%E3%82%BF%E3%82%92%E3%82%A8%E3%82%AF%E3%82%B9%E3%83%9D%E3%83%BC%E3%83%88%E3%81%99%E3%82%8B)してください。

以下のようにチャンネル単位でサブフォルダが作成され、 JSON ファイルがエクスポートされます。  

格納されている JSON ファイル( yyyy-mm-dd.json )を、 `source` フォルダ内に保存してください。  

```
./parseSlackMessage
```
同一フォルダ内に `SlackMessages.csv` ファイルが作成されます。



## 参考

- [Slack からエクスポートしたデータの読み方](https://slack.com/intl/ja-jp/help/articles/220556107-Slack-%E3%81%8B%E3%82%89%E3%82%A8%E3%82%AF%E3%82%B9%E3%83%9D%E3%83%BC%E3%83%88%E3%81%97%E3%81%9F%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E8%AA%AD%E3%81%BF%E6%96%B9)