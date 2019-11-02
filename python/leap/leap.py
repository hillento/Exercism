def leap_year(year):
    if(year % 4 == 0):
        if(year % 100 == 0 and year % 400 != 0 ):
            return False
        if(year % 400 == 0 or year == 200):
            return True
        return True
    return False
