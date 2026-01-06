# Hexa 🧩
Plataforma de subastas en tiempo real con arquitectura hexagonal.

## 🚀 Descripción
Hexa es una aplicación web que permite a usuarios autenticados crear y participar en subastas de productos, con soporte para pagos seguros, notificaciones en tiempo real y gestión de disputas.  
El proyecto está diseñado con **arquitectura hexagonal (Ports & Adapters)** para garantizar mantenibilidad, escalabilidad y separación de responsabilidades.

## 📐 Arquitectura
La aplicación sigue un enfoque modular:

- **Backend**
  - Framework principal: **Gin** (HTTP handler).
  - Runtime: **Bun**.
  - Configuración: **Godotenv**.
  - Comunicación en tiempo real: **Gorilla Websocket**.

- **Application Layer**
  - **Domain**: Entidades y lógica de negocio.
  - **Modules**: Servicios y casos de uso, DTOs, Ports.
  - **Ports**: Repository, Env, Realtime.

- **Infrastructure**
  - Adaptadores para HTTP, repositorios y servicios externos.
  - Manejo de configuración y dependencias.

El diseño asegura que la lógica de negocio esté desacoplada de frameworks y librerías externas.

## 📋 Historias de Usuario
Principales funcionalidades definidas:

- **Exploración**
  - Listado de subastas disponibles con filtros, buscador y paginación.
  - Vista individual de subasta con historial de ofertas en tiempo real.

- **Autenticación**
  - Registro de usuarios (con SSO opcional vía Google).
  - Inicio de sesión con usuario/contraseña o SSO.

- **Notificaciones**
  - Bandeja de notificaciones en tiempo real, con scroll infinito y componente reutilizable.

- **Pagos**
  - Vinculación de métodos de pago (tarjetas).
  - Depósitos y retiros de dinero.
  - Gestión de pagos congelados, completados y cancelados.

- **Subastas**
  - Creación de nuevas subastas con validaciones de producto, precio inicial y garantía.
  - Participación en subastas con bloqueo pesimista para evitar inconsistencias.
  - Confirmación de producto recibido o inicio de disputa.
  - Administración de subastas por moderadores/administradores.

- **Perfil de Usuario**
  - Información personal, métodos de pago vinculados y métricas.
  - Historial de transacciones con scroll infinito.

## ✅ Criterios Clave
- Ofertas y notificaciones en **tiempo real**.
- Manejo seguro de pagos con estados: disponible, congelado, completado, cancelado.
- **Bloqueos de base de datos (Pessimistic Locking)** para garantizar integridad en las ofertas.
- Moderación con reglas claras para disputas y penalizaciones.

## 🛠️ Tecnologías
- **Go** con Gin y Gorilla Websocket.
- **Bun** como runtime.
- **Arquitectura Hexagonal** (Ports & Adapters).
- **Godotenv** para configuración.
- Base de datos relacional (a definir, compatible con Bun ORM).

## 📂 Estructura del Proyecto

```
hexa/
├── backend/                 # Entry points y configuración principal
│   ├── main.go                            # Inicialización del servidor
│   ├── router.go                        # Definición de rutas HTTP (Gin)
│   └── websocket.go                  # Configuración de Gorilla Websocket
│
├── domain/                  # Núcleo de negocio (independiente de frameworks)
│   ├── entities/            # Entidades principales (Usuario, Subasta, Oferta, Pago)
│   └── value_objects/       # Objetos de valor y reglas de negocio
│
├── modules/                 # Casos de uso y servicios
│   ├── services/            # Lógica de aplicación (crear subasta, ofertar, confirmar)
│   ├── usecases/            # Coordinación de procesos de negocio
│   ├── dtos/                # Data Transfer Objects (entrada/salida)
│   └── ports/               # Interfaces hacia infraestructura
│
├── infrastructure/          # Adaptadores externos
│   ├── http/                # Handlers y controladores (Gin)
│   ├── repository/          # Implementaciones de persistencia (DB)
│   ├── env/                 # Configuración (Godotenv)
│   ├── realtime/            # Implementación de sockets (Gorilla Websocket)
│   └── payment/             # Integración con pasarela de pagos
│
├── config/                  # Archivos de configuración
│   ├── app.env                            # Variables de entorno
│   └── bun.toml                          # Configuración de Bun
│
├── docs/                    # Documentación técnica
│   ├── ERD.svg                            # Diagrama entidad-relación
│   ├── architecture.svg          # Diagrama de arquitectura
│   └── README.md                        # Documentación principal
│
├── scripts/                 # Scripts de utilidad (migraciones, seeds, etc.)
│
└── tests/                   # Pruebas unitarias e integración
├── domain/              # Tests de entidades y reglas de negocio
├── modules/             # Tests de casos de uso
└── infrastructure/      # Tests de adaptadores
```

---

### 🔑 Notas
- **domain/** es el corazón: no depende de nada externo.  
- **modules/** orquesta la lógica de aplicación y conecta con los **ports**.  
- **infrastructure/** implementa los adaptadores concretos (DB, HTTP, pagos, realtime).  
- **backend/** es el punto de entrada, donde se inicializa todo.  
- **tests/** mantiene la calidad y asegura integridad en cada capa.  

---