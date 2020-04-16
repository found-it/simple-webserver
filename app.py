from flask import Flask
flsky = Flask(__name__)

@flsky.route('/')
def homepage():
    return 'Hello!'

i = 0
@flsky.route('/increment')
def increment():
    global i
    i += 1
    return f'The number is now {i}'

@flsky.route('/hi')
def hi():
    return 'Hi there'


if __name__ == '__main__':
    flsky.run(debug=True, host='0.0.0.0')
