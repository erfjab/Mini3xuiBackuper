# Mini 3x-ui Backuper

A small tool that logs into a [3x-ui](https://github.com/MHSanaei/3x-ui) panel and downloads the database as a backup file.

---

## Download

| Architecture | File             |
|--------------|------------------|
| x86-64       | `backuper-amd64` |
| ARM64        | `backuper-arm64` |
| ARMv7        | `backuper-armv7` |

Download the binary from the [Releases](../../releases) page.

---

## Setup

**1. Create a directory with any name you like:**

```bash
mkdir mybackuprt
cd mybackuprt
```

**2. Download the binary into it** (example for x86-64):


```bash
wget https://github.com/erfjab/Mini3xuiBackuper/releases/latest/download/backuper-amd64
```

**3. Create a `.env` file:**

```bash
cat > .env << 'EOF'
#### Panel configuration

## admin username
## like: "admin" or "Jksd45"
PANEL_USERNAME="user"

## admin password
## like: "gewro" or "Nsdf8"
PANEL_PASSWORD="pass"

## panel host
## http or https, with port and path if needed  
## like: "http://domain.top:2053/dashboard"
PANEL_HOST="http://localhost:8080"
EOF
```

**4. Give execute permission:**

```bash
chmod +x backuper-amd64
```

**5. Test it:**

```bash
./backuper-amd64
```

If it works, you will see a `.db` file in the same directory:

```
2026-03-13_14-05-30_admin.db
```

---

## Cron Job

Once the test works, open `crontab -e` and add:

#### every 1 minute
```cron
* * * * * cd /root/mybackuprt && ./backuper-amd64 >> backup.log 2>&1
```
#### every 5 minutes
```cron
*/5 * * * * cd /root/mybackuprt && ./backuper-amd64 >> backup.log 2>&1
```

#### every 10 minutes
```cron
*/10 * * * * cd /root/mybackuprt && ./backuper-amd64 >> backup.log 2>&1
```

#### every 30 minutes
```cron
*/30 * * * * cd /root/mybackuprt && ./backuper-amd64 >> backup.log 2>&1
```

Pick one line and remove the rest.

> To run multiple panels at the same time, repeat the setup steps in a new directory for each panel, then add a separate cron line for each one.


