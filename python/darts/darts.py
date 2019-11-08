import math

def score(x, y):
    hyp = pythag(x, y)
    if(hyp > 10):
        return 0
    elif(hyp > 5):
        return 1
    elif(hyp > 1):
        return 5
    else:
        return 10
def pythag(x, y):
    return math.sqrt((x*x)+(y*y))
