## Shopping cart service 🛒
https://potent-afternoon-af0.notion.site/Shopping-cart-ac0c7820e1c34032a46eb98d8105db00

## Data seed

| Item_id | Name    | Price (USD) |
|---------|---------|-------------|
| 10      | T-shirt | 12.99       |
| 20      | Jeans   | 25.00       |
| 30      | Dress   | 20.65       |

User ID: bba82f7a-caa1-4587-819b-6db46e14fc60

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