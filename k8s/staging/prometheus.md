# Prometheus config for App Monitoring

```yaml
    - job_name: kaznet-status-staging
      metrics_path: /metrics
      static_configs:
      - targets: [ 'kaznet-status-staging-service.kaznet-status-staging.svc.cluster.local:2112']
```