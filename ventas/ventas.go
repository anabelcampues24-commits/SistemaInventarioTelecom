package ventas

import (
	"fmt"
	"inventariotelecom/inventario"
)

var ventasRegistradas []string

// RegistrarVenta registra una venta y descuenta el stock.
func RegistrarVenta(codigo string, cantidad int) bool {
	if !ValidarVenta(codigo, cantidad) {
		return false
	}

	if inventario.RegistrarSalida(codigo, cantidad) {
		detalle := fmt.Sprintf(
			"Producto: %s, Cantidad: %d",
			codigo, cantidad,
		)

		ventasRegistradas = append(ventasRegistradas, detalle)
		return true
	}

	return false
}

// CalcularTotal calcula el valor total de una venta.
func CalcularTotal(precio float64, cantidad int) float64 {
	if precio <= 0 || cantidad <= 0 {
		return 0
	}

	return precio * float64(cantidad)
}

// ConsultarVenta consulta una venta registrada.
func ConsultarVenta(idVenta int) string {
	if idVenta < 1 || idVenta > len(ventasRegistradas) {
		return "Venta no encontrada"
	}

	return ventasRegistradas[idVenta-1]
}

// ValidarVenta verifica que exista stock suficiente.
func ValidarVenta(codigo string, cantidad int) bool {
	return inventario.VerificarStock(codigo, cantidad)
}

// TotalVentasRegistradas cuenta las ventas realizadas.
func TotalVentasRegistradas() int {
	return len(ventasRegistradas)
}
