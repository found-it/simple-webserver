from flask import Flask
app = Flask(__name__)

@app.route('/')
def homepage():
    return 'Hello!'

i = 0
@app.route('/increment')
def increment():
    global i
    i += 1
    return f'The number is now {i}'

@app.route('/blue')
def blue():
    return 'Blue'

@app.route('/green')
def green():
    return 'Green'


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
