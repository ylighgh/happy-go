#!/bin/sh

# 获取命令行输入的提交消息
if [ $# -eq 0 ]; then
  echo "请输入提交消息"
  exit 1
fi
commit_message="$*"

# 提交和推送Git更改
git add -A
git commit -m "$commit_message"
git tag -d v2.0.0
git tag v2.0.0
git push origin v2.0.0
git push
go list -m github.com/ylighgh/happy-go@v2.0.0