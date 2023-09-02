package utils

func IsAllowedAddress(ip string, allowedIPs []string) bool {
	for _, allowedIP := range allowedIPs {
		if ip == allowedIP {
			return true
		}
	}
	return false
}
