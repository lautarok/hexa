![Hexa SVG logo](https://github.com/lautarok/hexa/blob/backend/assets/logo.svg?raw=true)

---

# Hexa
Plataforma de subastas en tiempo real con arquitectura hexagonal.

## Descripción
Hexa es una aplicación web que permite a usuarios autenticados crear y participar en subastas de productos, con soporte para pagos seguros, notificaciones en tiempo real y gestión de disputas.  
El proyecto está diseñado con **arquitectura hexagonal (Ports & Adapters)** para garantizar mantenibilidad, escalabilidad y separación de responsabilidades.

## Arquitectura
La aplicación sigue un enfoque modular:

- **Backend**
  - Lenguae: **Go**.
  - Framework principal: **Gin** (HTTP handler).
  - ORM: **Bun**.
  - Configuración: **Godotenv**.
  - Comunicación en tiempo real: **Gorilla Websocket**.
  - Base de datos: **PostgreSQL**.

- **Application Layer**
  - **Domain**: Entidades y lógica de negocio.
  - **Modules**: Servicios y casos de uso, controladores, DTOs, repositorios y puertos.
  - **Ports**: Repository, Env, Realtime.

- **Infrastructure**
  - Adaptadores para HTTP, repositorios y servicios externos.
  - Manejo de configuración y dependencias.

El diseño asegura que la lógica de negocio esté desacoplada de frameworks y librerías externas.

## Historias de Usuario
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

## Criterios Clave
- Ofertas y notificaciones en **tiempo real**.
- Manejo seguro de pagos con estados: disponible, congelado, completado, cancelado.
- **Bloqueos de base de datos (Pessimistic Locking)** para garantizar integridad en las ofertas.
- Moderación con reglas claras para disputas y penalizaciones.

## Documentación
Archivos de la documentación de Hexa

### Historias de usuario
[Descargar historias de usuario (en formato .docx)](https://github.com/lautarok/hexa/blob/backend/assets/Hexa%20-%20Historias%20de%20usuario.docx?raw=true)

### Diagrama de arquitectura
![Hexa SVG architecture diagram](https://github.com/lautarok/hexa/blob/backend/assets/Hexa%20-%20Diagrama%20de%20arquitectura.svg?raw=true)

### Diagrama de Entidad-Relación
![Hexa SVG entity-relation diagram](https://github.com/lautarok/hexa/blob/backend/assets/Hexa%20-%20Diagrama%20Entidad-Relación.svg?raw=true)
