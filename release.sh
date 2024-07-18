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
git tag -d v1.0.0
git push origin :refs/tags/v1.0.0
git tag v1.0.0
git push origin v1.0.0
go list -m github.com/ylighgh/happy-go@v1.0.0