# Goonstone

VueとLaravelのチュートリアル（[vue-splash](https://www.hypertextcandy.com/vue-laravel-tutorial-introduction)）を参考に、フロントをNuxt、サーバーサイドをGoに置き換えた写真共有アプリのプロジェクトです。

GoによるサーバーサイドAPIの学習を目的としています。

APIのディレクトリ構成はヘキサゴナルアーキテクチャを意識してレイヤーを分けています（[参考](https://qiita.com/rema424/items/9ffbdf584b705cae6a19)）。

TerraformでAWSのインフラを構築しており、GitHub Actionsを利用してCI/CD環境を構築、デプロイを自動化しています。

## 使用技術

- フロント
  - Nuxt 2.11
  - @nuxtjs/axios 5.9
  - etc...
  
- サーバーサイド
  - Go 1.13
  - echo 4.1
  - echo-swagger 0.0 (2020/03/14)
  - gorm 1.9
  - multierr 1.5
  - etc...
  
- データベース
  - MySQL 8
  - Redis 5
  
- インフラ
  - AWS (ECSがメイン)
  - Docker / docker-compose
  - Terraform
  - GitHub Actions
  


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
    
