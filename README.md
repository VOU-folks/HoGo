# hogo

### HoGo

Hosting management system

**Warning**: clone it and adapt for own use cases

---

### Deployment

---
#### env var:
This file is for sensitive variables (like keys, secrets and etc).

Make a copy of .env.example to .env
```shell
cp .env.example .env
```
... and define required variables

---

#### hogo.conf:
This file is for service specific non-sensitive configs (like flags, addresses and etc).

After building and copying to `/etc/hogo/hogo.conf` check for necessary parameters before starting as systemd service.

---

#### unit file in:
```
/usr/lib/systemd/system
```

```shell
mkdir -p /etc/hogo/
mkdir -p /var/log/hogo
rm /var/log/hogo/*
go build -o /etc/hogo/hogo-manager main.go
chmod +x /etc/hogo/hogo-manager
cp hogo.conf /etc/hogo/hogo.conf
cp .env /etc/hogo/.env

cp hogo-manager.service /usr/lib/systemd/system/
systemctl daemon-reload
systemctl enable hogo-manager
systemctl restart hogo-manager
systemctl status hogo-manager
```

```shell
# to cleanup previous deploy
systemctl stop hogo-manager
systemctl disable hogo-manager
pkill -f hogo-manager
```