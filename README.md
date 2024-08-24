## 概要
コミットが無い場合はLINEが通知してくれるアプリケーションです。
学習継続のために作成しました。

## 構成図
<img width="786" alt="構成図" src="https://github.com/user-attachments/assets/aac634d4-098a-4d50-89b9-2fd8cec68350">

## 使用技術
- golang
- aws (sam, eventbridge, ecr, lambda, iam, kms)
- terraform
- docker
- github actions
- line notify api, github api

## ローカル環境実行方法
1. asw sam cliをインストール
2. env.example.jsonを参考に、env.jsonを作り必要情報を記入する
3. sam local invoke　-n env.jsonを実行
