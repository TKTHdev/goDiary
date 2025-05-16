
# 📘 goDiary - Terminal-Driven Web Diary App in Go

**goDiary** is a lightweight web diary application written in Go. It allows users to submit diary entries from the terminal, stores them in a SQLite database, and displays them on a web page. Each entry includes a SHA256 hash and a timestamp in **Japan Standard Time (JST)**. Entries can be deleted using their hash.

---

## 🚀 Features

- ✅ Submit diary entries via terminal (`curl` or CLI)
- ✅ View entries in a web browser
- ✅ Delete entries by hash
- ✅ Auto-generated JST timestamp
- ✅ Lightweight SQLite database for persistence
- ✅ Can be run persistently using `screen` or `systemd`

---

## 🛠️ Tech Stack

- Go 1.16+
- Gin Web Framework
- SQLite3
- HTML Templates

---

## 📦 Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/goDiary.git
cd goDiary
```

### 2. Install dependencies

```bash
go mod tidy
```

SQLite DB file will be created automatically if not present.

---

## ▶️ Run the Application

### Development mode

```bash
go run main.go
```

### Persistent mode using `screen`

```bash
screen -S diary
go run main.go
# Detach with Ctrl + A, then D
```

To reattach the session later:

```bash
screen -r diary
```

---

## 🌐 Access the Web Interface

After running, open your browser and navigate to:

```
http://<your-server-ip>:8080
```

---

## 📤 Submit Diary Entries from Terminal

```bash
curl -X POST http://<your-server-ip>:8080/api/diary      -d "content=Today was a beautiful day."
```

---

## ❌ Delete Entries by Hash

```bash
curl -X DELETE http://<your-server-ip>:8080/api/diary/<hash>
```

Replace `<hash>` with the SHA256 hash shown in the web interface.

---

## 🧱 Database Schema (`diary.db`)

| Column       | Type     | Description                        |
|--------------|----------|------------------------------------|
| `id`         | INTEGER  | Auto-incrementing ID               |
| `hash`       | TEXT     | Unique SHA256 hash of the entry    |
| `content`    | TEXT     | Diary entry content                |
| `created_at` | DATETIME | Timestamp in JST                   |

---

## 📚 Libraries Used

- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
- [github.com/google/uuid](https://github.com/google/uuid)

---

## 📄 License

This project is licensed under the MIT License.
