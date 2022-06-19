#!/bin/bash
echo "download gs-reader-lite from github..."
wget -O gs-reader-lite https://github.com/GuoShaoOrg/GSReader-Lite/releases/download/Latest/gs-reader-lite

chmod 755 gs-reader-lite

echo "start running..."

# shellcheck disable=SC2046
# shellcheck disable=SC2006
kill -9 `cat pidfile.txt`

rm pidfile.txt

nohup ./gs-reader-lite & echo $! > pidfile.txt

echo "end"

exit