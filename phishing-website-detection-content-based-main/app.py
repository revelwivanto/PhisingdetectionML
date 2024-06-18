import machine_learning as ml
import feature_extraction_backup as fe
from bs4 import BeautifulSoup
import requests as re
from requests.exceptions import RequestException

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

def classify_url(models, vector):
    for mod in models:
        result = mod.predict(vector)
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
        
        # Parse the content with BeautifulSoup
        soup = BeautifulSoup(content, "html.parser")
        # Assuming fe.create_vector is defined elsewhere
        vector = [fe.create_vector(soup)]
        
        # Classify the URL using the models
        classify_url(models, vector)

if __name__ == "__main__":
    main()
