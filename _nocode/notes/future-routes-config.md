```yaml
version: v2

routes:
  - name: main_fanout
    source: webhook-main
    decoder: avro

    rules:
      - condition: event.error == "x"
        transforms:
          - type: redact
            fields: ["user.password", "meta.secret"]
        sinks:
          - notify-service

      - condition: event.type == "info"
        sinks:
          - stdout-log

      - condition: true   # default catch-all
        sinks:
          - file-logger
```
