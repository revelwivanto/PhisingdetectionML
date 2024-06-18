import machine_learning as ml
import requests as re
from requests.exceptions import RequestException
import json

# Define your models in a list
models = [
    ml.nb_model,
    ml.svm_model,
    ml.dt_model,
    ml.rf_model,
    ml.ab_model,
    ml.nn_model,
    ml.kn_model
]

def get_url_content(url):
    try:
        # Fetch the URL content with a timeout and without SSL verification
        response = re.get(url, verify=False, timeout=4)
        # Check if the response status code is not 200 (OK)
        if response.status_code != 200:
            print("HTTP connection was not successful for the URL: ", url)
            return None
        return response.content
    except RequestException as e:
        # Catch and print any request-related exceptions
        print("Error fetching URL:", e)
        return None

def extract_features(html_content):
    try:
        # Send a POST request to the Go feature extraction service
        response = re.post("http://localhost:8080/extract_features", data={"html": html_content})
        if response.status_code != 200:
            print("Feature extraction service responded with status code: ", response.status_code)
            return None
        
        # Parse the JSON response
        feature_response = response.json()
        return feature_response["features"]
    except RequestException as e:
        # Catch and print any request-related exceptions
        print("Error extracting features:", e)
        return None
    except json.JSONDecodeError as e:
        print("Error parsing JSON response:", e)
        return None

def classify_url(models, vector):
    for mod in models:
        result = mod.predict([vector])  # Wrap vector in a list to match expected input shape
        # Check the prediction result
        if result[0] == 0:
            print("This web page seems legitimate!")
        else:
            print("Attention! This web page is a potential PHISHING!")

def main():
    while True:
        # Prompt the user to enter a URL or 'end' to exit
        url = input("Enter the URL (enter 'end' to exit): ")
        if url.lower() == "end":
            break
        
        # Get the URL content
        content = get_url_content(url)
        if content is None:
            continue
        
        # Extract features using the Go service
        features = extract_features(content)
        if features is None:
            continue
        
        # Classify the URL using the models
        classify_url(models, features)

if __name__ == "__main__":
    main()
