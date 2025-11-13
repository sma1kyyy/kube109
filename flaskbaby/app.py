from flask import Flask

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"

@app.route("/greet/<name>")
def greet_user(name):
    return f"<h1>Hello, {name}!</h1>"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=80)