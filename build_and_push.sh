
go mod tidy
GOOS=linux GOARCH=amd64 go build

docker build --no-cache -t smart-mon --platform=linux/amd64 .
docker tag smart-mon docker.io/jbalcas/smart-mon:latest
docker push docker.io/jbalcas/smart-mon
