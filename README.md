## Shopping cart service 🛒
https://potent-afternoon-af0.notion.site/Shopping-cart-ac0c7820e1c34032a46eb98d8105db00

### Libs

- Decimal - https://pkg.go.dev/github.com/shopspring/decimal#section-readme

### Limitation

- Context and transaction

### Cmds

test 
```
docker build --target test -t shopping-cart . &&
docker run -t -i --rm \
	-v .:/usr/app:delegated \
		--name shopping-cart-test \
		shopping-cart \
		go clean --testcache && \
		go test -cover ./...
```

run
```
docker-compose up
```