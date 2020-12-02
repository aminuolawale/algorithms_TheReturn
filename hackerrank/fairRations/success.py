
# whoever submitted the challenge did not prepare it well for the go language so I had to
# just submit in python
def fairRations(B):
    count = 0
    for index, val in enumerate(B):
        if index == len(B)-1:
            if not isEven(val):
                return "NO"
        if not isEven(val):
            B[index] += 1
            B[index+1] += 1
            count += 2

    return count


def isEven(val):
    return val % 2 == 0


if __name__ == "__main__":
    fairRations([1, 2])
