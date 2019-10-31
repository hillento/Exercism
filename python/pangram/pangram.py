def is_pangram(sentence: str) -> bool:
    sentence = sentence.lower()
    letters = 'abcdefghijklmnopqrstuvwxyz'
    return set(sentence).issuperset(letters)
