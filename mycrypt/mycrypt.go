package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ")

func Krypter(melding []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], ALF_SEM03)
		if indeks+chiffer >= len(ALF_SEM03) {
			kryptertMelding[i] = ALF_SEM03[indeks+chiffer-len(ALF_SEM03)]
		} else {
			kryptertMelding[i] = ALF_SEM03[indeks+chiffer]
		}
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
		}
	}
	return -1
}
