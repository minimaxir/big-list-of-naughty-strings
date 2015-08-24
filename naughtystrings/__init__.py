import os


FILEPATH = os.path.join(
    os.path.abspath(os.path.dirname(os.path.dirname(__file__))), 'blns.txt')
"""Path to the file"""


def naughty_strings(filepath=FILEPATH):
    """Get the list of naughty_strings.

    By default this will get the strings from the blns.txt file

    Code is a simple port of what is already in the /scripts directory

    :param filepath: Optional filepath the the blns.txt file
    :returns: The list of naughty strings
    """

    strings = []
    with open(filepath, 'r') as f:

        # put all lines in the file into a Python list
        strings = f.readlines()
        
        # above line leaves trailing newline characters; strip them out
        strings = [x.strip(u'\n') for x in strings]
        
        # remove empty-lines and comments
        strings = [x for x in strings if x and not x.startswith(u'#')]
        
        # insert empty string since all are being removed
        strings.insert(0, u"")

    return strings
