#!/bin/bash

readonly NGROK_BASE_URL="http://localhost:4040"

docker compose up &
DC_PID=$!

until curl -s $NGROK_BASE_URL/api/tunnels >/dev/null; do
    sleep 2
done

NGROK_URL=$(curl -s $NGROK_BASE_URL/api/tunnels | grep -o '"public_url":"[^"]*' | grep -o '[^"]*$')
MESSAGE="Ngrok URL: $NGROK_URL"
echo -e "\033[1;31m$MESSAGE\033[0m"

wait $DC_PID
