# Docker Compose for Prometheus + Grafana

- https://dev.to/ninadingole/docker-compose-for-prometheus-grafana-3gie
- https://github.com/ninadingole/docker-images

```sh
cd grafana-prometheus

cp .env.grafana.example .env.grafana
# edit .env.grafana

docker compose up -d
```

# Install Prometheus with Tar.gz

```sh
cd grafana-prometheus

curl https://github.com/prometheus/prometheus/releases/download/v2.40.5/prometheus-2.40.5.linux-amd64.tar.gz -o prometheus-2.40.5.linux-amd64.tar.gz

# 1. for ubuntu
tar -xf prometheus-2.40.5.linux-amd64.tar.gz
(cd prometheus-2.40.5.linux-amd64 && ./prometheus --config.file=../prometheus/prometheus.yml --storage.tsdb.path=../prometheus_data --web.console.libraries=./console_libraries --web.console.templates=./consoles)
# 2. for macOS
tar -xf prometheus-2.40.5.darwin-amd64.tar.gz
(cd prometheus-2.40.5.darwin-amd64 && ./prometheus --config.file=../prometheus/prometheus.yml --storage.tsdb.path=../prometheus_data --web.console.libraries=./console_libraries --web.console.templates=./consoles)
```
