package usuarios

var nombres []string
var roles []string

// RegistrarUsuario permite registrar un usuario con su rol.
func RegistrarUsuario(nombre string, rol string) bool {
	if nombre == "" || rol == "" {
		return false
	}

	for _, n := range nombres {
		if n == nombre {
			return false
		}
	}

	nombres = append(nombres, nombre)
	roles = append(roles, rol)

	return true
}

// ConsultarUsuario verifica si un usuario esta registrado.
func ConsultarUsuario(nombre string) string {
	for i, n := range nombres {
		if n == nombre {
			return "Usuario: " + nombres[i] +
				" - Rol: " + roles[i]
		}
	}

	return "Usuario no encontrado"
}

// ConsultarRol devuelve el rol asignado al usuario.
func ConsultarRol(nombre string) string {
	for i, n := range nombres {
		if n == nombre {
			return roles[i]
		}
	}

	return "Rol no encontrado"
}

// ActualizarRol modifica el rol de un usuario.
func ActualizarRol(nombre string, rol string) bool {
	if rol == "" {
		return false
	}

	for i, n := range nombres {
		if n == nombre {
			roles[i] = rol
			return true
		}
	}

	return false
}
