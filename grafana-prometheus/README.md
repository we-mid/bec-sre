# Docker Compose for Prometheus + Grafana

- https://dev.to/ninadingole/docker-compose-for-prometheus-grafana-3gie
- https://github.com/ninadingole/docker-compose-stacks/blob/master/prometheus-grafana
- https://github.com/vegasbrianc/prometheus
- https://grafana.com/docs/grafana-cloud/quickstart/docker-compose-linux/

```sh
cd grafana-prometheus

cp .env.grafana.example .env.grafana
# edit .env.grafana

docker compose up -d
# visit http://localhost:4000/
```

- [ ] cadvisor
- [ ] alertmanager
- [ ] dashboard +docker
- [x] dashboard +nginx +prometheus
- [x] dashboard +pm2 (w/ docker extra_hosts)
- [x] dashboard +node-exporter
