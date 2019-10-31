from enum import Enum



class Drops(Enum):
    Pling = 3
    Plang = 5
    Plong = 7

def convert(number: int) -> str:
    sounds = ''
    for drop in Drops:
        if not number % drop.value:
            sounds += drop.name
    return sounds or str(number)




#def convert(number):
#    sounds = ''
#   if number % 3 == 0:
#      sounds = sounds + "Pling"
# if number % 5 == 0:
#    sounds = sounds + "Plang"
#    if number % 7 == 0:
#        sounds = sounds + "Plong"
#    if number % 7 != 0 and number % 5 != 0 and number % 3 != 0:
#        sounds = str(number)
#    return sounds
