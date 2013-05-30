""" The Programs submodule is used to have custom
    scripts to scrape and parse the website for each program."""

import os
import pkgutil
import requests
import json
from pkg_resources import parse_version

def get_all_modules():
    """ Lists all modules in this package. This is used so we can just drop
    another module in here for some program, and it will automatically be
    picked up."""
    pkgpath = os.path.dirname(__file__)
    programs = [name for _, name, _ in pkgutil.iter_modules([pkgpath])]
    return programs



def check_all_programs():
    """ Checks the versions of all programs we know about. This returns a list
    of the programs that are out of date, where each item is a tuple tuple of
    the form  (program_name, last_known_version, current_version)"""

    out_of_date = []
    modules = get_all_modules()
    for module_name in modules:
        module = __import__("updatechecks.programs."+module_name,globals(),locals(), ['a'], -1)
        current_version = module.get_version()
        last_known_version = get_aur_version(module_name)
        if parse_version(current_version) > parse_version(last_known_version):
            out_of_date.append( (module_name, last_known_version, current_version) )
    return out_of_date



def get_aur_version(program_name):
    """ Gets the version of the program that is currently in the AUR"""
    url = "https://aur.archlinux.org/rpc.php?type=info&arg=%(name)s" % {'name': program_name}
    r = requests.get(url)
    info = json.loads(r.text)
    # Version is actually <version>-<release>, so we ignore everything after
    # the last - to get only the version
    release_start = info['results']['Version'].rfind('-')
    return info['results']['Version'][:release_start]
