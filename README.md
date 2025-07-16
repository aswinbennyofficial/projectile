# 🚀 Projectile - Ongoing

**Projectile** is a lightweight, pluggable event router and fan-out engine.  
It connects sources (like webhooks, Kafka, or message queues) to multiple sinks (like HTTP endpoints, log files, Slack, S3, or databases), making it easy to build real-time data pipelines.

With a YAML-based config system and modular plugin support, Projectile is designed for flexibility — whether you're running it on a laptop or deploying to a cloud-native environment.


![Architecture diagram](./_nocode/images/architecture-diagram.png)


## 🔌 Supported Plugins

### Sources

| Type     | Supported |
|----------|-----------|
| `webhook` | ✅        |
| `kafka`   | ❌        |
| `rabbitMQ`| ❌        |
| `http`    | ❌        |
| `timer`   | ❌        |

### Sinks

| Type       | Supported |
|------------|-----------|
| `stdout`   | ✅        |
| `file`     | ✅        |
| `webhook`  | ✅        |
| `kafka`    | ❌        |
| `postgres` | ❌        |
| `slack`    | ❌        |
| `s3`       | ❌        |
| `rabbitMQ` | ❌        |



## 🧠 Use Cases

- 🔄 Forward webhook events to Kafka, Slack, or internal APIs
- 🔁 Mirror Kafka events into multiple systems
- 🚀 Trigger workflows or alerts from GitHub/GitLab events
- 🧪 Build pluggable event processing pipelines
- 📝 Log events to file for audit or debugging



## 🧩 Speciality

- ✅ **Modular Plugin System**: Easily add new sources or sinks with simple Go interfaces
- 🔁 **Dynamic Routing**: Define event flows declaratively via `routes.yaml`
- 🧠 **Decoupled Design**: Clean separation of concerns between config, orchestration, and plugins
- ⚙️ **Config-Driven**: All behavior driven by YAML config – no code change required for routing
- 📦 **Pluggable Transforms**: (WIP) Plan support for filters and `gojq`-style transformations
- 🧪 **Single Binary Runtime**: All functionality packed into a single Go binary – no external dependencies
- 🔧 **Unified Runtime**: Run all plugins (sources/sinks) together in one self-contained process



---

## 🧾 Example Configs

### `infra.yaml` – Secrets, DSNs, and Connection Setup

```yaml
version: v1

sources:
  webhook-main:
    type: webhook
    config:
      path: /webhook/main
      method: POST

  webhook-metrics:
    type: webhook
    config:
      path: /webhook/metrics
      method: POST

sinks:
  stdout-log:
    type: stdout
    config: {}

  file-logger:
    type: file
    config:
      path: ./logs/main/

  notify-service:
    type: webhook
    config:
      method: POST
      url: http://internal.service.local/notify
      headers:
        X-Auth-Token: abc123

```


### `routes.yaml` – Routing Logic

```yaml
version: v1

routes:
  - name: main_fanout
    source: webhook-main
    sinks:
      - stdout-log
      - file-logger
      - notify-service

  - name: metrics_alert
    source: webhook-metrics
    sinks:
      - stdout-log
      - notify-service
```


---

## 🚀 How to Run

```bash
go run cmd/projectile/main.go
```

Make sure you have the required configuration files:

- `configs/infra.yaml`
- `configs/routes.yaml`


