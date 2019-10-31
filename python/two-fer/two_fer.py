#Method accepts string input that is set to 'you' if no input is given
def two_fer(name: str = 'you') -> str:
    #Returns the string with the variable name inserted. This will be 'you' if no input is given.
    return f"One for {name}, one for me."