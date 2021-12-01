package decrypt

type Decode struct {
	Str   string
	Score int
}

func MakeDecryption(str []byte) []Decode {

	decode := make([]Decode, 256)

	for i := 0; i < 256; i++ {

		var xor byte
		for j := 0; j < len(str); j++ {
			xor = str[j] ^ byte(i)
			decode[i].Str += string(xor)
		}

		decode[i].Score += getCharWeight(xor)

	}

	return decode
}
func getCharWeight(char byte) int {
	wm := map[byte]int{
		byte('U'): 2,
		byte('u'): 2,
		byte('L'): 3,
		byte('l'): 3,
		byte('D'): 4,
		byte('d'): 4,
		byte('R'): 5,
		byte('r'): 5,
		byte('H'): 6,
		byte('h'): 6,
		byte('S'): 7,
		byte('s'): 7,
		byte(' '): 8,
		byte('N'): 9,
		byte('n'): 9,
		byte('I'): 10,
		byte('i'): 10,
		byte('O'): 11,
		byte('o'): 11,
		byte('A'): 12,
		byte('a'): 12,
		byte('T'): 13,
		byte('t'): 13,
		byte('E'): 14,
		byte('e'): 14,
	}
	return wm[char]
}
