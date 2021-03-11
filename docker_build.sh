# /bin/sh
if [ ! -n "$1" ] ;then
    echo "请指定镜像tag"
    exit 1
fi
cd $(dirname $0)
sh ./build_linux.sh
docker build -t registry.cn-beijing.aliyuncs.com/ivi/go-easy-trace-collector:$1 ./
echo "build img registry.cn-beijing.aliyuncs.com/ivi/go-easy-trace-collector:$1"