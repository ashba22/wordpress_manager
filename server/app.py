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

        response_object['websites'] = WEBSITES
    return jsonify(response_object)

@app.route('/website/<website_id>', methods=['PUT', 'DELETE', 'GET'])
def single_website(website_id):
    response_object = {'status': 'success'}
    ##### get website by id #####
    if request.method == 'GET':
        website = next((site for site in WEBSITES if site['id'] == website_id), None)
        if website:
            response_object['website'] = website
           
           
       
            favicon_url = f"{website['url']}/favicon.ico"
            plugins_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/installed-plugins?api_key={website['api_key']}"
            themes_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/installed-themes?api_key={website['api_key']}"
            website_info_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/website-info?api_key={website['api_key']}"
            posts_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/posts?api_key={website['api_key']}"
            pages_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/pages?api_key={website['api_key']}"
            comments_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/comments?api_key={website['api_key']}"
            users_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/users?api_key={website['api_key']}"
            media_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/media?api_key={website['api_key']}"
            updates_endpoint = f"{website['url']}/wp-json/wp-manager-plugin/v1/updates?api_key={website['api_key']}"
            
      
            with concurrent.futures.ThreadPoolExecutor() as executor:
                # Submitting concurrent requests
                futures = [
                    executor.submit(requests.get, url) for url in [favicon_url, plugins_endpoint, themes_endpoint, website_info_endpoint, posts_endpoint, pages_endpoint, comments_endpoint, users_endpoint, media_endpoint, updates_endpoint]
                ]
                
                # Getting results from concurrent requests
                results = [future.result() for future in futures]
                
                website['status_code'] = results[0].status_code
                favicon_base64 = base64.b64encode(results[0].content).decode('utf-8')
                website['favicon'] = favicon_base64
                website['plugins'] = results[1].json()
                website['themes'] = results[2].json()
                website['info'] = results[3].json()
                website['posts'] = results[4].json()
                website['pages'] = results[5].json()
                website['comments'] = results[6].json()
                website['users'] = results[7].json()
                website['media'] = results[8].json()
                website['updates'] = results[9].json()
                website['url'] = website['url']
        else:
            response_object['status'] = 'error'
            response_object['message'] = 'Website not found!'


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
