package fetcher

import (
	"fmt"
	"math"
)

func (f *Fetcher) GetURL() {

}

func (f *Fetcher) GetJSON() {

}

// EncryptedString :
func (f *Fetcher) EncryptedString(enc int, mo int, password string) string {
	var b uint
	var pos uint
	for _, char := range password {
		b += (uint(char) << pos)
		pos += 8
	}
	crypt := int(math.Pow(float64(b), float64(enc))) % int(mo)
	return fmt.Sprintf("%x", crypt)
}
