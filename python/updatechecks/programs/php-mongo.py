""" This module gets the latest version of php-mongo from pecl"""
import urllib2
import re
from BeautifulSoup import BeautifulSoup
def get_version():
    """ Gets the latest version of php-mongo from pecl. Returns the version as
    a string, or False if it can't be found"""
    url = "http://pecl.php.net/package/mongo"
    page = urllib2.urlopen(url)

    soup = BeautifulSoup(page)

    releases = [rel.text for rel in soup.findAll('a',{'href':re.compile('/package/mongo/*')})]
    releases.sort(reverse=True)
    return releases[0] if len(releases) else False

