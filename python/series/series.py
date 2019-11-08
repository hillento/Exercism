def slices(series, length):
    if(length <= len(series) and length > 0):
        subStrArr = []
        for i in range(0, (len(series) - length + 1)):
            sub_str = series[i:(i+length)]
            subStrArr.append(sub_str)
        return subStrArr
    else:
        raise ValueError("Cannot make substrings larger than original string")
