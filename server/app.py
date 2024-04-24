import uuid
from flask import Flask, jsonify, request
from flask_cors import CORS
import requests
import base64
WEBSITES = [
    {
        'id': '4', # uuid.uuid4().hex, # '4' for testing purpose
        'name': 'WP1',
        'url': 'https://wp1.goxcms.xyz/',
        'active': True
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
            ### ping the website to get the status of the website ###
            try:
                response = requests.get(website['url'])
                website['status_code'] = response.status_code
                favicon_url = website['url'] + '/favicon.ico'
                response = requests.get(favicon_url)
                favicon_base64 = base64.b64encode(response.content).decode('utf-8')
                website['favicon'] = favicon_base64
            except:
                website['status_code'] = 404
    
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
