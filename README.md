# awscli-lambda
## デプロイ
### IAMロールの作成
ポリシー作成
```bash
aws iam create-policy --policy-name LambdaExecute --policy-document file://iam-role/policy.json
```

信頼ポリシー作成
```bash
aws iam create-role --role-name lambda-execute-role --assume-role-policy-document file://iam-role/trust-policy.json
```

IAMロールにポリシーをアタッチ
```bash
aws iam attach-role-policy --role-name lambda-execute-role --policy-arn arn:aws:iam::{accountId}:policy/LambdaExecute
```

### デプロイパッケージの作成
```bash
make package
```
### デプロイ
```bash
make deploy
```
## ローカルでのテスト
```bash
make test
```