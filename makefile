all:
	forge test --force

test-all:
	forge test --force -vvvv

deploy:
	./scripts/deploy
