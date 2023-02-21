# /bin/sh
cd $(dirname $0)
echo "start building easycboot"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build easycboot.go
echo "finish building easycboot"
cd ./trace-ui
echo 'start build trace-ui'
pnpm build
echo 'finish build trace-ui'