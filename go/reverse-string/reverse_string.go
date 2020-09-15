package reverse

//Reverse returns a string oString that is a reverse of the input string iString
func Reverse(iString string)string {

	if len(iString) == 0 {
		return ""
	}

	rSlice:= []rune(iString)
	var oString string

	for i:=len(rSlice)-1; i >=0; i-- {
		oString += string(rSlice[i])
	}
	
	return oString
	
}
