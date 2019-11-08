import math

def score(x, y):
    hyp = pythag(x, y)
    if(hyp > 10 or (abs(x)>10 or abs(y) > 10)):
        return 0
    elif(hyp > 5 or (abs(x)>5 or abs(y) > 5)):
        return 1
    elif(hyp > 1  or (abs(x)>1 or abs(y) > 1)):
        return 5
    elif(hyp <= 1):
        return 10
def pythag(x, y):
    return math.sqrt((x*x)+(y*y))
