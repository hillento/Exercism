package accumulate

//Accumulate takes an array of strings and a function.
//It inputs the strings into the given function and returns the output
func Accumulate(collection []string, fnc func(string) string) []string {

	for i := range collection {
		collection[i] = fnc(collection[i])
	}

	return collection
}
