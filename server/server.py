from flask import Flask, request
import sys

app = Flask(__name__)

@app.route('/')
def home():
    return f"Hello from Server running on port {sys.argv[1]}!", 200

@app.route('/status')
def health_check():
    """Health check endpoint for load balancer"""
    return "OK", 200

if __name__ == '__main__':
    port = int(sys.argv[1])  # Read port number from command line argument
    app.run(host='127.0.0.1', port=port)
