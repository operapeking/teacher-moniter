#!/bin/bash
dir=$(pwd)
export GIN_MODE=release
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