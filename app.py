'''
    Flasky application

    Useful for testing Azure DevOps Pipeline with Anchore
'''
from flask import Flask
APP = Flask(__name__)


@APP.route('/')
def homepage():
    '''
        Root route
    '''
    return 'Hello!'


@APP.route('/blue')
def blue():
    '''
        Blue route
    '''
    return 'Blue'


@APP.route('/green')
def green():
    '''
        Green route
    '''
    return 'Green for dayz'


if __name__ == '__main__':
    APP.run(debug=True, host='0.0.0.0')
