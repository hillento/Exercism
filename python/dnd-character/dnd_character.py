from random import randint
from math import floor


class Character:

    def __init__(self):
        self.strength = self.ability()
        self.dexterity = self.ability()
        self.constitution = self.ability()
        self.intelligence = self.ability()
        self.wisdom = self.ability()
        self.charisma = self.ability()
        self.hitpoints = 10 + modifier(self.constitution)

    # return random ability attribute
    def ability(self):
        return sum(sorted([randint(1, 6) for l in range(4)])[1:4])


def modifier(ability):
    return floor((ability - 10)/2)