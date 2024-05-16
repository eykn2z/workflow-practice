# api-server

apiサーバーサンプルコード

## build & push
```
AWS_ACCOUNT_ID=xxxxxx
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin "${AWS_ACCOUNT_ID}.dkr.ecr.us-east-1.amazonaws.com"
IMAGE_URL="${AWS_ACCOUNT_ID}.dkr.ecr.us-east-1.amazonaws.com/api-server:latest"
docker build -t ${IMAGE_URL} -f ./Dockerfile .
docker push ${IMAGE_URL}
```
