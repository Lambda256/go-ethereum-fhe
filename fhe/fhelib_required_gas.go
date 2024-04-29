package fhevm

func fheAddSubRequiredGas(input []byte) uint64 {
	return 0
}

func fheMulRequiredGas(input []byte) uint64 {
	return 0
}

func fheLeRequiredGas(input []byte) uint64 {
	return 0
}

func fheLtRequiredGas(input []byte) uint64 {
	// Implement in terms of le, because le and lt costs are currently the same.
	return fheLeRequiredGas(input)
}

func fheEqRequiredGas(input []byte) uint64 {
	// Implement in terms of le, because comparison costs are currently the same.
	return fheLeRequiredGas(input)
}

func fheGeRequiredGas(input []byte) uint64 {
	// Implement in terms of le, because comparison costs are currently the same.
	return fheLeRequiredGas(input)
}

func fheGtRequiredGas(input []byte) uint64 {
	// Implement in terms of le, because comparison costs are currently the same.
	return fheLeRequiredGas(input)
}

func fheNeRequiredGas(input []byte) uint64 {
	// Implement in terms of le, because comparison costs are currently the same.
	return fheLeRequiredGas(input)
}

func fheNegRequiredGas(input []byte) uint64 {
	input = input[:minInt(32, len(input))]

	return 0
}

func fheNotRequiredGas(input []byte) uint64 {
	input = input[:minInt(32, len(input))]

	return 0
}

func trivialEncryptRequiredGas(input []byte) uint64 {
	input = input[:minInt(33, len(input))]

	return 0
}
func registerKeyRequiredGas(input []byte) uint64 {
	input = input[:minInt(33, len(input))]

	return 0
}

func fheDivRequiredGas(input []byte) uint64 {
	input = input[:minInt(65, len(input))]

	return 0
}

func fheRemRequiredGas(input []byte) uint64 {
	input = input[:minInt(65, len(input))]

	return 0
}
