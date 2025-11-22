package utils

func CheckRole(allowedRoles []string, userRole string) bool{
	for _, role := range allowedRoles {
			if role == userRole {
				return  true
			}
		}
		return false
}