from collections import Counter


# Score categories.
# Change the values as you see fit.
YACHT = 0
ONES = 1
TWOS = 2
THREES = 3
FOURS = 4
FIVES = 5
SIXES = 6
FULL_HOUSE = 7
FOUR_OF_A_KIND = 8
LITTLE_STRAIGHT = 9
BIG_STRAIGHT = 10
CHOICE = 11


def score(dice, category):
    # sorts the list to mkae # of a kind easier to calculate
    dice.sort()
    if category in (ONES, TWOS, THREES, FOURS, FIVES, SIXES):
        return category * dice.count(category)
    elif category == CHOICE:
        return sum(dice)
    elif category == FULL_HOUSE:
        if (sorted(Counter(dice).values()) == [2, 3]):
            return sum(dice)
        else:
            return 0
    elif category == FOUR_OF_A_KIND:
        if (max(Counter(dice).values()) >= 4):
            return max(dice)*4
        else:
            return 0
    elif category == LITTLE_STRAIGHT:
        if(dice == [1, 2, 3, 4, 5]):
            return 30
        else:
            return 0
    elif category == BIG_STRAIGHT:
        if(dice == [2, 3, 4, 5, 6]):
            return 30
        else:
            return 0
    elif category == YACHT:
        if(max(dice) == min(dice)):
            return 50
        else:
            return 0