#The instructions for this one are sort of strange. It will be easier to understand if you look at robot_name_test.py
from random import choices
from string import ascii_uppercase, digits

# standard class setting up the attributes of any robod object.
class Robot:
    def __init__(self):
        # When robots come off the factory floor, they have no name.
        self._name = None
        #stores all past names as to not repeat if reset
        self._past_names = set()
    
    @property
    def name(self):
        #If the robot's name is ever None...
        if self._name is None:
            while True:
                #...Access the random_name method...
                self._name = self.random_name()
                #If the given name is not in past names, add it and exit loop
                if self._name not in self._past_names:
                    self._past_names.add(self._name)
                    break
        return self._name
    
    def reset(self):
        """Reset robot to its factory settings."""
        # Returns the name to None
        self._name = None
    
    @staticmethod
    def random_name():
        # Two random uppercase letters followed by three random digits
        return ''.join(choices(ascii_uppercase, k=2) + choices(digits, k=3))
