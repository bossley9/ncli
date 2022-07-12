#!/bin/sh

echo "Enter Notion integration token:"
read -r token

if [ -z $token ]; then
  echo "Invalid token. Aborting"
  return
fi

echo "package client\n\nconst NOTION_TOKEN = \"${token}\"" > "./pkg/client/token.go"

echo "Token set!"
