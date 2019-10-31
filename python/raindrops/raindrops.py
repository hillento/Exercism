from enum import Enum


# Using enum to make this a more dynamic solution to the problem. It will enable a for loop.
class Drops(Enum):
    Pling = 3
    Plang = 5
    Plong = 7

def convert(number: int) -> str:
    # initialize sounds as an empty string in preperation for it to be loaded in the for loop
    sounds = ''
    # This for loop will execute once for every enum in Drops. 
    for drop in Drops:
        #An interesting property of python. Because a number that's truly divisible by a number returns 0 it is the same as false. So if the number is divisible by the enum add the enum name to the string.
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
