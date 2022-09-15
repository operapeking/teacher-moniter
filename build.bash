#!/bin/bash
echo "请确保环境变量 GIN_MODE 为 release。按任意键继续。"
read
echo "正在编译"
dir=$(pwd)
for target in "linux" "windows"
do
    export GOOS=$target
    for arch in "amd64"
    do
        export GOARCH=$arch
        for app in "api" "client" "server"
        do
            cd $dir/$app
            if test $target == "windows" ; then
                go build -o $dir/bin/$app-$target-$arch.exe
            else
                go build -o $dir/bin/$app-$target-$arch
            fi
        done
    done
done