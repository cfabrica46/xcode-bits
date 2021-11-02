package cmd

func GetWillBeDecoded() bool {
	return willBeDecoded
}

func GetValueToDecode() string {
	return valueToDecode
}

func GetValuesToEncode() (currentHuman, currentMale, currentLegal, currentAdmin bool, currentWeight int) {
	return isHuman, isMale, isLegal, isAdministrator, weight
}
