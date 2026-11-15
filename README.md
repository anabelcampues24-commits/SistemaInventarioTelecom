
# Sistema de Gestión de Inventario y Bodega para Telecomunicaciones

## Descripción
Proyecto académico desarrollado en Go para gestionar
equipos de telecomunicaciones destinados a su comercialización
a proveedores.

## Objetivo
Desarrollar un sistema modular que permita registrar equipos,
controlar inventarios, gestionar proveedores, registrar ventas
y consultar información básica del sistema.

## Módulos
1. Productos
2. Inventarios
3. Proveedores
4. Ventas
5. Reportes
6. Usuarios y Roles

## Tecnologías
- Lenguaje de programación: Go
- Entorno de desarrollo: Visual Studio Code
- Arquitectura: Modular
- Interfaz: Consola

## Funcionalidades
- Registro y consulta de productos.
- Control de entradas y salidas de inventario.
- Administración de proveedores.
- Registro y cálculo de ventas.
- Consulta de reportes.
- Registro de usuarios y asignación de roles.

## Ejecución
Abrir una terminal en la carpeta del proyecto y ejecutar:

    go run .

## Verificación
Para comprobar que todos los paquetes compilan:

    go build ./...

## Alcance
El sistema es un prototipo académico. Los registros se
almacenan temporalmente en memoria y no utiliza una base
de datos permanente.
