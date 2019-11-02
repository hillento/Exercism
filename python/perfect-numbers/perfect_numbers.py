def classify(number):
    factors = []
    if(str(number).isdigit() == False or number < 1):
        raise ValueError("Meaningful message indicating the source of the error")
    for i in range(1, number):
        if number % i == 0:
            factors.append(i)
    aliquot = sum(factors)
    if(aliquot == number):
        return "perfect"
    if(aliquot > number):
        return "abundant"
    if(aliquot < number):
        return "deficient"
