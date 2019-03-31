# デモサイト

[AWS Serverless Uploader](https://aws-serverless-uploader.netlify.com/)

# 概要

- AWS Lambda APIGateway DynamoDB S3 を用いたサーバレスアプリケーション
- DynamoDB や S3 へのアクセスを interface で抽象化
- Localstack を用いた自動テスト
- フロントエンドは React + TypeScript + Netlify

# フロントエンド

[aws-serverless-uploader-frontend](https://github.com/dyoshikawa/aws-serverless-uploader-frontend)

# セットアップ

## 環境

- Go 1.11.2
- dep
- Node 8.11.0
- npm or yarn
- aws-cli
- direnv

## デプロイ

```
git clone https://github.com/dyoshikawa/aws-serverless-uploader-go
cd aws-serverless-uploader-go
yarn
make deploy
```

## 自動テスト

### aws-cli Localstack 用の Profile を作成

```
aws configure --profile localstack
Default region name [None]: us-east-1
Default output format [None]: json
```

最初の 2 項目は任意の値で OK。

### Localstack 起動とセットアップ

```
cp .envrc.example .envrc
docker-compose up -d
make localstack-setup
```

### テスト開始

```
make test
```
