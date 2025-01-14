#!/bin/sh

# Установка необходимых утилит, если их нет
if ! command -v getent > /dev/null 2>&1; then
  echo "Устанавливаю необходимые утилиты..."
  apk add --no-cache busybox-extras
fi

if ! command -v awk > /dev/null 2>&1; then
  echo "Устанавливаю awk..."
  apk add --no-cache awk
fi

# Получение IP-адреса целевого контейнера
TARGET_IP=$(getent hosts keycloak | awk '{ print $1 }')

# Проверка, что IP был найден
if [ -z "$TARGET_IP" ]; then
  echo "Ошибка: Не удалось определить IP-адрес для keycloak."
  exit 1
fi

# Обновление /etc/hosts
echo "$TARGET_IP localhost" >> /etc/hosts

echo "localhost теперь резолвится в $TARGET_IP"

