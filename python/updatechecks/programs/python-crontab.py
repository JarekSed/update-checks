""" This module gets the latest version of python-crontab from PyPi"""
import urllib2
import re
from BeautifulSoup import BeautifulSoup
def get_version():
    """ Gets the latest version of python-crontab from PyPi. Returns the version as
    a string, or False if it can't be found"""
    url = "https://pypi.python.org/pypi/python-crontab"
    page = urllib2.urlopen(url)

    soup = BeautifulSoup(page)

    full_link = [rel['href'] for rel in soup.findAll('link') if 'version' in rel['href'] ][0]

    # version is the last part of the url, so everything after the last =
    release = full_link[ full_link.rfind('=') + 1 : ]
    return release

