//lets make the converter
package converter

import (
	"strconv"
	"strings"
)

//convert hex to decimal
func ConvertHex(text string) string{
	words:= strings.Fields(text)
	for i := 0; i< len(words); i++{
		if words[i]== "(hex)" && i > 0{
			//convert previous word from hex to decimal
			hexNum:= words[i+1]
			decimal,err := strconv.PerseInt(hexNum,16,64)
			if err == nil{
				words[i-1]= strconv
			}
		}
	}
}

//we will continue tomorrow