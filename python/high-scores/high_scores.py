def latest(scores):
    index = len(scores)-1
    return scores[index]


def personal_best(scores):
    return max(scores)


def personal_top_three(scores):
    return sorted(scores, reverse=True)[0:3]
    
