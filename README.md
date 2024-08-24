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
1. sam cli, aws cli, dockerのセッティング
- https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/getting-started-install.html
- https://docs.aws.amazon.com/ja_jp/serverless-application-model/latest/developerguide/install-sam-cli.html
- https://docs.docker.com/desktop/install/mac-install/

2. github tokenとline notify tokenを用意
- https://docs.github.com/ja/rest/authentication/authenticating-to-the-rest-api?apiVersion=2022-11-28#personal-access-token
- https://notify-bot.line.me/ja/

3. init-ssm.example.shをコピーし、init-ssm.shを作成
```sh
cp init-ssm.example.sh init-ssm.sh
```

4. init-ssm.shにそれぞれ 2 で取得したtokenを設定

5. init-ssm.shの権限変更
```sh
chmod +x init-ssm.sh
```

6. env.example.jsonをコピーし、env.jsonを作成
```sh
cp env.example.json env.json
```

7. env.jsonのgithub_userを自分のものに変更

8. dockerを起動
```sh
docker compose up -d
```

9. アプリケーションのビルドを実行
```sh
sam build
```

10. アプリケーションを実行
```sh
sam local invoke GoLambdaFunction -n env.json
```

## デプロイ手順
1. AWS認証用のroleを作成
https://zenn.dev/kou_pg_0131/articles/gh-actions-oidc-aws

2. Github Secretsに必要情報を入力
```sh
AWS_ACCOUNT_ID
AWS_REGION
AWS_ROLE_NAME <= 1で作成したrole名
ECR_REPOSITORY <= no-commit-notify
LAMBDA_FUNCTION_NAME <= no-commit-notify
USER_GITHUB <= githubのユーザー名
```
3. mainブランチにpush
(失敗した場合はgithub actionsのworkflowを手動で実行)
