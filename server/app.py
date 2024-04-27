import uuid
from flask import Flask, jsonify, request
from flask_cors import CORS
import requests
import base64
import threading
import concurrent.futures
WEBSITES = [
    {
        'id': '4', # uuid.uuid4().hex, # '4' for testing purpose
        'name': 'WP1',
        'url': 'https://wp1.goxcms.xyz/',
        'api_key': 'ewlg6v944t4qNrVb7R5iVijOgvETLEdqhIbheZwkJYs7RpeCyT7753cxStQLOPYp',
        'active': True,
        'description': 'This is a WordPress website'
    },
    {
        'id': '5', # uuid.uuid4().hex, # '5' for testing purpose
        'name': 'WP2',
        'url': 'https://wp2.goxcms.xyz/',
        'api_key': 'LaBslUiBe5M61DDgChntNaV2qSza4vUfEaoAAR6ASizrhRGRAf6a5p4RV6F3uTMc',
        'active': True,
        'description': 'This is a WordPress website'
    }
]

app = Flask(__name__)
app.config.from_object(__name__)


CORS(app, resources={r'/*': {'origins': '*'}})

def remove_website(website_id):
    for website in WEBSITES:
        if website['id'] == website_id:
            WEBSITES.remove(website)
            return True
    return False

@app.route('/ping', methods=['GET'])
def ping_pong():
    return jsonify('pong!')

@app.route('/websites', methods=['GET', 'POST'])
def all_websites():
    response_object = {'status': 'success'}
    if request.method == 'POST':
        post_data = request.get_json()
        WEBSITES.append({
            'id': uuid.uuid4().hex,
            'name': post_data.get('name'),
            'url': post_data.get('url'),
            'active': post_data.get('active')
        })
        response_object['message'] = 'Website added!'
    else:
        # Sorting logic
        sort_by = request.args.get('sort_by', 'name')
        reverse = request.args.get('reverse', 'false').lower() == 'true'
        sorted_websites = sorted(WEBSITES, key=lambda x: x[sort_by], reverse=reverse)
        response_object['websites'] = sorted_websites
    return jsonify(response_object)

@app.route('/website/<website_id>', methods=['PUT', 'DELETE', 'GET'])
def single_website(website_id):
    response_object = {'status': 'success'}
    ##### get website by id #####
    if request.method == 'GET':
        website = next((site for site in WEBSITES if site['id'] == website_id), None)
        if website:
            response_object['website'] = website
            try:
                response = requests.get(website['url'])
                website['status_code'] = response.status_code
                favicon_url = f"{website['url']}/favicon.ico"
                response = requests.get(favicon_url)
                favicon_base64 = base64.b64encode(response.content).decode('utf-8')
                website['favicon'] = favicon_base64
                plugins_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/installed-plugins?api_key={website['api_key']}"
                themes_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/installed-themes?api_key={website['api_key']}"
                website_info_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/website-info?api_key={website['api_key']}"
                posts_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/posts?api_key={website['api_key']}"
                pages_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/pages?api_key={website['api_key']}"
                comments_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/comments?api_key={website['api_key']}"
                users_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/users?api_key={website['api_key']}"
                media_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/media?api_key={website['api_key']}"
                updates_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/updates?api_key={website['api_key']}"
                def fetch_data(url, result):
                    try:
                        response = requests.get(url)
                        response.raise_for_status()
                        result.append(response.json())
                    except requests.exceptions.RequestException as e:
                        result.append(None)

                plugins_result = []
                themes_result = []
                website_result = []
                posts_result = []
                pages_result = []
                comments_result = []
                users_result = []
                media_result = []
                updates_result = []

                plugins_thread = threading.Thread(target=fetch_data, args=(plugins_endpoint, plugins_result))
                themes_thread = threading.Thread(target=fetch_data, args=(themes_endpoint, themes_result))
                website_thread = threading.Thread(target=fetch_data, args=(website_info_endpoint, website_result))
                posts_thread = threading.Thread(target=fetch_data, args=(posts_endpoint, posts_result))
                pages_thread = threading.Thread(target=fetch_data, args=(pages_endpoint, pages_result))
                comments_thread = threading.Thread(target=fetch_data, args=(comments_endpoint, comments_result))
                users_thread = threading.Thread(target=fetch_data, args=(users_endpoint, users_result))
                media_thread = threading.Thread(target=fetch_data, args=(media_endpoint, media_result))
                updates_thread = threading.Thread(target=fetch_data, args=(updates_endpoint, updates_result))
                website_thread.start()
                plugins_thread.start()
                themes_thread.start()
                posts_thread.start()
                pages_thread.start()
                comments_thread.start()
                users_thread.start()
                media_thread.start()
                updates_thread.start()

                website_thread.join()
                plugins_thread.join()
                themes_thread.join()
                posts_thread.join()
                pages_thread.join()
                comments_thread.join()
                users_thread.join()
                media_thread.join()
                updates_thread.join()

                website['plugins'] = plugins_result[0]
                website['themes'] = themes_result[0]
                website['info'] = website_result[0]
                website['posts'] = posts_result[0]
                website['pages'] = pages_result[0]
                website['comments'] = comments_result[0]
                website['users'] = users_result[0]
                website['media'] = media_result[0]
                website['updates'] = updates_result[0]
                    
            except requests.exceptions.RequestException as e:
                website['status_code'] = 404
                response_object['status'] = 'error'
                response_object['message'] = 'Error occurred while fetching website data'
    
    if request.method == 'PUT':
        post_data = request.get_json()
        remove_website(website_id)
        WEBSITES.append({
            'id': uuid.uuid4().hex,
            'name': post_data.get('name'),
            'url': post_data.get('url'),
            'active': post_data.get('active')
        })
        response_object['message'] = 'Website updated!'
    
    if request.method == 'DELETE':
        if remove_website(website_id):
            response_object['message'] = 'Website removed!'
        else:
            response_object['status'] = 'error'
            response_object['message'] = 'Website not found!'
    
    return jsonify(response_object)

if __name__ == '__main__':
    app.run()
