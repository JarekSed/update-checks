#!/usr/bin/python2
import urllib2
import re
from BeautifulSoup import BeautifulSoup


def get_version():
    url = "http://pecl.php.net/package/mongo"
    page = urllib2.urlopen(url)

    soup = BeautifulSoup(page)

    releases = [rel.text for rel in soup.findAll('a',{'href':re.compile('/package/mongo/*')})]
    releases.sort(reverse=True)
    return releases[0] if len(releases) else False

