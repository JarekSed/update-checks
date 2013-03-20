""" The Programs submodule is used to have custom
    scripts to scrape and parse the website for each program."""

import os
import pkgutil
from pkg_resources import parse_version

def get_all_modules():
    """ Lists all modules in this package. This is used so we can just drop
    another module in here for some program, and it will automatically be
    picked up."""
    pkgpath = os.path.dirname(__file__)
    programs = [name for _, name, _ in pkgutil.iter_modules([pkgpath])]
    return programs



def check_all_programs():
    modules = get_all_modules()
    for module_name in modules:
        module = __import__("updatechecks.programs."+module_name,globals(),locals(), ['a'], -1)
        current_version = module.get_version()
        if parse_version(current_version) > parse_version(module.get_last_known_version()):
            print module_name, "is out of date, newest version is", current_version
        else:
            print module_name, "looks good, current version is", current_version
