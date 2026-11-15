package reportes

import (
	"fmt"

	"inventariotelecom/inventario"
	"inventariotelecom/ventas"
)

// ReporteInventario muestra las existencias disponibles.
func ReporteInventario() string {
	stock := inventario.ConsultarStock("ONT001")

	return fmt.Sprintf(
		"Producto ONT001 - Stock disponible: %d",
		stock,
	)
}

// ReporteStockBajo identifica productos con pocas existencias.
func ReporteStockBajo(minimo int) string {
	stock := inventario.ConsultarStock("ONT001")

	if stock < minimo {
		return "ALERTA: Stock bajo del producto ONT001"
	}

	return "Stock suficiente del producto ONT001"
}

// ReporteVentas muestra el numero de ventas registradas.
func ReporteVentas() string {
	total := ventas.TotalVentasRegistradas()

	return fmt.Sprintf(
		"Total de ventas registradas: %d",
		total,
	)
}

// ReporteMovimientos presenta el estado del modulo.
func ReporteMovimientos() string {
	return "Modulo de movimientos preparado para futuras etapas"
}
