PHONY: mysql-up
mysql-up:
	podman run \
		--name "mysql" \
		-dt \
		-e MYSQL_ROOT_HOST='%' \
		-e MYSQL_ROOT_PASSWORD='password' \
		-p 3306:3306/tcp \
		"docker.io/mysql/mysql-server:8.0.28"

PHONY: mysql-down
mysql-down:
	podman stop "mysql"
	podman rm "mysql"