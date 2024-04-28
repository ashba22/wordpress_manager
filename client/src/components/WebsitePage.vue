<template>

<div class="container mx-auto p-4">
    <div class="grid gap-4">
        <div v-if="website.updates && website.updates.core.length > 0" class="lg:col-span-3">
            <div class="text-white p-1 rounded-lg shadow">
                <div>
                    <div v-for="update in website.updates.core" :key="update.version"
                        class="flex flex-row justify-between items-center bg-gray-800 hover:bg-gray-600 transition duration-300 ease-in-out shadow-md rounded-lg p-3 mb-2">
                        <div class="flex flex-col">
                            <h2 class="text-xl font-bold text-green-400">Wordpress Core Update</h2>
                            <h4 class="text-lg font-bold">WordPress Core</h4>
                            <p class="">Current Version: {{ update.current }}</p>
                            <a :href="update.download" class="text-blue-500 hover:text-blue-300">Manual Update download
                                link</a>
                        </div>
                        <button
                            class="bg-green-500 hover:bg-green-400 text-white px-3 py-2 rounded focus:outline-none transition ease-in-out">
                            Update Now
                        </button>
                    </div>

                </div>
            </div>
        </div>


        <div class="col-span-2">
         
            <div class="bg-gray-800 border border-gray-700 shadow-lg rounded-lg p-6">

                <div class="flex justify-start mb-6">
                    <div v-if="website.favicon" class="h-16 w-16 rounded-full overflow-hidden">
                        <img :src="'data:image/x-icon;base64,' + website.favicon" alt="favicon"
                            class="h-full w-full object-cover" />
                    </div>
                    <div v-else
                        class="h-16 w-16 rounded-full bg-gray-800 flex items-center justify-center text-white text-xl">
                        <i class="pi pi-image"></i> <!-- Placeholder icon -->
                    </div>
                </div>

                <div class="grid grid-cols-1 gap-4">
                    <Card title="Website Info"
                        class="shadow-md rounded-lg">
                        <template #title>
                            <h2 class="text-lg font-bold text-white">Website Info</h2>
                        </template>
                        <template #content>
                            <div class="flex flex-col">
                                <div class="flex justify-between mb-2" v-for="(value, key) in website.info" :key="key">
                                    <span class=" font-semibold">{{ key }}:</span>
                                    <span class="">{{ value }}</span>
                                </div>
                            </div>
                        </template>
                    </Card>
                    <Card title="Website Statistics"
                        class="shadow-md rounded-lg">
                        <template #title>
                            <h2 class="text-lg font-bold text-white">Website Statistics</h2>
                        </template>
                        <template #content>
                            <div class="grid grid-cols-2 gap-4">
                                <div class="flex items-center justify-center bg-gray-800 hover:bg-gray-600 transition duration-300 ease-in-out shadow-md rounded-lg p-4">
                                    <div>
                                        <i class="pi pi-folder-plus text-white"></i>
                                        <span class="block mt-2">{{ website.plugins_count }} Plugins</span>
                                    </div>
                                </div>
                                <div class="flex items-center justify-center bg-gray-800 hover:bg-gray-600 transition duration-300 ease-in-out shadow-md rounded-lg p-4">
                                    <div>
                                        <i class="pi pi-palette text-white"></i>
                                        <span class="block mt-2">{{ website.themes_count }} Themes</span>
                                    </div>
                                </div>
                                <div class="flex items-center justify-center bg-gray-800 hover:bg-gray-600 transition duration-300 ease-in-out shadow-md rounded-lg p-4">
                                    <div>
                                        <i class="pi pi-file text-white"></i>
                                        <span class="block mt-2">{{ website.posts_count }} Posts</span>
                                    </div>
                                </div>
                                <div class="flex items-center justify-center bg-gray-800 hover:bg-gray-600 transition duration-300 ease-in-out shadow-md rounded-lg p-4">
                                    <div>
                                        <i class="pi pi-file-o text-white"></i>
                                        <span class="block mt-2">{{ website.pages_count }} Pages</span>
                                    </div>
                                </div>
                                <div class="flex items-center justify-center bg-gray-800 hover:bg-gray-600 transition duration-300 ease-in-out shadow-md rounded-lg p-4">
                                    <div>
                                        <i class="pi pi-users text-white"></i>
                                        <span class="block mt-2">{{ website.users_count }} Users</span>
                                    </div>
                                </div>
                                <div class="flex items-center justify-center bg-gray-800 hover:bg-gray-600 transition duration-300 ease-in-out shadow-md rounded-lg p-4">
                                    <div>
                                        <i class="pi pi-image text-white"></i>
                                        <span class="block mt-2">{{ website.media_count }} Media</span>
                                    </div>
                                </div>
                                <div class="flex items-center justify-center bg-gray-800 hover:bg-gray-600 transition duration-300 ease-in-out shadow-md rounded-lg p-4">
                                    <div>
                                        <i class="pi pi-comment text-white"></i>
                                        <span class="block mt-2">{{ website.comments_count }} Comments</span>
                                    </div>
                                </div>
                            </div>
                        </template>
                       

                    </Card>
                </div>
            </div>
        </div>
       
        <div class="col-span-1">
          
            <DataTable
            stripedRows
       
                ref="dt"
                :value="website.plugins"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.pluginsFilter }"
                filterMode="global"
                class="bg-gray-800
                    stripedRows shadow-md rounded-lg px-6 py-4 mb-3">
                <h2 class="text-lg font-bold text-white">Plugins <i class="pi pi-folder-plus"></i></h2>
                <InputText v-model="filters.pluginsFilter.value" placeholder="Search Plugins" class="mb-4 mt-4 shadow-lg rounded-lg p-2" @input="filterData" />
                
                <Column field="name" header="Name" :sortable="true" class="" />
                <Column field="version" header="Version" :sortable="true" class="" />
                <Column field="author" header="Author" :sortable="true" class="" />
            </DataTable>

            <DataTable
            stripedRows
                ref="dt"
                :value="website.themes"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.themesFilter }"
                filterMode="global"
                class="bg-gray-800
                    stripedRows shadow-md rounded-lg px-6 py-4 mb-3">
                <h2 class="text-lg font-bold text-white">Themes  <i class="pi pi-palette"></i></h2>
                <InputText v-model="filters.themesFilter.value" placeholder="Search Themes" class="mb-4 mt-4 shadow-lg rounded-lg p-2" @input="filterData" />
                
                <Column field="name" header="Name" :sortable="true" class="" />
                <Column field="version" header="Version" :sortable="true" class="" />
                <Column field="author" header="Author" :sortable="true" class="" />
            </DataTable>

            <DataTable
            stripedRows
                ref="dt"
                :value="website.users"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.usersFilter }"
                filterMode="global"
                class="bg-gray-800
                    stripedRows shadow-md rounded-lg px-6 py-4 mb-3">
                <h2 class="text-lg font-bold text-white">Users <i class="pi pi-users"></i></h2>
                <InputText v-model="filters.usersFilter.value" placeholder="Search Users" class="mb-4 mt-4 shadow-lg rounded-lg p-2" @input="filterData" />
                
                <Column field="user_login" header="Login" :sortable="true" class="" />
                <Column field="user_email" header="Email" :sortable="true" class="" />
                <Column field="user_registered" header="Registered" :sortable="true" class="" />
            </DataTable>

            <DataTable
            stripedRows
                ref="dt"
                :value="website.media"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.mediaFilter }"
                filterMode="global"
                class="bg-gray-800
                    stripedRows shadow-md rounded-lg px-6 py-4 mb-3">
                <h2 class="text-lg font-bold text-white">Media <i class="pi pi-image"></i></h2>
                <InputText v-model="filters.mediaFilter.value" placeholder="Search Media" class="mb-4 mt-4 shadow-lg rounded-lg p-2" @input="filterData" />
               <!-- if media type image/jpg or image/png show column with image-->
                <Column field="preview" header="Preview" :sortable="true" class="">
                    <template #body="{ data }">
                        <img v-if="data.media_type === 'image/jpg' || data.media_type === 'image/png'" :src="data.media_url" alt="media" class="h-8 w-8 rounded-full" />
                        <img v-else src="" alt="empty" class="h-8 w-8 rounded-full" />
                    </template>
                  
                </Column>

                <Column field="media_title" header="Title" :sortable="true" class="" />
                <Column field="media_date" header="Date" :sortable="true" class="" />
                <Column field="media_type" header="Type" :sortable="true" class="" />
               
            </DataTable>


            <DataTable
            stripedRows
                ref="dt"
                :value="website.comments"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.commentsFilter }"
                filterMode="global"
                class="bg-gray-800
                    stripedRows shadow-md rounded-lg px-6 py-4 mb-3">
                <h2 class="text-lg font-bold text-white">Comments <i class="pi pi-comment"></i></h2>
                <InputText v-model="filters.commentsFilter.value" placeholder="Search Comments" class="mb-4 mt-4 shadow-lg rounded-lg p-2" @input="filterData" />
                
                <Column field="comment_author" header="Author" :sortable="true" class="" />
                <Column field="comment_date" header="Date" :sortable="true" class="" />
                <Column field="comment_content" header="Content" :sortable="true" :truncate="true" class="" />

            </DataTable>

           
            <DataTable
                stripedRows
                ref="dt"
                :value="website.posts"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.global }"
                filterMode="global"
                class="bg-gray-800
                    stripedRows shadow-md rounded-lg px-6 py-4 mb-3"
            >
                <h2 class="text-lg font-bold text-white">Posts <i class="pi pi-file"></i></h2>
                <InputText
                    v-model="filters.global.value"
                    placeholder="Search Posts"
                    class="mb-4 mt-4 shadow-lg rounded-lg p-2"
                    @input="filterData"
                />

                <Column field="post_title" header="Title" :sortable="true" class="" />
                <Column field="post_date" header="Date" :sortable="true" class="" />

            <Column field="post_status" header="Status" :sortable="true" class="">
                <template #body="{ data }">
                    <span :class="{'bg-green-500': data.post_status === 'publish', 'bg-red-500': data.post_status !== 'publish'}" class="px-2 py-1 rounded-lg text-white">
                        {{ data.post_status === 'publish' ? 'Published' : data.post_status }}
                    </span>
                </template>

            </Column>
            </DataTable>
            

            <DataTable
            stripedRows
                ref="dt"
                :value="website.pages"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.global }"
                filterMode="global"
                class="bg-gray-800
                    stripedRows shadow-md rounded-lg px-6 py-4 mb-3">
                <h2 class="text-lg font-bold text-white">Pages <i class="pi pi-file-o"></i></h2>
                <InputText v-model="filters.global.value" placeholder="Search Pages" class="mb-4 mt-4 shadow-lg rounded-lg p-2" @input="filterData" />
                
                <Column field="page_title" header="Title" :sortable="true" class="" />
                <Column field="page_date" header="Date" :sortable="true" class="" />
                <Column field="page_status" header="Status" :sortable="true" class="" />
            </DataTable>
        </div>
    </div>
</div>
</template>






<script>
import axios from "axios";
import { onMounted } from "vue";
import { FilterMatchMode, FilterOperator } from 'primevue/api';
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Badge from "primevue/badge";
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
            }
        };
    },
    methods: {

        filterData() {
            this.$refs.dt.filter(this.filters.global.value, 'name', this.filters.global.matchMode);
        },
 
        getWebsiteDetail() {
            // Show the progress bar
            var progressBar = document.querySelector(".p-progressbar");
            progressBar.style.display = "block";

            axios.get(`http://localhost:5001/website/${this.id}`)
                .then((response) => {
                    this.website = response.data.website;
                    console.log(this.website.updates);

                    // Update the counts
                    this.website.plugins_count = this.website.plugins.length;
                    this.website.themes_count = this.website.themes.length;
                    this.website.posts_count = this.website.posts.length;
                    this.website.pages_count = this.website.pages.length;
                    this.website.users_count = this.website.users.length;
                    this.website.media_count = this.website.media.length;
                    this.website.comments_count = this.website.comments.length;

                    // Hide the progress bar
                    progressBar.style.display = "none";
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
/* add PrimeVue styles */



</style>
