""" This module gets the latest version of neocomplecache from from vimscripts"""
import urllib2
import re
from BeautifulSoup import BeautifulSoup
def get_version():
    """ Gets the latest version of neocomplecache from vimscripts . Returns the version as
    a string, or False if it can't be found"""
    url = "http://www.vim.org/scripts/script.php?script_id=2620"
    page = urllib2.urlopen(url)

    soup = BeautifulSoup(page)

    vers = [str(link.contents[0]) for link in soup.findAll('a')]

    vers = [ver for ver in vers if re.search('(neocomplcache-\d.*(zip|tar\.gz))',ver)]

    vers.sort(reverse=True)

    match =  re.search('cache-(.*?).(zip|tar)',vers[0])

    return match.group(1)

