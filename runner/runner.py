import html
import json, sys, requests
import urllib

api_endpoint = 'http://127.0.0.1:3000/search-by-title'
search_terms = urllib.parse.quote(str(sys.argv[1:]))
search_params = {'title': search_terms}

"""
Function to query a search api like OpenLibrary.

Parameters (dict): The search terms.
result (dict): The search result.

"""
def search_api(search_params):
    get_response = requests.get(api_endpoint, params=search_params)
    result = html.unescape(json.loads(get_response.content))
    return result

    """
    Function to display the results of a search query.

    Parameters:
    result (dict): The search result.

    """
def show_results(result):
    num_results = result['numFound']
    for i in (range(0, num_results)):
        print(f"{[i]}.\t{result['docs'][i]['title']}\n")
    print(f"Encontrado(s) {num_results} resultado(s) para sua busca")

def main():
    try:
        api_result = search_api(search_params)
        show_results(api_result)
    except Exception as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    main()