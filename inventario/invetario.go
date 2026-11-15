package inventario

var codigos []string
var cantidades []int

// RegistrarEntrada aumenta las existencias.
func RegistrarEntrada(codigo string, cantidad int) bool {
	if codigo == "" || cantidad <= 0 {
		return false
	}

	for i, c := range codigos {
		if c == codigo {
			cantidades[i] += cantidad
			return true
		}
	}

	codigos = append(codigos, codigo)
	cantidades = append(cantidades, cantidad)
	return true
}

// RegistrarSalida descuenta productos del inventario.
func RegistrarSalida(codigo string, cantidad int) bool {
	if cantidad <= 0 {
		return false
	}

	for i, c := range codigos {
		if c == codigo && cantidades[i] >= cantidad {
			cantidades[i] -= cantidad
			return true
		}
	}

	return false
}

// ConsultarStock muestra las existencias disponibles.
func ConsultarStock(codigo string) int {
	for i, c := range codigos {
		if c == codigo {
			return cantidades[i]
		}
	}

	return 0
}

// VerificarStock comprueba la disponibilidad.
func VerificarStock(codigo string, cantidad int) bool {
	return cantidad > 0 && ConsultarStock(codigo) >= cantidad
}
