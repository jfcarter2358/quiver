import re

class Token:
    def __init__(self):
        self.m_val = ""
        self.m_type = ""
        self.m_left = []
        self.m_right = []

    def __repr__(self):
        left_string = "null"
        right_string = "null"
        if self.m_right > 0:
            left_string = str(self.m_left[0])
        if self.m_right:
            right_string = str(self.m_right[0])
        cleaned_val = self.m_val.replace("\"", "\\\"")
        return f'{{"val": "{cleaned_val}", "type": "{self.m_type}", "left": "{left_string}", "right": "{right_string}"}}'

    def __str__(self):
        return self.__repr__()
