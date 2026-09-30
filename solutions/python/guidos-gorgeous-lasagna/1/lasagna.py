"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language:
https://en.wikipedia.org/wiki/Guido_van_Rossum

This is a module docstring, used to describe the functionality
of a module and its functions and/or classes.
"""


#TODO (student): define your EXPECTED_BAKE_TIME (required) and PREPARATION_TIME (optional) constants below.
EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2

#TODO (student): Remove 'pass' and complete the 'bake_time_remaining()' function below.
def bake_time_remaining(elapsed_bake_time: int) -> int:
    """Calculate the bake time remaining.

    Parameters:
        elapsed_bake_time (int): The baking time already elapsed.

    Returns:
        int: The remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """

    return EXPECTED_BAKE_TIME - elapsed_bake_time


#TODO (student): Define the 'preparation_time_in_minutes()' function below.
def preparation_time_in_minutes(number_of_layers: int) -> int:
    """
    Calculate the preparation time.
    
    Parameters:
        number_of_layers (int): The number of layers in the lasagna
    Returns:
        int: Time taken to prepare the lasgana given the number of layers

    Function that takes the number of layers as input and returns the time in minutes
    needed to prepare the lasagna. Assuming that time taken to prepare each layer is 2
    minutes, that comes out to be `number_of_layers` * 2. 
    """
    return number_of_layers * PREPARATION_TIME

# To avoid the use of magic numbers (see: https://en.wikipedia.org/wiki/Magic_number_(programming)), you should define a PREPARATION_TIME constant.
# You can do that on the line below the 'EXPECTED_BAKE_TIME' constant.
# This will make it easier to do calculations, and make changes to your code.
    

#TODO (student): define the 'elapsed_time_in_minutes()' function below.

def elapsed_time_in_minutes(number_of_layers: int, elapsed_bake_time: int) -> int:
    """Calculate the elapsed cooking time.
    
    Parameters:
        number_of_layers (int): The number of layers in the lasagna.
        elapsed_bake_time (int): Time the lasagna has been baking in the oven.
    
    Returns:
        int: The total time elapsed (in minutes) preparing and baking.

    This function takes two integers representing the number of lasagna 
    layers and the time already spent baking. It calculates 
    the total elapsed minutes spent cooking (preparing + baking)
    """
    prep_time = preparation_time_in_minutes(number_of_layers)
    return prep_time + elapsed_bake_time


# TODO (student): Remember to go back and add docstrings to all your functions
#  (you can copy and then alter the one from bake_time_remaining.)
