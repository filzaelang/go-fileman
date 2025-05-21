dev-client:
	BUILD_ENV=development bun --prefix ./views run dev #(Belum diinstall)

dev-server:
	BUILD_ENV=development go run main.go