#!/bin/bash

# run monday to friday at 0830h
aRON_JOB="30  8 * *  1-5 /usr/local/go/bin/go run /path/to/your/app.go"

if ! crontab -l | grep -Fxq "$CRON_JOB"; then
  (crontab -l ; echo "$CRON_JOB") | crontab -
else
  echo "automata already awake!"
fi

echo "automata has been awaken"
