# NotasGo API 2.0

Una **API REST refactorizada** en Go para gestión de notas y usuarios, construida con arquitectura moderna usando **Gin**, **GORM** y **SQLite**. Implementa patrones de **service layer**, validación robusta, y API versionada con documentación completa en **Swagger**.

---

## ✨ Características Principales

- 🏗️ **Arquitectura Limpia** - Service layer pattern con separación de responsabilidades
- 🔐 **Autenticación Segura** - Registro y login con hash bcrypt
- 👥 **Gestión de Usuarios** - CRUD completo con validación
- 📝 **Sistema de Notas** - Notas asociadas a usuarios con relaciones FK
- 🌐 **API Versionada** - Endpoints v1 con compatibilidad legacy
- ✅ **Validación Robusta** - DTOs con validación de entrada
- 📚 **Documentación Swagger** - API completamente documentada
- 🎯 **Respuestas Consistentes** - Formato estandarizado de respuestas

---

## 🚀 Requisitos

- **Go** >= 1.25
- **SQLite** (incluido)
- **Git**
- Dependencias de Go (gestionadas automáticamente por `go.mod`)

---

## 🏗️ Arquitectura del Proyecto

```
├── controllers/           # HTTP handlers con service layer
│   ├── users.go              # Controlador de usuarios
│   ├── notes.go              # Controlador de notas
│   └── home.go               # Dashboard
├── services/              # Lógica de negocio
│   ├── user_service.go       # Servicios de usuario
│   └── note_service.go       # Servicios de notas
├── models/                # Modelos y DTOs
│   ├── user.go               # Entidad usuario
│   ├── note.go               # Entidad nota
│   ├── requests.go           # DTOs de entrada
│   └── responses.go          # DTOs de salida
├── utils/                 # Utilidades
│   └── responses.go          # Helpers de respuesta HTTP
├── database/              # Capa de datos
│   └── database.go           # Conexión GORM
├── routes/                # Definición de rutas
│   └── routes.go             # Router principal
├── docs/                  # Documentación Swagger
├── static/                # Archivos estáticos
├── templates/             # Templates HTML
└── main.go                # Punto de entrada
```

---

## ⚡ Instalación y Uso

### 1. Clonar e Instalar

```bash
git clone git@github.com:CristianZArellano/NotasGo.git
cd NotasGo
go mod tidy
go mod vendor  # El proyecto usa vendor mode
```

### 2. Ejecutar la API

```bash
go run main.go
```

La API estará disponible en `http://localhost:8080`

### 3. Desarrollo con Hot Reload

```bash
air  # Requiere air: go install github.com/cosmtrek/air@latest
```

---

## 🌐 API Endpoints

### 🔐 Autenticación

| Método | Endpoint                | Descripción           |
|--------|------------------------|-----------------------|
| POST   | `/api/v1/auth/register` | Registro de usuario  |
| POST   | `/api/v1/auth/login`   | Login de usuario     |

### 👥 Usuarios

| Método | Endpoint              | Descripción                    |
|--------|-----------------------|--------------------------------|
| GET    | `/api/v1/users`       | Listar todos los usuarios      |
| GET    | `/api/v1/users/:id`   | Obtener usuario por ID         |
| PUT    | `/api/v1/users/:id`   | Actualizar usuario             |
| DELETE | `/api/v1/users/:id`   | Eliminar usuario y sus notas   |

### 📝 Notas

| Método | Endpoint                     | Descripción                    |
|--------|------------------------------|--------------------------------|
| GET    | `/api/v1/notes`             | Listar todas las notas         |
| GET    | `/api/v1/notes/:id`         | Obtener nota por ID            |
| POST   | `/api/v1/notes`             | Crear nueva nota               |
| PUT    | `/api/v1/notes/:id`         | Actualizar nota completa       |
| PATCH  | `/api/v1/notes/:id`         | Actualización parcial          |
| DELETE | `/api/v1/notes/:id`         | Eliminar nota                  |

### 🔗 Relaciones

| Método | Endpoint                        | Descripción              |
|--------|---------------------------------|--------------------------|
| GET    | `/api/v1/user/:user_id/notes`  | Obtener notas de usuario |

### 🔄 Compatibilidad Legacy

Todos los endpoints están disponibles también sin el prefijo `/api/v1/` para compatibilidad con versiones anteriores.

---

## 📚 Documentación Swagger

### Generar Documentación

```bash
# Instalar swag
go install github.com/swaggo/swag/cmd/swag@latest

# Generar docs
swag init
```

### Acceder a Swagger UI

```
http://localhost:8080/swagger/index.html
```

---

## 🔧 Ejemplos de Uso

### Registro de Usuario

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Crear Nota

```bash
curl -X POST http://localhost:8080/api/v1/notes \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Mi Primera Nota",
    "content": "Contenido de la nota",
    "user_id": 1
  }'
```

### Obtener Notas de Usuario

```bash
curl http://localhost:8080/api/v1/user/1/notes
```

---

## 🔒 Características de Seguridad

- **Hashing de Contraseñas** - bcrypt con salt automático
- **Validación de Entrada** - DTOs con validación robusta
- **Sanitización de Respuestas** - Exclusión de datos sensibles
- **Validación de Unicidad** - Email y username únicos
- **Relaciones Seguras** - Foreign keys con validación

---

## 📊 Modelos de Datos

### Usuario

```go
{
  "id": 1,
  "username": "johndoe",
  "email": "john@example.com",
  "role": "user",
  "status": "activo",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

### Nota

```go
{
  "id": 1,
  "title": "Mi Nota",
  "content": "Contenido de la nota",
  "user_id": 1,
  "user": { /* objeto usuario */ },
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

---

## 🛠️ Desarrollo

### Comandos Útiles

```bash
# Construir aplicación
go build -o ./tmp/main .

# Ejecutar tests (cuando estén implementados)
go test ./...

# Linting (si está configurado)
golangci-lint run

# Actualizar vendor
go mod vendor
```

### Estructura de Respuestas

#### Respuesta Exitosa
```json
{
  "success": true,
  "message": "Operación exitosa",
  "data": { /* datos solicitados */ }
}
```

#### Respuesta de Error
```json
{
  "success": false,
  "message": "Error en la operación",
  "error": "Detalles del error"
}
```

---

## 🚀 Patrones Implementados

### Service Layer Pattern
- **Controllers** - Manejo HTTP y validación
- **Services** - Lógica de negocio y reglas
- **Models** - Entidades y DTOs
- **Utils** - Funciones auxiliares

### Beneficios Arquitectónicos
- ✅ **Separación de Responsabilidades**
- ✅ **Código Testeable y Modular**
- ✅ **Fácil Mantenimiento**
- ✅ **Escalabilidad Horizontal**
- ✅ **Reutilización de Servicios**

---

## 📈 Próximas Mejoras

- [ ] **Testing Suite** - Unit e integration tests
- [ ] **JWT Authentication** - Tokens para sesiones
- [ ] **Rate Limiting** - Protección contra spam
- [ ] **Logging Estructurado** - Logs con formato JSON
- [ ] **Paginación** - Para listas grandes
- [ ] **Cache Layer** - Redis para performance
- [ ] **Docker Support** - Containerización
- [ ] **CI/CD Pipeline** - Automatización

---

## 🔧 Contribuciones

### Cómo Contribuir

1. **Fork** del repositorio
2. **Crear rama**: `git checkout -b feature/nueva-funcionalidad`
3. **Seguir patrones** establecidos (service layer, DTOs, etc.)
4. **Añadir tests** para nueva funcionalidad
5. **Commit**: `git commit -m "feat: descripción del cambio"`
6. **Push**: `git push origin feature/nueva-funcionalidad`
7. **Pull Request** con descripción detallada

### Estándares de Código

- Seguir **Go conventions**
- Usar **service layer** para lógica de negocio
- Implementar **DTOs** para requests/responses
- Añadir **documentación Swagger**
- Mantener **compatibilidad de API**

---

## 📄 Licencia

Este proyecto está bajo la licencia MIT. Ver el archivo `LICENSE` para más detalles.

---

## 📞 Contacto

**Desarrollador:** Cristian Arellano  
**Email:** zaidarellano@gmail.com
**GitHub:** [@CristianZArellano](https://github.com/CristianZArellano)

---

## 🙏 Agradecimientos

- **Gin Framework** - HTTP router rápido y minimalista
- **GORM** - ORM potente y fácil de usar
- **Swaggo** - Generación automática de documentación
- **Comunidad Go** - Por las excelentes librerías y herramientas

---

**¡Construido con ❤️ usando Go y arquitectura moderna!**