#!/bin/bash
rm -f gotro
go build .
sstrip -z gotro
upx --lzma -9 gotro
ls -al gotro
./gotro

