from setuptools import setup

setup(name='updatechecks',
      version='0.3',
      description='Checks web pages to determine if programs have updates',
      url='http://github.com/JarekSed/update-check',
      author='Jarek Sedlacek',
      author_email='jareksedlacek@gmail.com',
      license='MIT',
      packages=['updatechecks', 'updatechecks.programs'],
      zip_safe=False)
