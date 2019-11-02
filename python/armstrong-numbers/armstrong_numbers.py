import math

def is_armstrong_number(number):
    digits = len(str(number))
    lst = [int(x) for x in str(number)]
    armstrong = [x**digits for x in lst]
    if(sum(armstrong) == number):
        return True
    return False
