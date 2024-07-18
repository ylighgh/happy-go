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

LATEST_TAG=$(git describe --tags --abbrev=0)

if [ -z "$LATEST_TAG" ]; then
  echo "未能获取到最新的tag"
  exit 1
fi

IFS='.' read -r -a VERSION_PARTS <<< "$LATEST_TAG"
MAJOR=${VERSION_PARTS[0]}
MINOR=${VERSION_PARTS[1]}
PATCH=${VERSION_PARTS[2]}

PATCH=$((PATCH + 1))

NEW_TAG="$MAJOR.$MINOR.$PATCH"
git tag "$NEW_TAG"

git push origin $NEW_TAG
go list -m github.com/ylighgh/happy-go@$NEW_TAG