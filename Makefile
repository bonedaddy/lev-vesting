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
	./scripts/solc_compile.sh \
		contracts/Vesting.sol \
		bin/vesting \
		compilers/solc-v0.7.3

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

	./scripts/abigen.sh \
		bin/vesting/Vesting.bin \
		bin/vesting/Vesting.abi \
		vesting \
		bindings/vesting/bindings.go

.PHONY: interfaces
interfaces:
	./scripts/abi2solidity.sh \
		bin/datetime/BokkyPooBahsDateTimeContract.abi \
		contracts/interfaces/DateTimeInterface.sol \
		DateTimeInterface \
		"pragma solidity 0.7.3;"

.PHONY: flatten
flatten:
	./scripts/solidityFlattener.pl --mainsol=Vesting.sol --outputsol=flattened/Vesting_Flattened.sol

.PHONY: all
all: compile abigen interfaces