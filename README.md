# 時間割システム

## リポジトリの準備
```bash
# リポジトリのクローン
git clone https://github.com/ninebolt6/tcu-timetable.git
cd tcu-timetable

# 新規ブランチを生やす
git checkout -b ブランチ名
# 既存のブランチに移動する
git branch ブランチ名
```

## 開発ルール
- issueをTODOリストの要領で使う。自分が担当するissueには自分をアサインする。
- issueにはできる限りラベルをつける。
- ブランチの命名規則を守る。新機能は`feature/XXX`、バグ修正は`hotfix/XXX`。
- コミットメッセージの頭には`feat: XXX`のように接頭辞をつける。詳細は`.gitmessage`を参照されたし。
- プルリクエストにもコミットメッセージと同様に接頭辞をつける。
- プルリクエストは必ずメンバーにレビューしてもらってからマージする。

## Dockerでの開発
```bash
docker compose build
docker compose up -d
docker compose exec go go run main.go
docker compose stop
```

```
# フロントエンド
localhost:3000

# バックエンド
localhost:8080

# データベース
localhost:3306
```