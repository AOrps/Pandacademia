from flask import Flask, render_template, request, jsonify, url_for, redirect  #pip install flask
from back_end import *
app = Flask(__name__)


# @app.route("/")
# def index():
#     if request.method == 'POST':
#         return redirect(url_for("results"))
#     return render_template('index.html', title="🩺 Daily Screening")

# @app.route("/res", methods = ['GET', 'POST'])
# def res():
#     return render_template("results.html", code=1, name="Hi")

@app.route("/", methods = ['GET', 'POST'])
def index():
    if request.method == 'POST':
        nm = request.form['nm']
        feeling = request.form['feeling']
        temp = request.form['temp']
        sick = request.form['sick']
        sore = request.form['sore']
        contact = request.form['contact']
        location = request.form['location']
        return redirect(url_for('results', nm=nm, feel=feeling, 
        temp=temp, sick=sick, sore=sore, contact=contact, 
        location=location))
    return render_template('index.html', title="🩺 Daily Screening")

@app.route("/info")
def info():
    return render_template("info.html", title="📚 CDC Info")

@app.route("/results", methods=['POST','GET'])
def results(nm, feel, temp, sick, sore, contact, location):
    ANSWER_SET = set()
    ANSWER_SET.add(feeling(feel))
    ANSWER_SET.add(temperature(temp))
    ANSWER_SET.add(yn(sick))
    ANSWER_SET.add(yn(sore))
    ANSWER_SET.add(yn(contact))
    ANSWER_SET.add(yn(location))
    
    if( 1 in ANSWER_SET):
        return render_template("results.html", code=1, name=nm)
    elif( 2 in ANSWER_SET and 1 not in ANSWER_SET):
        return render_template("results.html", code=2, name=nm)
    else:
        return render_template("results.html", code=3, name=nm)

if __name__ == "__main__":
    app.run(port=8080, debug=True)