## 概要
コミットが無い場合はLINEが通知してくれるアプリケーションです。
学習継続のために作成しました。

## 構成図
<img width="600" alt="スクリーンショット 2024-07-07 午後3 41 04" src="https://github.com/user-attachments/assets/5bb5058a-5a47-43ab-bd7d-7fe8361bdb92">

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
