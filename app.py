
from flask import Flask, render_template, request, make_response
app = Flask(__name__)

@app.route('/')
def index():
    return render_template("index.html")

@app.route('/dev')
def index_dev():
    return render_template("index_dev.html")

@app.route('/get_shortest_route', methods=['POST'])
def get_shortest_route():
    
    res = make_response("sdf")
    res.headers['Access-Control-Allow-Origin'] = "*"

    return res
