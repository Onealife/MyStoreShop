# Nutchapholshop (เพื่อการศึกษาเท่านัน)

แอปพลิเคชัน E-Commerce ขนาดกลาง ที่พัฒนาด้วย Go และ PostgreSQL

## 📋 รายละเอียดโปรเจกต์

Nutchapholshop เป็นระบบ e-commerce ที่มีฟีเจอร์ครบครัน เช่น:
- 🔐 ระบบการตรวจสอบสิทธิ์ (Authentication)
- 🛒 ระบบตะกร้าสินค้า
- 📦 ระบบจัดการสินค้า
- 📂 ระบบหมวดหมู่สินค้า
- 🛍️ ระบบคำสั่งซื้อ
- 💳 ระบบชำระเงิน
- 📊 ระบบสถิติ
- 👤 ระบบผู้ใช้งาน
- 🔑 ระบบการจัดการสิทธิ์ (Role Management)
- ⚙️ แผงควบคุมผู้ดูแลระบบ (Admin Panel)

## 🛠️ เทคโนโลยีที่ใช้

- **Language**: Go 1.x
- **Database**: PostgreSQL
- **API Documentation**: Swagger/OpenAPI
- **Container**: Docker & Docker Compose

## 📁 โครงสร้างโปรเจกต์

```
├── cmd/                          # จุดเริ่มต้นของแอปพลิเคชัน
│   ├── api/                      # API Server
│   └── migrate/                  # Database Migration Tool
├── internal/                     # โค้ดภายในของแอปพลิเคชัน
│   ├── adapters/                 # Adapter & Implementation
│   │   ├── db/                   # Database connection
│   │   ├── http/                 # HTTP handlers & middleware
│   │   └── persistence/          # Data models & repositories
│   └── core/                     # Core business logic
│       ├── domain/               # Domain entities & models
│       ├── ports/                # Interfaces/Contracts
│       └── services/             # Business logic services
├── pkg/                          # Utility packages
│   └── utils/                    # Helper functions
├── config/                       # Configuration files
├── docs/                         # API documentation
└── docker-compose.yml            # Docker configuration
```

## 🚀 เริ่มต้นใช้งาน

### ความต้องการเบื้องต้น

- Go 1.18+
- PostgreSQL 12+
- Docker & Docker Compose (สำหรับการรัน Container)

### การติดตั้ง

1. **Clone repository**
   ```bash
   git clone <repository-url>
   cd Nutchapholshop
   ```

2. **ติดตั้ง dependencies**
   ```bash
   go mod download
   ```

3. **ตั้งค่าตัวแปรสภาพแวดล้อม**
   ```bash
   # สร้างไฟล์ .env จากตัวอย่าง
   cp .env.example .env
   ```

4. **รัน Database migrations**
   ```bash
   go run cmd/migrate/main.go
   ```

5. **เริ่มต้น API Server**
   ```bash
   go run cmd/api/main.go
   ```

### การรัน Docker

```bash
docker-compose up -d
```

## 📚 API Documentation

API documentation สามารถดูได้ที่:
- Swagger UI: `http://localhost:8080/swagger/index.html`

API endpoints สำหรับ:
- **Authentication** - เข้าสู่ระบบ, ลงทะเบียน
- **Products** - ดูสินค้า, ค้นหา
- **Categories** - ดูหมวดหมู่สินค้า
- **Cart** - จัดการตะกร้าสินค้า
- **Orders** - สร้างและติดตามคำสั่งซื้อ
- **Payments** - จัดการการชำระเงิน
- **Users** - จัดการข้อมูลผู้ใช้
- **Admin** - ฟีเจอร์สำหรับผู้ดูแลระบบ
- **Stats** - ดูสถิติต่างๆ

## 🔧 Configuration

ตัวแปรสภาพแวดล้อมที่สำคัญ:

```
DATABASE_URL=postgres://user:password@localhost:5432/nutchapholshop
JWT_SECRET=your-secret-key
SERVER_PORT=8080
ENVIRONMENT=development
```

## 📝 Scripts

- `go run cmd/api/main.go` - รัน API Server
- `go run cmd/migrate/main.go` - รัน Database Migrations
- `go test ./...` - รัน Unit Tests

## 🤝 มีส่วนร่วมในโปรเจกต์

ยินดีต้อนรับการมีส่วนร่วม! โปรดดูไฟล์ CONTRIBUTING.md สำหรับรายละเอียดเพิ่มเติม

## 📄 License

โปรเจกต์นี้ได้รับอนุญาตภายใต้ MIT License - ดูไฟล์ LICENSE สำหรับรายละเอียด

## 📧 ติดต่อ

สำหรับคำถามหรือข้อเสนอแนะ โปรดติดต่อทีมพัฒนา

---

**Note**: โปรเจกต์นี้ยังอยู่ในระหว่างพัฒนา อาจมีการเปลี่ยนแปลงในเร็วๆ นี้
