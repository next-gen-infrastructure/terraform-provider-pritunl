build:
	go build -gcflags="all=-N -l" -o ~/.terraform.d/registry.terraform.io/next-gen-infrastructure/pritunl/0.5.3/darwin_arm64/terraform-provider-pritunl_v0.5.3 main.go
	go build -gcflags="all=-N -l" -o ~/.terraform.d/registry.opentofu.org/next-gen-infrastructure/pritunl/0.5.3/darwin_arm64/terraform-provider-pritunl_v0.5.3 main.go

test:
	@docker rm tf_pritunl_acc_test -f || true
	@docker run \
		--platform linux/amd64 \
		--name tf_pritunl_acc_test \
		--hostname pritunl.local \
		--rm -d --privileged \
		-p 1194:1194/udp \
		-p 1194:1194/tcp \
		-p 80:80/tcp \
		-p 443:443/tcp \
		-p 27017:27017/tcp \
		ghcr.io/jippi/docker-pritunl:latest

	sleep 20

	@chmod +x ./tools/wait-for-it.sh
	./tools/wait-for-it.sh localhost:27017 -- echo "mongodb is up"

	# enables an api access for the pritunl user, updates an api token and secret
	@docker exec -i tf_pritunl_acc_test mongo pritunl < ./tools/mongo.js

	TF_ACC=1 \
	PRITUNL_URL="https://localhost/" \
	PRITUNL_INSECURE="true" \
	PRITUNL_TOKEN=tfacctest_token \
	PRITUNL_SECRET=tfacctest_secret \
	go test -v -cover -count 1 ./internal/provider

	@docker rm tf_pritunl_acc_test -f
