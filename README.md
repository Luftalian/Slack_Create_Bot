# Slack_Create_Bot

## 環境変数

### mariadb関連
- DB_USERNAME
- DB_PASSWORD
- DB_HOSTNAME
- DB_PORT
- DB_DATABASE

### HackMD APIのトークン、チームID
- TOKEN
- TEAMID

### SlackのWebhooks用のURL
- SLACKURL

## 定期実行
- MTGの前日の通知

  例：
```
@channel
明日のMTGは [時間] からで [場所] 実施します 欠席する場合は:kesseki: 、遅刻する場合は:tikoku: を押してください
議事録：[リンク]
```
  [時間]→MTGの開始時刻が表示される

  [場所]→MTGの場所が表示される

  [リンク]→HackMD apiで作られた新たな議事録が表示される

- MTG開始時間の通知

  例：
```
@channel MTGの開始時刻です
```

## 各種設定
- ヘルプの表示
```
!md-bot-help
```
### MTG設定
- MTG時間の変更

[minute]は0または30
```
!md-bot-set [id] time [hour] [minute]
```
- MTG曜日の変更
0:Sunday, 1:Monday, 2:Tuesday, 3:Wednesday, 4:Thursday, 5:Friday, 6:Saturday
```
!md-bot-set [id] date [value]
```
- MTG場所の変更
```
!md-bot-set [id] place [value]
```
### MTG用HackMDの設定
- permissionの変更
freely, editable, limited, locked, protected, private
```
!md-bot-set [id] permission [value]
```
- commentPermissionの変更
disable, forbidden, owners, signed_in_users
```
!md-bot-set [id] commentPermission [value]
```
- readPermissionの変更
owner, sign_in, guest
```
!md-bot-set [id] commentPermission [value]
```
- writePermissionの変更
owner, sign_in, guest
```
!md-bot-set [id] commentPermission [value]
```
- MDのタイトルの変更
```
!md-bot-set [id] title [value]
```
- MDの内容の変更
```
!md-bot-set [id] content [value]
```
### Slackの設定
- 前日の通知文の変更
```
!md-bot-set [id] sendText [value]
```

## idについて
idは基本的に1

このBOTをほかのMTGで使いたくなったときに拡張できるように設定
