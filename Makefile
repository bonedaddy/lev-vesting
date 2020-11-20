.PHONY: compile
compile:
	./scripts/solc_compile.sh \
		contracts/LevToken.sol \
		bin/levtoken \
		compilers/solc-v0.6.2

.PHONY: abigen
abigen:
	./scripts/abigen.sh \
		bin/levtoken/LeverageToken.bin \
		bin/levtoken/LeverageToken.abi \
		levtoken \
		bindings/levtoken/bindings.go