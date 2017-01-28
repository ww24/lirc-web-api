LIRC Web API
===

[![wercker status](https://app.wercker.com/status/b9a4cc28becedddbff9ee59b19a54f47/m/master "wercker status")](https://app.wercker.com/project/byKey/b9a4cc28becedddbff9ee59b19a54f47)

[LIRC](http://www.lirc.org/) client implemented by golang.

Install
---
```
curl https://raw.githubusercontent.com/ww24/lirc-web-api/master/install.sh | sh
```

Usage
---
### Health check
```
curl http://localhost:3000/status
```

### List
```
curl http://localhost:3000/api/v1/signals
```

```json
{
  "status":"ok",
  "signals":[
    {
      "remote":"aircon",
      "name":"on"
    },
    {
      "remote":"aircon",
      "name":"off"
    }
  ]
}
```

### Send
```
curl -XPOST http://localhost:3000/status/aircon/on
```

```
remote --> aircon
name --> on
```

### Web Frontend
Open `http://localhost:3000/` in Google Chrome or Android Chrome.
