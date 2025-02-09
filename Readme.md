# Ecommerce-UMS

Ecommerce-UMS adalah sistem manajemen pengguna (User Management System) berbasis Echo, sebuah Golang Web Framework, dengan database PostgreSQL yang dikemas menggunakan Docker Compose.

## 🚀 Fitur

- **CRUD User**: Fasilitas untuk membuat, membaca, memperbarui, dan menghapus data pengguna.
- **Autentikasi JWT**: Sistem autentikasi berbasis token menggunakan JSON Web Token (JWT).
- **Manajemen Sesi**: Kemampuan untuk menyimpan sesi pengguna melalui user_sessions.
- **Dockerized**: Pemanfaatan Docker Compose untuk menjalankan aplikasi dengan PostgreSQL secara efisien.

## 📦 Cara Instalasi & Menjalankan Aplikasi
### 1️⃣ Clone Repository
```bash
git clone https://github.com/username/ecommerce-ums.git
cd ecommerce-ums
```

### 2️⃣ Menjalankan dengan Docker Compose
Pastikan Docker dan Docker Compose telah terinstal di sistem Anda sebelum melanjutkan.
```bash
docker-compose up --build -d
```
### 3️⃣ Memeriksa Container yang Berjalan
Gunakan perintah berikut untuk memastikan container berjalan dengan baik:
```bash
docker ps
```
### 4️⃣ Mengakses API
Untuk mengakses API login, gunakan perintah curl berikut:
```bash
curl -X POST http://localhost:9001/user/v1/login -H "Content-Type: application/json" -d '{
  "username": "admin",
  "password": "admin789"
}'
```
