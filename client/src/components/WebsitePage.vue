<template>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 px-4">
        <div>
            <h1 class="text-3xl font-bold mb-4 text-dark">Website Detail</h1>

            <div class="bg-gray-900 shadow-md rounded-lg p-6">
                <div class="flex justify-start mb-6">
                    <img v-if="website.favicon" :src="'data:image/x-icon;base64,' + website.favicon" alt="favicon" class="h-16 w-16 rounded-full" />
                    <div v-else class="h-16 w-16 rounded-full bg-gray-800"></div> <!-- Placeholder for favicon if not available -->
                </div>

                <div class="grid grid-cols-2 gap-4">
                    <div>
                        <p class="text-lg font-semibold mb-2 text-white">Website ID:</p>
                        <p class="text-gray-300">{{ website.id }}</p>
                    </div>
                    <div>
                        <p class="text-lg font-semibold mb-2 text-white">Website Name:</p>
                        <p class="text-gray-300">{{ website.name }}</p>
                    </div>
                    <div>
                        <p class="text-lg font-semibold mb-2 text-white">Website URL:</p>
                        <p class="text-gray-300">{{ website.url }}</p>
                    </div>
                    <div>
                        <p class="text-lg font-semibold mb-2 text-white">Website Description:</p>
                        <p class="text-gray-300">{{ website.description }}</p>
                    </div>
                    <div class="col-span-2">
                        <p class="text-lg font-semibold mb-2 text-white">Website Status:</p>
                        <p :class="{
                            'text-green-500': website.status_code === 200,
                            'text-yellow-500': website.status_code >= 400 && website.status_code < 500,
                            'text-red-500': website.status_code >= 500 && website.status_code < 600,
                            'text-gray-500': ![200, 400, 500].includes(website.status_code)
                        }">
                            <span class="inline-block mr-1 animate-pulse">●</span>
                            {{ website.status_code === 200 ? 'Website is up and running' : 
                               website.status_code >= 400 && website.status_code < 500 ? 'Client error: Website may have an issue on the client side' : 
                               website.status_code >= 500 && website.status_code < 600 ? 'Server error: Website may have an issue on the server side' : 
                               'Unknown status: Unable to determine website status' }}
                        </p>
                    </div>
                </div>
            </div>
        </div>
        <div>
            <h1 class="text-3xl font-bold mb-4 text-dark">Website Plugins</h1>
            <div class="bg-gray-900 shadow-md rounded-lg p-6">
                <div class="overflow-x-auto">
                    <table class="min-w-full divide-y divide-gray-700">
                        <thead>
                            <tr>
                                <th class="px-4 py-2 text-left text-gray-300">Name</th>
                                <th class="px-4 py-2 text-left text-gray-300">Version</th>
                                <th class="px-4 py-2 text-left text-gray-300">Author</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="plugin in website.plugins" :key="plugin.name">
                                <td class="px-4 py-2 text-gray-300">{{ plugin.name.substring(0, 20) }}</td>
                                <td class="px-4 py-2 text-gray-300">{{ plugin.version }}</td>
                                <td class="px-4 py-2 text-gray-300">{{ plugin.author }}</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
        
    </div>
 
</template>





<script>
import axios from "axios";


export default {
    props: {
        id: {
            type: String,
            required: true
        }
    },
    data() {
        return {
            website: {},
        };
    },
    methods: {
        getWebsiteDetail() {
            axios.get(`http://localhost:5001/website/${this.id}`)
                .then((response) => {
                    this.website = response.data.website;
                   
                    console.log(this.website.name);
                })
                .catch((error) => {
                    console.log(error);
                });
        
        }
     
    },
    mounted() {
       this.getWebsiteDetail();
    },
};

</script>


<style scoped>
/* Your component styles here */
</style>
