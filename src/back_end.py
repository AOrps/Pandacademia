import sys
"""
Anything from this script that returns 0, is pretty cash money.

Should it return 1 (for any function), they might have covidyboi and there would be a 
        solid recommendation is to quarantine.

Should a function return 2 then recommendations will be issued to stay healthy and maybe 
        how to useful things and they will be recommendated to stay home.

tldr:
if func:
    return 1 -> Please Quarantine or stay home.                                         *High Covid Alert
    return 2 -> Recommendation is to stay home, maybe send stuff to stay healthy        *Indeterminate
    return 0 -> Safe and good to go                                                     *A1 Unit
"""

YN_SET = {"Yes", "yes", "Y", "y"}
feeling = lambda user_input: 0 if int(user_input) >= 3 else int(user_input)
temperature = lambda user_input: 0 if 97.0 < float(user_input) < 100.4 else 1
yn = lambda user_input: 1 if user_input in YN_SET else 0


        

# References
"""
# feeling: Determines the general health of the individual
def Feeling(user_input):    # user_input: int -> Number between 1 to 5, inclusive. (1 being bad, and 5 being amazing)
    uIN = int(user_input)
    if uIN >= 3:
        return 0
    else:
        return uIN

def temp(user_input):
#temperature: Determines if it is a 'sickly' temperature
#user_input: float -> used to figure out whether the temperature is that of a sickly person
    if(97.0 < float(user_input) < 100.4):
        return 0
    else:
        return 1

def yN(user_input):
    if(user_input in YN_SET):
        return 1
    else:
        return 0
"""


if __name__ == "__main__":
    pass 
    # is this unit testing?
    ## print(temperature(sys.argv[1]))
    ## print(yn(sys.argv[1]))
    ## print(feeling(sys.argv[1]))
