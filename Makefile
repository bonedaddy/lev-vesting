.PHONY: compile
compile:
	./scripts/solc_compile.sh \
		contracts/LevToken.sol \
		bin/levtoken \
		compilers/solc-v0.6.2
	./scripts/solc_compile.sh \
		contracts/utils/time/BokkyPooBahsDateTimeContract.sol \
		bin/datetime \
		compilers/solc-v0.6.0

.PHONY: abigen
abigen:
	./scripts/abigen.sh \
		bin/levtoken/LeverageToken.bin \
		bin/levtoken/LeverageToken.abi \
		levtoken \
		bindings/levtoken/bindings.go

	./scripts/abigen.sh \
		bin/datetime/BokkyPooBahsDateTimeContract.bin \
		bin/datetime/BokkyPooBahsDateTimeContract.abi \
		datetime \
		bindings/datetime/bindings.go \
		--alias="getDaysInMonth=getDaysInMonth2,_isLeapYear=IsLeapYear2"