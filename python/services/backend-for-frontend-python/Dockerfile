# Use the official Python image as a base
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /app

# Copy the dependencies file to the working directory
COPY requirements.txt .

# Install Flask and other dependencies
RUN pip install --no-cache-dir -r requirements.txt

#open tel
RUN opentelemetry-bootstrap --action=install

# Copy the rest of the application code to the working directory
COPY . .

# Expose port to the outside world
EXPOSE 10115

ENV FLASK_APP=server.py

#run pls dear god work
CMD ["opentelemetry-instrument", "python", "./run"]

# Command to run the Flask application
CMD ["flask", "run", "-p 10115", "--host=0.0.0.0"]
