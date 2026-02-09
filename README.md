# 🎓 Student API with Gin Framework

REST API สำหรับจัดการข้อมูลนักศึกษา พัฒนาด้วย Go, Gin Framework และ SQLite

## ✨ Features

- ✅ Get all students (GET)
- ✅ Get student by ID (GET)
- ✅ Create new student (POST)
- ✅ Update student (PUT)
- ✅ Delete student (DELETE)

## 🏗️ Architecture

โปรเจคใช้ **Layered Architecture** แบ่งออกเป็น 4 ชั้น:

```
┌─────────────────────┐
│   Handler Layer     │  ← HTTP Request/Response
├─────────────────────┤
│   Service Layer     │  ← Business Logic & Validation
├─────────────────────┤
│  Repository Layer   │  ← Database Operations
├─────────────────────┤
│    Model Layer      │  ← Data Structures
└─────────────────────┘
```

## 📁 Project Structure

```
go-api-gin-lab/
├── main.go                    # Entry point & Routes
├── config/
│   └── database.go            # Database configuration
├── models/
│   └── student.go             # Student data model
├── repositories/
│   └── student_repository.go  # Database operations
├── services/
│   └── student_service.go     # Business logic
├── handlers/
│   └── student_handler.go     # HTTP handlers
├── go.mod                     # Go modules
├── go.sum                     # Dependencies checksum
└── students.db                # SQLite database (auto-generated)
```

## 🚀 How to Run

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/tha15thai/go-api-gin-lab
cd go-api-gin-lab
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run main.go
```

Server will start at `http://localhost:8080`

## 📡 Testing API Endpoints with Postman

### การเตรียมความพร้อม

1. **ติดตั้ง Postman**: ดาวน์โหลดจาก https://www.postman.com/downloads/
2. **ตั้งค่า Base URL**: `http://localhost:8080`
3. **ตั้งค่า Headers** (สำหรับ POST และ PUT):
   - Key: `Content-Type`
   - Value: `application/json`

---

### 1. Get All Students ✅
**ดึงข้อมูลนักศึกษาทั้งหมด**

```http
GET /students
```

**ขั้นตอนใน Postman:**
1. เลือก Method: `GET`
2. ใส่ URL: `http://localhost:8080/students`
3. กด `Send`

**Response ที่คาดหวัง:**
```json
[
  {
    "id": "6609650111",
    "name": "ยอดนักสืบจิ๋ว",
    "major": "Computer Science",
    "gpa": 4
  },
  {
    "id": "6609650222",
    "name": "คุโด้ ชินอิจิ",
    "major": "Computer Science",
    "gpa": 3.9
  }
]
```

---

### 2. Get Student by ID ✅
**ดึงข้อมูลนักศึกษาด้วย ID**

```http
GET /students/:id
```

**ขั้นตอนใน Postman:**
1. เลือก Method: `GET`
2. ใส่ URL: `http://localhost:8080/students/6609650222`
3. กด `Send`

**Response สำเร็จ (200 OK):**
```json
{
  "id": "6609650222",
  "name": "คุโด้ ชินอิจิ",
  "major": "Computer Science",
  "gpa": 3.9
}
```

**Error Response (404 Not Found):**
```json
{
  "error": "Student not found"
}
```

---

### 3. Create Student ✅
**เพิ่มข้อมูลนักศึกษาใหม่**

```http
POST /students
```

**ขั้นตอนใน Postman:**
1. เลือก Method: `POST`
2. ใส่ URL: `http://localhost:8080/students`
3. ไปที่แท็บ `Headers`
   - Key: `Content-Type`
   - Value: `application/json`
4. ไปที่แท็บ `Body`
   - เลือก `raw`
   - เลือกประเภท `JSON`
5. ใส่ JSON ด้านล่าง
6. กด `Send`

**Request Body:**
```json
{
  "id": "6609650444",
  "name": "โมริ รัน",
  "major": "Karate",
  "gpa": 4
}
```

**Response สำเร็จ (201 Created):**
```json
{
  "id": "6609650444",
  "name": "โมริ รัน",
  "major": "Karate",
  "gpa": 4
}
```

**Validation Errors (400 Bad Request):**

กรณี ID ว่าง:
```json
{
  "error": "id must not be empty"
}
```

กรณี Name ว่าง:
```json
{
  "error": "name must not be empty"
}
```

กรณี GPA ไม่อยู่ในช่วง 0-4:
```json
{
  "error": "gpa must be between 0.00 and 4.00"
}
```

---

### 4. Update Student ✅
**แก้ไขข้อมูลนักศึกษาที่มีอยู่แล้ว**

```http
PUT /students/:id
```

**ขั้นตอนใน Postman:**
1. เลือก Method: `PUT`
2. ใส่ URL: `http://localhost:8080/students/6609650333`
3. ไปที่แท็บ `Headers`
   - Key: `Content-Type`
   - Value: `application/json`
4. ไปที่แท็บ `Body`
   - เลือก `raw`
   - เลือกประเภท `JSON`
5. ใส่ JSON ด้านล่าง (ข้อมูลที่ต้องการอัปเดต)
6. กด `Send`

**Request Body:**
```json
{
  "id": "6609650333",
  "name": "คุโรบะ ไคโตะ",
  "major": "Jewelry",
  "gpa": 4.0
}
```

**Response สำเร็จ (200 OK):**
```json
{
  "id": "6609650333",
  "name": "คุโรบะ ไคโตะ",
  "major": "Jewelry",
  "gpa": 4.0
}
```

**Error Response (404 Not Found):**
```json
{
  "error": "Student not found"
}
```

**Validation Errors (400 Bad Request):**
- เหมือนกับ POST (ตรวจสอบ ID, Name, GPA)

---

### 5. Delete Student ✅
**ลบข้อมูลนักศึกษา (ต้องแน่ใจว่ามี ID นี้อยู่ในฐานข้อมูล)**

```http
DELETE /students/:id
```

**ขั้นตอนใน Postman:**
1. เลือก Method: `DELETE`
2. ใส่ URL: `http://localhost:8080/students/6609650333`
3. กด `Send`

**Response สำเร็จ (204 No Content):**
- ไม่มี Response Body
- HTTP Status: `204 No Content`
- ดูสถานะได้ที่มุมล่างขวาของ Postman

**Error Response (404 Not Found):**
```json
{
  "error": "Student not found"
}
```

---

### 6. Test Validation Errors ✅
**ทดสอบการตรวจสอบข้อมูล**

#### 6.1 ทดสอบ ID ว่าง

**ขั้นตอนใน Postman:**
1. เลือก Method: `POST`
2. URL: `http://localhost:8080/students`
3. Headers: `Content-Type: application/json`
4. Body (raw JSON):

```json
{
  "id": "",
  "name": "Test Student",
  "major": "Test",
  "gpa": 3.0
}
```

**Expected Response (400 Bad Request):**
```json
{
  "error": "id must not be empty"
}
```

---

#### 6.2 ทดสอบ Name ว่าง

**Body:**
```json
{
  "id": "6609650999",
  "name": "",
  "major": "Test",
  "gpa": 3.0
}
```

**Expected Response (400 Bad Request):**
```json
{
  "error": "name must not be empty"
}
```

---

#### 6.3 ทดสอบ GPA > 4.00

**Body:**
```json
{
  "id": "6609650999",
  "name": "Test Student",
  "major": "Test",
  "gpa": 5.0
}
```

**Expected Response (400 Bad Request):**
```json
{
  "error": "gpa must be between 0.00 and 4.00"
}
```

---

#### 6.4 ทดสอบ GPA < 0.00

**Body:**
```json
{
  "id": "6609650999",
  "name": "Test Student",
  "major": "Test",
  "gpa": -1.0
}
```

**Expected Response (400 Bad Request):**
```json
{
  "error": "gpa must be between 0.00 and 4.00"
}
```

---

### 7. Test 404 Errors ✅
**ทดสอบเมื่อไม่พบข้อมูล**

#### 7.1 GET Student ที่ไม่มี

**ขั้นตอนใน Postman:**
1. Method: `GET`
2. URL: `http://localhost:8080/students/999999`
3. กด `Send`

**Expected Response (404 Not Found):**
```json
{
  "error": "Student not found"
}
```

---

#### 7.2 UPDATE Student ที่ไม่มี

**ขั้นตอนใน Postman:**
1. Method: `PUT`
2. URL: `http://localhost:8080/students/999999`
3. Headers: `Content-Type: application/json`
4. Body:

```json
{
  "id": "999999",
  "name": "Not Exist",
  "major": "Test",
  "gpa": 3.0
}
```

**Expected Response (404 Not Found):**
```json
{
  "error": "Student not found"
}
```

---

#### 7.3 DELETE Student ที่ไม่มี

**ขั้นตอนใน Postman:**
1. Method: `DELETE`
2. URL: `http://localhost:8080/students/999999`
3. กด `Send`

**Expected Response (404 Not Found):**
```json
{
  "error": "Student not found"
}
```
---

## ✅ Validation Rules

ระบบจะตรวจสอบข้อมูลก่อนบันทึก:

| Field | Rule | Example Error |
|-------|------|---------------|
| `id` | ต้องไม่ว่าง | `"id must not be empty"` |
| `name` | ต้องไม่ว่าง | `"name must not be empty"` |
| `gpa` | ต้องอยู่ระหว่าง 0.00 - 4.00 | `"gpa must be between 0.00 and 4.00"` |

---

## 🎯 HTTP Status Codes

| Code | Meaning | When |
|------|---------|------|
| 200 | OK | GET, PUT สำเร็จ |
| 201 | Created | POST สำเร็จ |
| 204 | No Content | DELETE สำเร็จ |
| 400 | Bad Request | ข้อมูลไม่ถูกต้อง, validation ผิด |
| 404 | Not Found | ไม่พบข้อมูลที่ต้องการ |
| 500 | Internal Server Error | เกิดข้อผิดพลาดในระบบ |

---

## 💡 What I Learned

1. **REST API Design** - ออกแบบ API ตามมาตรฐาน RESTful
2. **Layered Architecture** - แยก logic ออกเป็นชั้นต่างๆ อย่างชัดเจน
3. **Input Validation** - ตรวจสอบข้อมูลก่อนบันทึกลงฐานข้อมูล
4. **Error Handling** - จัดการ error และส่ง HTTP status code ที่เหมาะสม
5. **Database Operations** - UPDATE และ DELETE ใน SQLite

---

## 👨‍💻 Author

**Student ID**: 6609650111  
**Student Name**: Thawanhathai T.  
**Course**: CS367 WEB SERVICE DEVELOPMENT CONCEPTS

---

**Last Updated**: February 2026