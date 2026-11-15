package proveedores

var rucs []string
var nombres []string
var telefonos []string

// RegistrarProveedor permite ingresar un nuevo proveedor.
func RegistrarProveedor(ruc string, nombre string, telefono string) bool {
	if ruc == "" || nombre == "" || telefono == "" {
		return false
	}

	for _, r := range rucs {
		if r == ruc {
			return false
		}
	}

	rucs = append(rucs, ruc)
	nombres = append(nombres, nombre)
	telefonos = append(telefonos, telefono)

	return true
}

// ConsultarProveedor busca un proveedor por su RUC.
func ConsultarProveedor(ruc string) string {
	for i, r := range rucs {
		if r == ruc {
			return nombres[i] + " - " + telefonos[i]
		}
	}

	return "Proveedor no encontrado"
}

// ActualizarProveedor modifica el telefono registrado.
func ActualizarProveedor(ruc string, telefono string) bool {
	if telefono == "" {
		return false
	}

	for i, r := range rucs {
		if r == ruc {
			telefonos[i] = telefono
			return true
		}
	}

	return false
}

// EliminarProveedor elimina un proveedor del registro.
func EliminarProveedor(ruc string) bool {
	for i, r := range rucs {
		if r == ruc {
			rucs = append(rucs[:i], rucs[i+1:]...)
			nombres = append(nombres[:i], nombres[i+1:]...)
			telefonos = append(telefonos[:i], telefonos[i+1:]...)

			return true
		}
	}

	return false
}
