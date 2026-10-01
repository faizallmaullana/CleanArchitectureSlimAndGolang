# Rekrutment GBU

Project ini memiliki dua stack aplikasi:

- Go + Gin + GORM + PostgreSQL
- PHP Slim + MongoDB

Keduanya dapat dijalankan bersamaan karena menggunakan port host yang berbeda.

## Prasyarat

- Docker
- Docker Compose

Cek instalasi:

```bash
docker --version
docker compose version
```

## Menjalankan Go + PostgreSQL

Jalankan dari root repository:

```bash
docker compose -f docker-compose.go.yaml up -d --build
```

Service yang dijalankan:

| Service | Host | Container | Keterangan |
| --- | ---: | ---: | --- |
| API Go | `8090` | `8080` | Gin + GORM |
| PostgreSQL | `5433` | `5432` | Database aplikasi |
| Adminer | `9000` | `8080` | Web UI PostgreSQL |

Akses:

- API: http://localhost:8090
- Health check: http://localhost:8090/healthz
- Adminer: http://localhost:9000

Login Adminer:

```text
System:   PostgreSQL
Server:   postgres
Username: app_user
Password: app_password
Database: app_db
```

Seeder dijalankan sebagai service Compose. Untuk menjalankannya ulang:

```bash
docker compose -f docker-compose.go.yaml run --rm seeder
```

## Menjalankan PHP + MongoDB

```bash
docker compose -f docker-compose.php.yaml up -d --build
```

Service yang dijalankan:

| Service | Host | Container | Keterangan |
| --- | ---: | ---: | --- |
| PHP API melalui Nginx | `8080` | `80` | Slim Framework |
| MongoDB | `27017` | `27017` | Database aplikasi |
| Mongo Express | `8082` | `8081` | Web UI MongoDB |
| PHP-FPM | internal | `9000` | Tidak diekspos ke host |

Akses:

- PHP API: http://localhost:8080
- Mongo Express: http://localhost:8082

Login Mongo Express:

```text
Username: admin
Password: admin_password
```

Seeder MongoDB dijalankan sebagai service Compose. Untuk menjalankannya ulang:

```bash
docker compose -f docker-compose.php.yaml run --rm seeder
```

## Endpoint API Go

Base URL: `http://localhost:8090`

| Method | Endpoint | Keterangan |
| --- | --- | --- |
| GET | `/healthz` | Health check |
| GET | `/api/products` | Mengambil semua produk |
| GET | `/api/products/{id}` | Mengambil satu produk |
| POST | `/api/products` | Membuat produk |
| PUT | `/api/products/{id}` | Mengubah produk |
| DELETE | `/api/products/{id}` | Menghapus produk |

Contoh request membuat produk:

```bash
curl -X POST http://localhost:8090/api/products \
	-H 'Content-Type: application/json' \
	-d '{"name":"Keyboard","price":250000}'
```

Contoh request mengubah produk:

```bash
curl -X PUT http://localhost:8090/api/products/1 \
	-H 'Content-Type: application/json' \
	-d '{"name":"Mechanical Keyboard","price":350000}'
```

## Endpoint API PHP

Base URL: `http://localhost:8080`

| Method | Endpoint | Keterangan |
| --- | --- | --- |
| GET | `/api/products` | Mengambil semua produk dari MongoDB |
| GET | `/api/products/{id}` | Mengambil satu produk |
| POST | `/api/products` | Membuat produk |
| PUT | `/api/products/{id}` | Mengubah produk |
| DELETE | `/api/products/{id}` | Menghapus produk |

Contoh request:

```bash
curl http://localhost:8080/api/products
```

## Melihat Status dan Log

Go stack:

```bash
docker compose -f docker-compose.go.yaml ps
docker compose -f docker-compose.go.yaml logs -f api
```

PHP stack:

```bash
docker compose -f docker-compose.php.yaml ps
docker compose -f docker-compose.php.yaml logs -f nginx
```

## Menghentikan Stack

```bash
docker compose -f docker-compose.go.yaml down
docker compose -f docker-compose.php.yaml down
```

Perintah `down` tidak menghapus volume database. Untuk menghapus container
sekaligus data database, gunakan `down -v` dengan hati-hati:

```bash
docker compose -f docker-compose.go.yaml down -v
docker compose -f docker-compose.php.yaml down -v
```
