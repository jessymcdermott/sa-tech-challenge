import os
from flask import Flask, jsonify
import json
import random

app = Flask(__name__)

image_urls = []

with app.app_context():
    """Load images from JSON file before the first request."""
    try:
        if not os.path.exists("images.json"):
            raise FileNotFoundError("images.json can not be found")

        with open("images.json", 'r') as file:
            data = json.load(file)
            images = data.get("images", [])

        bucket_name = os.environ.get("BUCKET_NAME", "random-pictures")
        image_urls = [f"https://{bucket_name}.s3.amazonaws.com/{image}" for image in images]
        app.logger.info("Images loaded successfully.")

    except Exception as e:
        app.logger.error(f"Failed to load images: {e}")
        images = []


# Route for health check
@app.route('/health')
def health():
    return jsonify({"message": "I am here, ready to pick an image", "status_code": 0})


# Route for getting a random phrase
@app.route('/imageUrl')
def get_image_url():
    phrase = choose(image_urls)
    # You can implement tracing logic here if needed
    return jsonify({"imageUrl": phrase})


# Helper function to choose a random item from a list
def choose(array):
    return random.choice(array)


if __name__ == '__main__':
    app.run(port=10116)
