def value(colors):
    colorVals = [
        'black', 
        'brown', 
        'red', 
        'orange', 
        'yellow',
        'green',
        'blue',
        'violet',
        'grey',
        'white'
        ]

    return int(str(colorVals.index(colors[0])) + str(colorVals.index(colors[1])))
