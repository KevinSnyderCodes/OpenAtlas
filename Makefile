.PHONY: test
test:
	go test -v ./...

.PHONY: certificate
certificate:
	cd ./cert && mkcert -key-file key.pem -cert-file cert.pem localhost