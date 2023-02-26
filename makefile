all:
	forge test --force
	cd luz ; make ; cd -

test-all:
	forge test --force -vvvv

deploy:
	./scripts/deploy
