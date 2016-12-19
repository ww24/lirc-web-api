package config

const (
	dev  = "development"
	prod = "production"
)

// IsProd check production mode
func IsProd() bool {
	return Mode == prod
}

// IsDev check development mode
func IsDev() bool {
	return Mode == dev
}
