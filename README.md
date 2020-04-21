# Goonstone
A picture sharing web-app written in Go

## 0.0.1

- web-appの基礎構成を構築
    - Ping用ルーティングの作成
    - testを実行する
- Terraformでインフラを構築
    - applyしたあとにECRのリポジトリURLを確認しておく
- github actions
    - githubにシークレットの追加
    - pushしてみる
    
    
    
## Memo

- エラーを共通化することで、各層でエラーを分岐しクライアント側とサーバー側で伝える情報を分けるようにした
    - multierrを使用している
    
