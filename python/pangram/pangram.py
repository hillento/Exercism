# This method accepts a string called sentence and returns a bool (T/F)
def is_pangram(sentence: str) -> bool:
    #makes the string all lowercase
    sentence = sentence.lower()
    #a string containing all letters of the alphabet
    letters = 'abcdefghijklmnopqrstuvwxyz'
    #If the variable sentence has all letter contained in string it returns true
    return set(sentence).issuperset(letters)
