package iiko

import "strconv"

var (
	baseURL string = "https://api-ru.iiko.services"
	// iikoCloud API Timeout in seconds by default
	defaultTimeout = "15"
)

// SetDefaultTimeout sets default Timeout header for all requests. By default 15 seconds.
func SetDefaultTimeout(seconds int) {
	defaultTimeout = strconv.Itoa(seconds)
}
