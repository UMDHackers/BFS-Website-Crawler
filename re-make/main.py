from flask import Flask, url_for, request, render_template, send_from_directory
app = Flask(__name__)


@app.route("/")
def index():
    return render_template('index.html')

@app.route("/results", methods = ['POST'])
def results():
    website = request.form['website']
    return render_template('results.html', website = website)


if __name__ == "__main__":
    app.run()
