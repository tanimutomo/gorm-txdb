# gorm-txdb

Gormのテストで，シード挿入と並列実行をいい感じにやる目的で，以下のことを実装したサンプルです．

- Daoパッケージのテストを並列実行する
- テストの冒頭でいい感じにシードを挿入する

詳しくは以下の記事で解説しています．

https://zenn.dev/tanimutomo/articles/gorm-txdb-test