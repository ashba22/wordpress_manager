<template>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 px-4">
        <div>
            <h1 class="text-3xl font-bold mb-4 text-white">Website Detail</h1>

            <div class="bg-gray-900 shadow-md rounded-lg p-6">
                <div v-if="website.favicon" class="flex justify-start mb-6">
                    <img
                        :src="'data:image/x-icon;base64,' + website.favicon"
                        alt="favicon"
                        class="h-16 w-16 rounded-full"
                    />
                </div>

                <div class="mb-4">
                    <p class="text-lg font-semibold mb-2 text-white">Website ID:</p>
                    <p class="text-gray-300">{{ website.id }}</p>
                </div>
                <div class="mb-4">
                    <p class="text-lg font-semibold mb-2 text-white">Website Name:</p>
                    <p class="text-gray-300">{{ website.name }}</p>
                </div>
                <div class="mb-4">
                    <p class="text-lg font-semibold mb-2 text-white">Website URL:</p>
                    <p class="text-gray-300">{{ website.url }}</p>
                </div>
                <div class="mb-4">
                    <p class="text-lg font-semibold mb-2 text-white">Website Description:</p>
                    <p class="text-gray-300">{{ website.description }}</p>
                </div>
                <div class="mb-4">
                    <p class="text-lg font-semibold mb-2 text-white">Website Status Code:</p>
                    <p class="text-gray-300">{{ website.status_code }}</p>
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
