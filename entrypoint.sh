#!/usr/bin/env bash

set -e

syslogd

echo "[ENTRYPOINT] - Starting go_simple_rest."
echo $ENVIRONMENT

# /usr/sbin/consul-template \
#     -template "/etc/go_simple_rest/config/config.toml.ctmpl:/etc/go_simple_rest/config/dev.toml" \
#     -consul-retry-attempts=0 \
#     -exec "/usr/bin/supervisord -c /etc/supervisord.conf"

go run main.go
