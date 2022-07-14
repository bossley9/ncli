#!/bin/sh

echo "Enter Notion integration token:"
read -r token

if [ -z $token ]; then
  echo "Invalid token. Aborting"
  return
fi

echo "package notion\n\nconst NOTION_TOKEN = \"${token}\"" > "./pkg/notion/token.go"

echo "Token set!"
