package main

import (
	"fmt"

	"inventariotelecom/inventario"
	"inventariotelecom/productos"
	"inventariotelecom/proveedores"
	"inventariotelecom/reportes"
	"inventariotelecom/usuarios"
	"inventariotelecom/ventas"
)

func main() {

	fmt.Println("==========================================")
	fmt.Println(" SISTEMA DE GESTION DE INVENTARIO Y BODEGA")
	fmt.Println(" EQUIPOS DE TELECOMUNICACIONES")
	fmt.Println("==========================================")

	// MODULO 1: PRODUCTOS
	fmt.Println("\n1. REGISTRO DE PRODUCTOS")

	registrado := productos.RegistrarProducto(
		"ONT001", "ONT Huawei", 45.50,
	)

	fmt.Println("Producto registrado:", registrado)

	fmt.Println("Consulta:",
		productos.ConsultarProducto("ONT001"))

	// MODULO 2: INVENTARIO
	fmt.Println("\n2. CONTROL DE INVENTARIO")

	entrada := inventario.RegistrarEntrada("ONT001", 20)
	fmt.Println("Entrada registrada:", entrada)

	fmt.Println("Stock inicial:",
		inventario.ConsultarStock("ONT001"))

	salida := inventario.RegistrarSalida("ONT001", 5)
	fmt.Println("Salida registrada:", salida)

	fmt.Println("Stock disponible:",
		inventario.ConsultarStock("ONT001"))

	// MODULO 3: PROVEEDORES
	fmt.Println("\n3. GESTION DE PROVEEDORES")

	proveedor := proveedores.RegistrarProveedor(
		"1790012345001",
		"Proveedor Telecom",
		"0991234567",
	)

	fmt.Println("Proveedor registrado:", proveedor)

	fmt.Println("Consulta:",
		proveedores.ConsultarProveedor("1790012345001"))

	actualizado := proveedores.ActualizarProveedor(
		"1790012345001", "0987654321",
	)

	fmt.Println("Telefono actualizado:", actualizado)

	fmt.Println("Datos actualizados:",
		proveedores.ConsultarProveedor("1790012345001"))

	// MODULO 4: VENTAS
	fmt.Println("\n4. GESTION DE VENTAS")

	disponible := ventas.ValidarVenta("ONT001", 2)
	fmt.Println("Stock suficiente:", disponible)

	venta := ventas.RegistrarVenta("ONT001", 2)
	fmt.Println("Venta registrada:", venta)

	total := ventas.CalcularTotal(45.50, 2)
	fmt.Printf("Total de la venta: $%.2f\n", total)

	fmt.Println("Detalle:",
		ventas.ConsultarVenta(1))

	fmt.Println("Stock despues de la venta:",
		inventario.ConsultarStock("ONT001"))

	// MODULO 5: REPORTES
	fmt.Println("\n5. REPORTES DEL SISTEMA")

	fmt.Println(reportes.ReporteInventario())

	fmt.Println(reportes.ReporteStockBajo(15))

	fmt.Println(reportes.ReporteVentas())

	fmt.Println(reportes.ReporteMovimientos())

	// MODULO 6: USUARIOS Y ROLES
	fmt.Println("\n6. GESTION DE USUARIOS Y ROLES")

	registradoUsuario := usuarios.RegistrarUsuario(
		"Cristina", "Administrador",
	)

	fmt.Println("Usuario registrado:", registradoUsuario)

	fmt.Println("Consulta:",
		usuarios.ConsultarUsuario("Cristina"))

	fmt.Println("Rol actual:",
		usuarios.ConsultarRol("Cristina"))

	actualizadoRol := usuarios.ActualizarRol(
		"Cristina", "Operador",
	)

	fmt.Println("Rol actualizado:", actualizadoRol)

	fmt.Println("Nuevo rol:",
		usuarios.ConsultarRol("Cristina"))

	fmt.Println("\n==========================================")
	fmt.Println(" EJECUCION DEL SISTEMA FINALIZADA")
	fmt.Println("==========================================")
}
