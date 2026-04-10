#!/usr/bin/env bash
set -euo pipefail

echo "Instalando scout-cli..."

if ! command -v go >/dev/null 2>&1; then
  echo "Go no está instalado."
  exit 1
fi

mkdir -p bin
go build -o bin/scout-cli ./cmd/scout-cli

echo "Listo. Ejecuta:"
echo "  ./bin/scout-cli help"