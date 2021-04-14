request:
	curl -k -sL https://localhost/hello

gen_server_crt:
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

gen_server_key:
	openssl genrsa -out server.key 1024
