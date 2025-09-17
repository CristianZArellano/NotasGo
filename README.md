
# NotasGo

Una API REST en Go para manejar notas usando **Gin** y **GORM** con SQLite como base de datos. La API permite crear, leer, actualizar y eliminar notas, y está documentada con **Swagger**.

---

## 🚀 Requisitos

- Go >= 1.25
- SQLite
- Git
- Dependencias de Go (gestionadas por `go.mod`)

---

## 📂 Estructura del proyecto

```

.
├── controllers      # Lógica de los endpoints
├── database         # Conexión y configuración de la base de datos
├── docs             # Documentación Swagger
├── models           # Modelos de datos
├── routes           # Rutas de la API
├── main.go          # Archivo principal
├── go.mod
├── go.sum
├── notas.db         # Base de datos SQLite
└── tmp              # Archivos temporales y logs

````

---

## ⚡ Instalación

1. Clona el repositorio:

```bash
git clone git@github.com:CristianZArellano/NotasGo.git
cd NotasGo
````

2. Instala dependencias:

```bash
go mod tidy
```

3. Ejecuta la API:

```bash
go run main.go
```

La API estará corriendo por defecto en `http://localhost:8080`.

---

## 📖 Endpoints principales

| Método | Ruta          | Descripción                      |
| ------ | ------------- | -------------------------------- |
| GET    | `/`           | Mensaje de bienvenida            |
| GET    | `/notes`      | Obtiene todas las notas          |
| GET    | `/notes/{id}` | Obtiene una nota por ID          |
| POST   | `/notes`      | Crea una nueva nota              |
| PUT    | `/notes/{id}` | Actualiza completamente una nota |
| PATCH  | `/notes/{id}` | Actualiza parcialmente una nota  |
| DELETE | `/notes/{id}` | Elimina una nota                 |

---

## 📌 Uso de Swagger

La documentación Swagger se encuentra en la carpeta `docs`. Para usarla:

1. Instala `swag`:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Genera la documentación:

```bash
swag init
```

3. Accede a la documentación en:

```
http://localhost:8080/swagger/index.html
```

---

## 🔧 Contribuciones

Si quieres contribuir:

1. Haz un fork del repositorio
2. Crea tu rama: `git checkout -b feature/nueva-funcionalidad`
3. Realiza tus cambios
4. Haz commit: `git commit -m "Agrega nueva funcionalidad"`
5. Haz push a tu rama
6. Crea un Pull Request

---

