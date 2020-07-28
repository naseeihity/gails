package helpers

func isActive(action string, param string) string {
	var cls string
	if action == param {
		cls = "active"
	}
	return cls
}
