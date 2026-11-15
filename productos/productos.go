package productos

import "fmt"

var codigos []string
var nombres []string
var precios []float64

func RegistrarProducto(codigo string, nombre string, precio float64) bool {
	if codigo == "" || nombre == "" || precio <= 0 {
		return false
	}

	for _, c := range codigos {
		if c == codigo {
			return false
		}
	}

	codigos = append(codigos, codigo)
	nombres = append(nombres, nombre)
	precios = append(precios, precio)
	return true
}

func ConsultarProducto(codigo string) string {
	for i, c := range codigos {
		if c == codigo {
			return fmt.Sprintf("%s - $%.2f", nombres[i], precios[i])
		}
	}
	return "Producto no encontrado"
}

func ActualizarProducto(codigo string, precio float64) bool {
	if precio <= 0 {
		return false
	}

	for i, c := range codigos {
		if c == codigo {
			precios[i] = precio
			return true
		}
	}
	return false
}

func EliminarProducto(codigo string) bool {
	for i, c := range codigos {
		if c == codigo {
			codigos = append(codigos[:i], codigos[i+1:]...)
			nombres = append(nombres[:i], nombres[i+1:]...)
			precios = append(precios[:i], precios[i+1:]...)
			return true
		}
	}
	return false
}
