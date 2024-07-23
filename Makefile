PHONY: create
create:
	podman play kube mysql.yaml

PHONY: up
up:
	podman pod start mysql_pod

PHONY: down
down:
	podman pod stop mysql_pod

PHONY: rm
rm:
	podman pod rm mysql_pod
