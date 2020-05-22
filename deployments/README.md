# Goonstone

Vue と Laravel のチュートリアル（[vue-splash](https://www.hypertextcandy.com/vue-laravel-tutorial-introduction)）を参考に、フロントを Nuxt、サーバーサイドを Go に置き換えた写真共有アプリのプロジェクトです。

Go によるサーバーサイド API の学習を目的としています。

API のディレクトリ構成はヘキサゴナルアーキテクチャを意識してレイヤーを分けています（[参考](https://qiita.com/rema424/items/9ffbdf584b705cae6a19)）。

Terraform で AWS のインフラを構築しており、GitHub Actions を利用して CI/CD 環境を構築、デプロイを自動化しています。

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
  - redigo 2.0
  - multierr 1.5
  - etc...

- データベース

  - MySQL 8
  - Redis 5

- インフラ
  - AWS (ECS がメイン)
  - Docker / docker-compose
  - Terraform
  - GitHub Actions


