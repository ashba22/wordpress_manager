<template>
    <div class="grid gap-4">
        <div class="col-span-5">
            <div class="flex items-center justify-between py-4 px-2 rounded-lg mt-3">
                <!-- WordPress Core Update Section -->
                <div v-if="website.updates && website.updates.core.length > 0" class="lg:col-span-3">
                    <div class="p-1 rounded-lg">
                        <div v-for="update in website.updates.core" :key="update.version"
                            class="flex flex-row justify-between items-center hover:bg-gray-100 transition duration-300 ease-in-out rounded-lg p-3 mb-2">
                            <div class="flex flex-col">
                                <h2 class="text-xl font-bold text-green-400">Wordpress Core Update</h2>
                                <p>Current / Latest Version: {{ update.current }} / {{ update.version }}</p>
                            </div>
                            <button
                                class="bg-green-500 hover:bg-green-400 px-3 py-2 rounded focus:outline-none transition ease-in-out ml-4">
                                Update Now
                            </button>
                        </div>
                    </div>
                </div>
                <div v-else class="lg:col-span-3 p-1 rounded-lg">
                    <h3>No WordPress Core updates available.</h3>
                </div>

                <!-- Plugins Update Section -->
                <div v-if="plugins_update_bool === true" class="lg:col-span-3">
                    <div class="p-1 rounded-lg">
                        <div class="flex items-center">
                            <i class="pi pi-folder-plus"></i>
                            <h3 class="ml-2">Plugins Updates: {{ website.updates.plugins_update_count }}</h3>
                        </div>
                    </div>
                </div>
                <div v-else class="lg:col-span-3 p-1 rounded-lg">
                    <h3>No plugin updates available.</h3>
                </div>

                <!-- Themes Update Section -->
                <div v-if="website.updates && themes_update_bool === true" class="lg:col-span-3">
                    <div class="p-1 rounded-lg">
                        <div class="flex items-center">
                            <i class="pi pi-palette"></i>
                            <h3 class="ml-2">Themes Updates: {{ website.updates.plugins_update_count }}</h3>
                        </div>
                    </div>
                </div>
                <div v-else class="lg:col-span-3 p-1 rounded-lg">
                    <h3>No theme updates available.</h3>
                </div>
            </div>
            <hr class="my-2">
        </div>
        <div class="col-span-1">
            <div class="shadow-lg rounded-lg bg-gray-800 p-4 mb-4">
                    <div class="flex items-center">
                        <i class="pi pi-globe text-2xl"></i>
                        <h1 class="text-2xl font-bold ml-2">{{ website.name }}</h1>
                        <div class="h-16 w-16 rounded-full overflow-hidden ml-4">
                            <img :src="website.url + '/' + 'favicon.ico'" alt="favicon" class="h-full w-full object-cover" />
                        </div>
                    </div>
                <div class="flex justify-between items-center bg-purple-900 text-white p-4 rounded-t-lg">
                    <div class="flex items-center">
                        <a :href="website.url" target="_blank" rel="noopener noreferrer">
                            <button class="btn-primary">Visit Website</button>
                        </a>
                        <a :href="website.url + '/wp-admin'" target="_blank" rel="noopener noreferrer">
                            <button class="btn-secondary">WP-Admin</button>
                        </a>
                        <router-link :to="`/website/${website.id}/edit`">
                            <button class="btn-info">Edit Website</button>
                        </router-link>
                    </div>
                </div>
                <div class="border-3 bg-purple-800 text-white mb-4 p-3 rounded-b-lg">
                    <div class="grid grid-cols-3 gap-4 text-center">
                        <a href="#plugins-table" class="flex flex-col items-center justify-center hover-effect">
                            <i class="pi pi-folder-plus text-3xl"></i>
                            <span class="mt-2">{{ website.plugins_count }} Plugins</span>
                        </a>
                        <a href="#themes-table" class="flex flex-col items-center justify-center hover-effect">
                            <i class="pi pi-palette text-3xl"></i>
                            <span class="mt-2">{{ website.themes_count }} Themes</span>
                        </a>
                        <a href="#posts-table" class="flex flex-col items-center justify-center hover-effect">
                            <i class="pi pi-file text-3xl"></i>
                            <span class="mt-2">{{ website.posts_count }} Posts</span>
                        </a>
                        <a href="#pages-table" class="flex flex-col items-center justify-center hover-effect">
                            <i class="pi pi-file-o text-3xl"></i>
                            <span class="mt-2">{{ website.pages_count }} Pages</span>
                        </a>
                        <a href="#users-table" class="flex flex-col items-center justify-center hover-effect">
                            <i class="pi pi-users text-3xl"></i>
                            <span class="mt-2">{{ website.users_count }} Users</span>
                        </a>
                        <a href="#media-table" class="flex flex-col items-center justify-center hover-effect">
                            <i class="pi pi-image text-3xl"></i>
                            <span class="mt-2">{{ website.media_count }} Media</span>
                        </a>
                        <a href="#comments-table" class="flex flex-col items-center justify-center hover-effect">
                            <i class="pi pi-comment text-3xl"></i>
                            <span class="mt-2">{{ website.comments_count }} Comments</span>
                        </a>
                    </div>
                    <hr class="my-4">
                    <div class="grid grid-cols-1 gap-2">
                        <div v-for="(value, key) in website.info" :key="key">
                            <span class="font-semibold">{{ key }}:</span>
                            <span>{{ value }}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        
        <div class="col-span-4 shadow-lg">
            <DataTable stripedRows ref="dt" :value="website.plugins" :paginator="true" :rows="5"
                :filters="{ global: filters.pluginsFilter }" filterMode="global" id="plugins-table">
                <h2 class="text-lg font-bold ">Plugins <i class="pi pi-folder-plus"></i></h2>
                <InputText v-model="filters.pluginsFilter.value" placeholder="Search Plugins" />

                <Column field="name" header="Name" :sortable="true" />
                <Column field="version" header="Version" :sortable="true">
                    <template #body="{ data }">
                        <span v-if="data.updates !== 'Available'">{{ data.version }}</span>
                        <span v-else>{{ data.version }} <i class="pi pi-arrow-up"></i></span>
                    </template>
                </Column>
                <Column field="author" header="Author" :sortable="true" />
                <Column field="status" header="Status" :sortable="true">
                    <template #body="{ data }">
                        <div class="flex items-center">
                            <span v-if="data.status === 'Active'" class="bg-green-500 px-2 py-1 rounded-lg"
                                @click="togglePluginByName(data.name, data.path)">Active</span>
                            <span v-else @click="togglePluginByName(data.name, data.path)"
                                class="bg-red-500 px-2 py-1 rounded-lg">
                                Inactive</span>
                        </div>
                    </template>
                </Column>
            </DataTable>

            <DataTable stripedRows ref="dt" :value="website.themes" :paginator="true" :rows="5"
                :filters="{ global: filters.themesFilter }" filterMode="global" id="themes-table">
                <h2 class="text-lg font-bold ">Themes <i class="pi pi-palette"></i></h2>
                <InputText v-model="filters.themesFilter.value" placeholder="Search Themes" />

                <Column field="name" header="Name" :sortable="true" />
                <Column field="version" header="Version" :sortable="true" />
                <Column field="author" header="Author" :sortable="true" />
            </DataTable>


            <DataTable stripedRows ref="dt" :value="website.users" :paginator="true" :rows="5"
                :filters="{ global: filters.usersFilter }" filterMode="global" id="users-table">
                <h2 class="text-lg font-bold ">Users <i class="pi pi-users"></i></h2>
                <InputText v-model="filters.usersFilter.value" placeholder="Search Users" />

                <Column field="user_login" header="Login" :sortable="true" />
                <Column field="user_email" header="Email" :sortable="true" />
                <Column field="user_registered" header="Registered" :sortable="true" />
            </DataTable>

            <DataTable stripedRows ref="dt" :value="website.media" :paginator="true" :rows="5"
                :filters="{ global: filters.mediaFilter }" id="media-table" filterMode="global">
                <h2 class="text-lg font-bold ">Media <i class="pi pi-image"></i></h2>
                <InputText v-model="filters.mediaFilter.value" placeholder="Search Media" />
                <!-- if media type image/jpg or image/png show column with image-->
                <Column field="preview" header="Preview" :sortable="true">
                    <template #body="{ data }">
                        <img v-if="data.media_type === 'image/jpg' || data.media_type === 'image/png'"
                            :src="data.media_url" alt="media" />
                        <i v-else class="pi pi-file text-2xl"></i>
                    </template>
                </Column>

                <Column field="media_title" header="Title" :sortable="true" />
                <Column field="media_date" header="Date" :sortable="true" />
                <Column field="media_type" header="Type" :sortable="true" />
            </DataTable>

            <DataTable stripedRows ref="dt" :value="website.comments" :paginator="true" :rows="5"
                :filters="{ global: filters.commentsFilter }" filterMode="global" id="comments-table">
                <h2 class="text-lg font-bold ">Comments <i class="pi pi-comment"></i></h2>
                <InputText v-model="filters.commentsFilter.value" placeholder="Search Comments" />

                <Column field="comment_author" header="Author" :sortable="true" />
                <Column field="comment_date" header="Date" :sortable="true" />
                <Column field="comment_content" header="Content" :sortable="true" :truncate="true" />
            </DataTable>

            <DataTable stripedRows ref="dt" :value="website.posts" :paginator="true" :rows="5"
                :filters="{ global: filters.postsFilter }" filterMode="global" id="posts-table">
                <h2 class="text-lg font-bold ">Posts <i class="pi pi-file"></i></h2>
                <InputText v-model="filters.postsFilter.value" placeholder="Search Posts" />
                <Column field="post_url" header="URL" :sortable="true">
                    <template #body="{ data }">
                        <a :href="data.post_url" target="_blank" rel="noopener noreferrer">
                            <Button key="view" icon="pi pi-external-link" class="p-button-info" label="" />
                        </a>
                    </template>
                </Column>
                <Column field="post_image" header="Image" :sortable="true">
                    <template #body="{ data }">
                        <img v-if="data.post_image" :src="data.post_image" alt="post" />
                        <i v-else class="pi pi-file text-2xl"></i>
                    </template>
                </Column>
                <Column field="post_title" header="Title" :sortable="true" />
                <Column field="post_date" header="Date" :sortable="true" />
                <Column field="post_status" header="Status" :sortable="true">
                    <template #body="{ data }">
                        <span
                            :class="{ 'bg-green-500': data.post_status === 'publish', 'bg-red-500': data.post_status !== 'publish' }"
                            class="px-2 py-1 rounded-lg ">
                            {{ data.post_status === 'publish' ? 'Published' : data.post_status }}
                        </span>
                    </template>
                </Column>
            </DataTable>

            <DataTable stripedRows ref="dt" :value="website.pages" :paginator="true" :rows="5"
                :filters="{ global: filters.pagesFilter }" filterMode="global" id="pages-table">
                <h2 class="text-lg font-bold ">Pages <i class="pi pi-file-o"></i></h2>
                <InputText v-model="filters.pagesFilter.value" placeholder="Search Pages" />
                <Column field="page_url" header="URL" :sortable="true">
                    <template #body="{ data }">
                        <a :href="data.page_url" target="_blank" rel="noopener noreferrer">
                            <Button key="view" icon="pi pi-external-link" class="p-button-info" />
                        </a>
                    </template>
                </Column>
                <Column field="page_image" header="Image" :sortable="true">
                    <template #body="{ data }">
                        <img v-if="data.page_image" :src="data.page_image" alt="page" />
                        <i v-else class="pi pi-file text-2xl"></i>
                    </template>
                </Column>
                <Column field="page_title" header="Title" :sortable="true" />
                <Column field="page_date" header="Date" :sortable="true" />
                <Column field="page_status" header="Status" :sortable="true">
                    <template #body="{ data }">
                        {{ data.page_status }}
                        <span
                            :class="{ 'bg-green-500': data.page_status === 'publish', 'bg-red-500': data.page_status !== 'publish' }"
                            class="px-2 py-1 rounded-lg ">
                            {{ data.page_status === 'publish' ? 'Published' : data.page_status }}
                        </span>
                    </template>
                </Column>
            </DataTable>
        </div>
    </div>
</template>




<script>
import axios from "axios";
import { onMounted } from "vue";
import { FilterMatchMode, FilterOperator } from 'primevue/api';




export default {
    name: "WebsitePage",
    props: {
        id: {
            type: String,
            required: true
        },
        progressbar: {
            type: Boolean,
            default: true
        },
        plugins_update_count: {
            type: Number,
            default: 0
        },
        themes_update_count: {
            type: Number,
            default: 0
        }
    },


    data() {
        return {
            website: {},
            filters: {
                global: { value: "", matchMode: FilterMatchMode.CONTAINS },
                pluginsFilter: { value: "", matchMode: FilterMatchMode.CONTAINS },
                themesFilter: { value: "", matchMode: FilterMatchMode.CONTAINS },
                usersFilter: { value: "", matchMode: FilterMatchMode.CONTAINS },
                mediaFilter: { value: "", matchMode: FilterMatchMode.CONTAINS },
                commentsFilter: { value: "", matchMode: FilterMatchMode.CONTAINS },
                postsFilter: { value: "", matchMode: FilterMatchMode.CONTAINS },
                pagesFilter: { value: "", matchMode: FilterMatchMode.CONTAINS }
            }


        };
    },
    methods: {

        filterData() {
            this.$refs.dt.filter(this.filters.global.value, 'name', this.filters.global.matchMode);
        },

        async getWebsiteDetail() {
            // Show the progress bar
            var progressBar = document.querySelector(".p-progressbar");
            progressBar.style.display = "block";

            let retryCount = 0;
            const maxRetries = 3;

            while (retryCount < maxRetries) {
                try {
                    const response = await axios.get(`http://localhost:5001/website/${this.id}`);
                    this.website = response.data.website;

                    // Update the counts
                    this.website.plugins_count = this.website.plugins.length;
                    this.website.themes_count = this.website.themes.length;
                    this.website.posts_count = this.website.posts.length;
                    this.website.pages_count = this.website.pages.length;
                    this.website.users_count = this.website.users.length;
                    this.website.media_count = this.website.media.length;
                    this.website.comments_count = this.website.comments.length;

                    var themes_update_count = parseInt(this.website.updates.themes_update_count);
                    var plugins_update_count = parseInt(this.website.updates.plugins_update_count);

                    var plugins_update_bool = plugins_update_count > 0 ? true : false;
                    var themes_update_bool = themes_update_count > 0 ? true : false;

                    this.plugins_update_bool = plugins_update_bool;
                    this.themes_update_bool = themes_update_bool;

                    progressBar.style.display = "none";


                    break; // Exit the loop if successful
                } catch (error) {
                    console.log(error);

                    retryCount++;
                    this.$toast.open({
                        message: "Failed to fetch website details. Retrying nr. " + retryCount + " ...", // Show the error message
                        type: "error",
                        position: "top-right",
                    });
                    // wait for 2 seconds before retrying
                    await new Promise((resolve) => setTimeout(resolve, 2000));

                }
            }

            if (retryCount === maxRetries) {
                // Hide the progress bar
                progressBar.style.display = "none";
                this.$toast.open({
                    message: "Failed to fetch website details after multiple retries.", // Show the error message
                    type: "error",
                    position: "top-right",
                });
            }
        },

        togglePluginByName(name, path) {
            const api_url = `http://localhost:5001/website/${this.id}/toggle-plugin`;
            axios.post(api_url, { name: name, path: path })
                .then((response) => {
                    console.log(response.data);
                    this.$toast.open({
                        message: response.data.data, // Show the message from the server
                        type: response.data.status === "success" ? "success" : "error", // Show success or error message
                        position: "top-right",

                    });

                    this.getWebsiteDetail();
                })
                .catch((error) => {
                    console.log(error);
                });
        }
    },

    created() {
        //this.getWebsiteDetail();
    },

    mounted() {
        this.getWebsiteDetail();

    },
};

</script>


<style scoped>
#plugins-table,
#themes-table,
#users-table,
#media-table,
#comments-table,
#posts-table,
#pages-table {
    margin-top: 1rem;
    padding: 2rem;
    border-radius: 0.5rem;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);

}

img {
    width: 100px;
    height: 100px;
    object-fit: cover;
    border-radius: 0.5rem;
}



.p-inputtext {
    margin-bottom: 1rem;
    border: none;
    padding: 0.5rem;
    border-radius: 0.25rem;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    transition: background-color 0.3s ease;
}


.p-datatable {
    margin-top: 1rem;
    padding: 2rem;
    border-radius: 0.5rem;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}
</style>