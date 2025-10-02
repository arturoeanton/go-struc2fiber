# MEJORAS.md

## 🎯 Objetivo: Escribir Menos, Hacer Más

Este documento propone mejoras para maximizar la productividad reduciendo el código boilerplate al mínimo, aprovechando generics, reflection y configuración declarativa.

## 1. 🚀 Auto-Registro de Rutas con Anotaciones

### Problema Actual
```go
// main.go - Mucho código repetitivo
app.Get("/internal_users", handlers.NewHandler[model.InternalUser]().GetAll)
app.Get("/internal_users/:id", handlers.NewHandler[model.InternalUser]().GetByID)
app.Post("/internal_users", handlers.NewHandler[model.InternalUser]().FxCreate(GetValidator("schemas/internal_users.yaml")))
// ... repetir para cada modelo
```

### Solución Propuesta
```go
// main.go - Una sola línea por modelo
app := fiber.New()
RegisterCRUD(app, model.InternalUser{}, "schemas/internal_users.yaml")
RegisterCRUD(app, model.Skill{}, "schemas/skill.yaml")
```

### Implementación
```go
// pkg/commons/crud_registry.go
func RegisterCRUD[T any](app *fiber.App, model T, validationPath string) {
    modelType := reflect.TypeOf(model)
    resourceName := strings.ToLower(modelType.Name())
    
    // Auto-genera todas las rutas CRUD
    handler := handlers.NewHandler[T]()
    validator := validator.NewStructValidator()
    validator.LoadSchemaFromFile(validationPath)
    
    app.Get("/"+resourceName, handler.GetAll)
    app.Get("/"+resourceName+"/:id", handler.GetByID)
    app.Post("/"+resourceName, handler.FxCreate(validator))
    app.Put("/"+resourceName+"/:id", handler.FxUpdate(validator))
    app.Delete("/"+resourceName+"/:id", handler.DeleteByID)
}
```

## 2. 📝 Validación YAML Mejorada con Auto-Generación

### Problema Actual
- Crear manualmente archivos YAML de validación
- Duplicación entre modelo Go y esquema YAML

### Solución Propuesta
```go
// Generar YAML desde tags del modelo
type User struct {
    ID       int64  `json:"id" gorm:"primaryKey"`
    Username string `json:"username" validate:"required,min=3,max=20,pattern=^[a-zA-Z0-9_]+$"`
    Email    string `json:"email" validate:"required,email"`
    Age      int    `json:"age" validate:"min=18,max=120"`
}

// Auto-generar schema YAML con comando
go run cmd/generate-schemas/main.go
```

### Herramienta de Generación
```go
// cmd/generate-schemas/main.go
func GenerateSchemaFromModel(model interface{}) string {
    schema := ValidationSchema{
        Name: reflect.TypeOf(model).Name(),
    }
    
    t := reflect.TypeOf(model)
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        validateTag := field.Tag.Get("validate")
        if validateTag != "" {
            rule := parseValidateTag(field.Name, validateTag)
            schema.Rules = append(schema.Rules, rule)
        }
    }
    
    return toYAML(schema)
}
```

## 3. 🔄 Migraciones Automáticas con Versionado

### Problema Actual
- No hay sistema de migraciones
- Cambios en modelos requieren intervención manual

### Solución Propuesta
```go
// pkg/migrations/auto_migrate.go
func AutoMigrate() {
    models := []interface{}{
        &model.InternalUser{},
        &model.Skill{},
        &model.Connect{},
        // Auto-detectar modelos con reflection
    }
    
    for _, m := range models {
        DB.AutoMigrate(m)
        GenerateValidationSchema(m) // Auto-generar validación
        RegisterHooks(m)            // Auto-registrar hooks
    }
}
```

## 4. 🎨 Configuración Declarativa con un Solo Archivo

### Problema Actual
- Configuración dispersa entre código y archivos YAML

### Solución Propuesta
```yaml
# config/api.yaml
models:
  - name: InternalUser
    path: /users
    validation:
      username:
        required: true
        min: 3
        max: 20
      email:
        type: email
        required: true
    hooks:
      before_create: HashPassword
      after_update: SendNotification
    
  - name: Skill
    path: /skills
    validation: auto  # Genera desde tags del modelo
    middleware:
      - AuthRequired
      - RateLimit
```

```go
// main.go - Todo el API en 5 líneas
config := LoadConfig("config/api.yaml")
app := fiber.New()
AutoRegisterFromConfig(app, config)
app.Listen(":3000")
```

## 5. 🔍 Búsqueda y Filtrado Automático

### Problema Actual
- No hay soporte para queries complejas sin escribir código

### Solución Propuesta
```go
// Automáticamente disponible en todos los endpoints
GET /users?filter[email]=*@gmail.com&sort=-created_at&page=2&limit=10
GET /users?search=john&fields=id,username,email
GET /users?where=age>18 AND status='active'
```

### Implementación
```go
// pkg/handlers/query_builder.go
func (h *Handler[T]) GetAll(c *fiber.Ctx) error {
    query := ParseQueryParams(c)
    
    items, total, err := h.service.GetWithQuery(query)
    
    return c.JSON(Response{
        Data: items,
        Meta: Meta{
            Total: total,
            Page: query.Page,
            Limit: query.Limit,
        },
    })
}
```

## 6. 🚦 Middleware Automático por Modelo

### Solución Propuesta
```go
// pkg/model/internal_user.go
type InternalUser struct {
    // ... campos
}

// Método opcional que define middlewares
func (InternalUser) Middlewares() []fiber.Handler {
    return []fiber.Handler{
        middleware.RequireAuth(),
        middleware.RateLimit(100),
        middleware.ValidateRole("admin"),
    }
}

// Se aplican automáticamente al registrar rutas
```

## 7. 📊 Generación Automática de Documentación OpenAPI

### Solución Propuesta
```go
// Auto-generar desde modelos y validaciones
go run cmd/generate-docs/main.go

// Produce: swagger.yaml con toda la API documentada
// Sirve automáticamente en /docs
```

## 8. 🎯 CLI para Scaffolding

### Solución Propuesta
```bash
# Generar modelo completo con una línea
go-struc2fiber generate model Product --fields "name:string,price:float,stock:int"

# Genera automáticamente:
# - pkg/model/product.go
# - schemas/product.yaml
# - Registra en main.go
# - Migración de DB
# - Tests básicos
```

## 9. 🔌 Hooks y Eventos Declarativos

### Solución Propuesta
```yaml
# schemas/user_hooks.yaml
model: User
hooks:
  before_create:
    - validate_email_domain
    - hash_password
  after_create:
    - send_welcome_email
    - create_audit_log
  before_update:
    - check_permissions
  after_delete:
    - cleanup_related_data
```

## 10. 🧪 Tests Automáticos

### Solución Propuesta
```go
// pkg/testing/crud_test.go
func TestCRUD[T any](t *testing.T, model T, testData T) {
    // Tests automáticos para:
    // - Create
    // - Read
    // - Update
    // - Delete
    // - Validación
    // - Permisos
}

// Uso: Una línea por modelo
TestCRUD(t, model.User{}, testUser)
```

## 📋 Resumen de Beneficios

| Mejora | Líneas Antes | Líneas Después | Reducción |
|--------|--------------|----------------|-----------|
| Registro de rutas | 5 por modelo | 1 por modelo | 80% |
| Validación | 20+ líneas YAML | 0 (auto-generado) | 100% |
| Configuración API | 50+ líneas | 10 líneas YAML | 80% |
| Tests CRUD | 100+ por modelo | 1 por modelo | 99% |
| Documentación | Manual | Auto-generada | 100% |

## 🎯 Objetivo Final

```go
// main.go completo
package main

import "github.com/arturoeanton/go-struc2fiber/pkg/auto"

func main() {
    auto.StartAPI("config/api.yaml", ":3000")
}
```

**¡TODO el CRUD, validación, documentación y tests en una sola línea!**

## 🚀 Próximos Pasos

1. **Fase 1**: Implementar auto-registro de rutas
2. **Fase 2**: Generación automática de schemas YAML
3. **Fase 3**: CLI para scaffolding
4. **Fase 4**: Configuración unificada YAML
5. **Fase 5**: Tests automáticos

Cada mejora es independiente y puede implementarse gradualmente.