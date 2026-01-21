# Stock Ratings Service

Stock Ratings Service es una aplicación full-stack diseñada para gestionar, almacenar y consultar recomendaciones de analistas financieros, precios objetivo y métricas clave de acciones del mercado.

El sistema integra información desde servicios externos, calcula indicadores financieros relevantes como el *upside* y expone los datos a través de una API REST consumida por un frontend web moderno.

El foco principal del proyecto es la **calidad arquitectónica**, la **separación de responsabilidades** y la **testabilidad**.

---

## Stack tecnológico

### Backend
- **Go 1.25.5**
- **Echo v4** (API REST)
- **CockroachDB** (compatible con PostgreSQL)
- **pgx v5** (driver de base de datos)
- **sqlc** (queries tipadas)
- **testify** (testing)
- **mockery** (mocks para pruebas unitarias)

### Frontend
- **Vue 3**
- **TypeScript**
- **Pinia** (manejo de estado)
- **Tailwind CSS** (estilos)
- **Vite** (build tool)

---

## Visión general de la arquitectura

La aplicación sigue principios de **Domain-Driven Design (DDD)** combinados con **CQRS**, **Vertical Slice Architecture** y el patrón **Unit of Work**.

El sistema se divide conceptualmente en tres partes:

- Un **servicio de sincronización** encargado de traer y procesar datos desde servicios externos.
- Un **backend API REST** enfocado en exponer información ya procesada y validada.
- Un **frontend web** que consume la API y presenta los datos al usuario final.

Cada parte cumple una responsabilidad clara y se comunica únicamente mediante contratos bien definidos.

---

## Backend

El backend está desarrollado en Go y organizado por *features*, no por capas técnicas genéricas.  
Esto facilita el crecimiento del sistema sin generar dependencias innecesarias entre módulos.

### Estructura general

- **cmd/api**  
  Punto de entrada del servidor HTTP.  
  Aquí se configuran:
  - Rutas
  - Middlewares
  - Inyección de dependencias
  - Arranque del servidor

- **cmd/sync**  
  Servicio independiente de sincronización de datos.  
  Su responsabilidad es:
  - Consumir un servicio externo de stock ratings
  - Consultar Yahoo Finance para obtener el precio actual
  - Consumir el handler que contiene la logica necesaria para procesar y posteriromente persistir la informacion
    en la base de datos local

- **internal/features**  
  Contiene la lógica del negocio organizada por dominio.

---

## Dominio y patrones aplicados

### Domain-Driven Design (DDD)

El dominio de *stock ratings* está modelado explícitamente mediante:

- **Entidades enriquecidas**, responsables de mantener su estado válido.
- **Objetos de valor**, utilizados para representar conceptos como precios y evitar valores inconsistentes.
- **Validaciones de dominio**, ejecutadas antes de persistir o modificar datos.

La lógica de negocio vive en el dominio, no en los controladores ni en la infraestructura.

---

### CQRS (Command Query Responsibility Segregation)

Las operaciones están claramente separadas entre:

- **Commands**  
  Encargados de modificar el estado del sistema.

- **Queries**  
  Encargadas únicamente de lectura, devolviendo DTOs listos para consumo por la API.

Esto mejora la claridad del código y permite optimizar lectura y escritura de forma independiente.

---

## Vertical Slice Architecture

Cada feature se implementa como un **slice vertical completo**, que encapsula todo lo necesario para ejecutar un caso de uso específico.

### Estructura por feature

- **app**
  - commands
  - queries
  - handlers
  - dtos
  - interfaces (contratos de entrada)

- **domain**
  - entities
  - value objects
  - validation
  - pagination
  - repositories (interfaces)
  - unit of work (interfaces)

- **infra**
  - cockroach: connection, transactions, unit of work implementation
  - sqlc: repositorios generados por sqlc

- **interface**
  - API REST
  - requests
  - controllers

---

## Validaciones y guardas de request

Antes de ejecutar cualquier caso de uso, el sistema aplica **guardas de validación** a nivel de aplicación:

- Validación de campos requeridos
- Validación de tipos y rangos numéricos
- Normalización de valores de entrada
- Prevención de requests inválidas antes de llegar al dominio

Estas guardas:
- Protegen el dominio de estados inválidos
- Simplifican los handlers
- Mejoran la seguridad y robustez del sistema

Las validaciones de negocio más complejas permanecen exclusivamente en el dominio.

---

### Unit of Work

El patrón Unit of Work se utiliza para manejar transacciones de base de datos de forma segura:

- La transacción se inicia en la capa de infraestructura.
- Los handlers trabajan sobre una abstracción.
- Se garantiza commit o rollback según el resultado de la operación.

Esto evita estados inconsistentes y simplifica el manejo transaccional.

---

### Factories e Inyección de Dependencias

- Las entidades se crean mediante **factories**, centralizando reglas y validaciones de dominio.
- Las dependencias se inyectan desde el `main`.
- Esto permite pruebas unitarias aisladas y uso extensivo de mocks.

---

## Persistencia y base de datos

La base de datos utilizada es **CockroachDB**, elegida por su compatibilidad con PostgreSQL y su capacidad de escalar.

- Acceso mediante **pgx**
- Queries tipadas con **sqlc**
- Migraciones SQL versionadas
- Configuración vía variables de entorno usando **godotenv**

---

## Sincronización de datos externos

El sistema integra dos fuentes externas:

- **Servicio de Stock Ratings**  
  Proporciona recomendaciones, acciones y precios objetivo.

- **Yahoo Finance**  
  Proporciona el precio actual de mercado de cada acción.

El servicio `cmd/sync`:
- Obtiene los datos
- Normaliza la información
- Consume el handler que contiene la logica necesaria para procesar, calcular upside y posteriromente persistir la informacion en la base de datos local

Esto desacopla completamente la ingestión de datos del API principal.

---

## API REST

El backend expone endpoints REST para:

- Listar acciones con filtros por texto, rango de precios y upside mínimo
- Mostrar la mejor recomendacion de compra
- Ordenar y paginar resultados
---

## Frontend

El frontend está desarrollado en **Vue 3 con TypeScript**.

- **Pinia** se utiliza para el manejo de estado global.
- **Tailwind CSS** permite construir una interfaz consistente y mantenible.
- Consume la API REST del backend.
- Presenta un listado de acciones con:
  - Búsqueda por ticker o compañía
  - Filtros financieros
  - Visualización de la recomendacion de compra a la accion mas favorable

El frontend está desacoplado del backend y se comunica exclusivamente por HTTP.

---

## Testing

El proyecto cuenta con pruebas unitarias enfocadas en:

- Handlers de commands y querie
- Controladores
- Guardas
- Helpers
- Coneccion a BD y transacciones
- Validaciones de dominio
- Lógica de negocio crítica

Herramientas utilizadas:
- **testify** para assertions
- **mockery** para generación de mocks

Para ejecutar todos los tests:

```bash
go test ./...
