# Prometheus config for App Monitoring

```yaml
    - job_name: kaznet-status-production 
      metrics_path: /metrics
      static_configs:
      - targets: [ 'kaznet-status-production-service.kaznet-status-production.svc.cluster.local:2112']
```