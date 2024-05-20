<template>
    <!-- TODO: Create Modules for each element-->
    <div class="d-flex flex-row flex-wrap">


        
        <div class="grid mt-2 mx-auto surface-0">
            <div class="col-12 lg:col-6">
                <div class="grid">
                    <div class="col mb-2">
                        <h1>Website: {{ website.name }}</h1>
                        <div class="flex gap-2">
                            <Button type="button" key="view" class="p-button-success"
                                :href="website.url" target="_blank" rel="noopener noreferrer">
                                <i class="pi pi-external-link mr-2"></i> View Website
                            </Button>

                            <Button type="button" key="edit" class="p-button-warning"
                                :href="`/website/${website.id}`">
                                <i class="pi pi-pencil mr-2"></i> Edit Website
                            </Button>

                            <Button type="button" key="admin" class="p-button-secondary"
                                :href="website.url + '/wp-admin'">
                                <i class="pi pi-cog mr-2"></i> WP Admin
                            </Button>
                            
                        </div>
                    </div>
                </div>
                <div class="grid gap-2 p-2 rounded-3">
                    <div v-for="(item, index) in [
                        { icon: 'pi pi-folder-plus', label: 'Plugins', count: website.plugins_count },
                        { icon: 'pi pi-palette', label: 'Themes', count: website.themes_count },
                        { icon: 'pi pi-file', label: 'Posts', count: website.posts_count },
                        { icon: 'pi pi-file-o', label: 'Pages', count: website.pages_count },
                        { icon: 'pi pi-users', label: 'Users', count: website.users_count },
                        { icon: 'pi pi-image', label: 'Media', count: website.media_count },
                        { icon: 'pi pi-comment', label: 'Comments', count: website.comments_count }
                    ]" :key="index" class="flex flex-column align-items-center justify-content-center p-3 hover:surface-100 transition-duration-2s cursor-pointer">
                        <i :class="[item.icon, 'text-2xl']"></i>
                        <span class="mt-1 text-sm">{{ item.label }}: {{ item.count }}</span>
                    </div>
                </div>
            </div>
            
            <div class="col-12 lg:col-6">
                <h1 class="text-2xl font-bold">Website Info</h1>
                <div class="grid gap-3 rounded-3 p-2">
                  
                    <div v-for="(value, key) in website.info" :key="key" class="flex justify-content-between align-items-center p-2 border-round surface-100">
                        <span class="font-bold">{{ key }}:</span>
                        <Badge :value="value" class="ml-2" />
                    </div>
                </div>
            </div>
        </div>

        <div class="surface-50 py-2 px-4">
            <h2 class="my-2">
            Updates <i class="pi pi-arrow-up"></i>
            </h2>
            <hr>

            <div class="grid">
                <!-- WordPress Core Update Section -->
                <div class="col-12 md:col-4">
                    <div v-if="website.updates && website.updates.core.length > 0" class="card rounded shadow ">
                    <div v-for="update in website.updates.core" :key="update.version" class="flex flex-col align-items-center">
                        <h2 class="text-xl font-bold text-primary">
                        WordPress Core Update: {{ update.version }}
                        </h2>
                    </div>
                    </div>
                    <div v-else class="card rounded shadow ">
                    <h3 class="text-center text-muted">No WordPress Core updates available.</h3>
                    </div>
                </div>

                <!-- Plugins Update Section -->
                <div class="col-12 md:col-4">
                    <div v-if="plugins_update_bool" class="card rounded shadow">
                    <div class="flex align-items-center">
                        <i class="pi pi-folder-plus text-2xl text-primary ml-2"></i>
                        <h3 class="ml-2 ">Plugins Updates: {{ website.updates.plugins_update_count }}</h3>
                    </div>
                    </div>
                    <div v-else class="card rounded shadow ">
                    <h3 class="text-center text-muted">No plugin updates available.</h3>
                    </div>
                </div>

                <!-- Themes Update Section -->
                <div class="col-12 md:col-4">
                    <div v-if="website.updates && themes_update_bool" class="card rounded shadow ">
                    <div class="flex align-items-center">
                        <i class="pi pi-palette text-2xl text-purple"></i>
                        <h3 class="ml-2">Themes Updates: {{ website.updates.themes_update_count }}</h3>
                    </div>
                    </div>
                    <div v-else class="card rounded shadow ">
                    <h3 class="text-center text-muted">No theme updates available.</h3>
                    </div>
                </div>
                
            </div>
        </div>
        
        <div class="col-12">
            <DataTable
                ref="dt"
                :value="website.plugins"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.pluginsFilter }"
                filterMode="global"
                id="plugins-table"
                class="p-datatable-gridlines"
            >
                <div class="d-flex jc-between items-center">
                    <div class="d-flex jc-between items-center">
                        <h2 class="text-1xl font-bold space-x-2">
                            Plugins <i class="pi pi-folder-plus"></i>
                        </h2>
                    </div>
                    <InputText v-model="filters.pluginsFilter.value" placeholder="Search Plugins" />
                </div>
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
                        <div class="d-flex items-center">
                            <span
                                v-if="data.status === 'Active'"
                                class="bg-green-500 px-2 py-1 rounded-lg"
                                @click="togglePluginByName(data.name, data.path)"
                            >
                                Active
                            </span>
                            <span
                                v-else
                                @click="togglePluginByName(data.name, data.path)"
                                class="bg-red-500 px-2 py-1 rounded-lg"
                            >
                                Inactive
                            </span>
                        </div>
                    </template>
                </Column>
            </DataTable>
        </div>

        <div class="col-12">
            <DataTable
                ref="dt"
                :value="website.themes"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.themesFilter }"
                filterMode="global"
                id="themes-table"
                class="p-datatable-gridlines"
            >
                <div class="d-flex jc-between items-center">
                    <h2 class="text-1xl font-bold space-x-2">Themes <i class="pi pi-palette"></i></h2>
                    <InputText v-model="filters.themesFilter.value" placeholder="Search Themes" />
                </div>
                <Column field="name" header="Name" :sortable="true" />
                <Column field="version" header="Version" :sortable="true" />
                <Column field="author" header="Author" :sortable="true" />
            </DataTable>
        </div>

        <div class="col-12">
            <DataTable
                ref="dt"
                :value="website.users"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.usersFilter }"
                filterMode="global"
                id="users-table"
                class="p-datatable-gridlines"
            >
                <div class="d-flex jc-between items-center">
                    <h2 class="text-1xl font-bold space-x-2">Users <i class="pi pi-users"></i></h2>
                    <InputText v-model="filters.usersFilter.value" placeholder="Search Users" />
                </div>
                <Column field="user_login" header="Login" :sortable="true" />
                <Column field="user_email" header="Email" :sortable="true" />
                <Column field="user_registered" header="Registered" :sortable="true" />
            </DataTable>
        </div>

        <div class="col-12">
            <DataTable
                ref="dt"
                :value="website.media"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.mediaFilter }"
                id="media-table"
                filterMode="global"
                class="p-datatable-gridlines"
            >
                <div class="d-flex jc-between items-center">
                    <h2 class="text-1xl font-bold space-x-2">Media <i class="pi pi-image"></i></h2>
                    <InputText v-model="filters.mediaFilter.value" placeholder="Search Media" />
                </div>
                <Column field="preview" header="Preview" :sortable="true">
                    <template #body="{ data }">
                        <img
                            v-if="data.media_type === 'image/jpg' || data.media_type === 'image/png'"
                            :src="data.media_url"
                            alt="media"
                        />
                        <i v-else class="pi pi-file text-2xl"></i>
                    </template>
                </Column>
                <Column field="media_title" header="Title" :sortable="true" />
                <Column field="media_date" header="Date" :sortable="true" />
                <Column field="media_type" header="Type" :sortable="true" />
            </DataTable>
        </div>

        <div class="col-12">
            <DataTable
                ref="dt"
                :value="website.comments"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.commentsFilter }"
                filterMode="global"
                id="comments-table"
                class="p-datatable-gridlines"
            >
                <div class="d-flex jc-between items-center">
                    <h2 class="text-1xl font-bold space-x-2">Comments <i class="pi pi-comment"></i></h2>
                    <InputText v-model="filters.commentsFilter.value" placeholder="Search Comments" />
                </div>
                <Column field="comment_author" header="Author" :sortable="true" />
                <Column field="comment_date" header="Date" :sortable="true" />
                <Column field="comment_content" header="Content" :sortable="true" :truncate="true" />
            </DataTable>
        </div>

        <div class="col-12">
            <DataTable
                ref="dt"
                :value="website.posts"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.postsFilter }"
                filterMode="global"
                id="posts-table"
                class="p-datatable-gridlines"
            >
                <div class="d-flex jc-between items-center">
                    <h2 class="text-1xl font-bold space-x-2">Posts <i class="pi pi-file"></i></h2>
                    <InputText v-model="filters.postsFilter.value" placeholder="Search Posts" />
                </div>
                <Column field="post_id" header="ID" :sortable="true" />
                <Column field="post_url" header="URL" :sortable="true">
                    <template #body="{ data }">
                        <a :href="data.post_url" target="_blank" rel="noopener noreferrer">
                            <Button key="view" icon="pi pi-external-link" class="button-info" label="" />
                        </a>
                    </template>
                </Column>
                <Column field="post_image" header="Image" :sortable="true">
                    <template #body="{ data }">
                        <img v-if="data.post_image" :src="data.post_image" alt="post" />
                    </template>
                </Column>
                <Column field="post_title" header="Title" :sortable="true" />
                <Column field="post_date" header="Date" :sortable="true" />
                <Column field="post_status" header="Status" :sortable="true">
                    <template #body="{ data }">
                        <Dropdown
                            v-model="data.post_status"
                            @change="changePostStatusById(data.post_id, data.post_status)"
                            :options="['publish', 'draft', 'pending', 'auto-draft', 'future', 'private', 'inherit', 'trash']"
                        >
                            <template #optiongroup="slotProps">
                                <div class="d-flex ai-center">
                                    <div>{{ slotProps.option }}</div>
                                    <div>{{ slotProps.option.label }}</div>
                                </div>
                            </template>
                        </Dropdown>
                    </template>
                </Column>
            </DataTable>
        </div>

        <div class="col-12">
            <DataTable
                ref="dt"
                :value="website.pages"
                :paginator="true"
                :rows="5"
                :filters="{ global: filters.pagesFilter }"
                filterMode="global"
                id="pages-table"
                class="p-datatable-gridlines"
            >
                <div class="d-flex jc-between items-center">
                    <h2 class="text-1xl font-bold space-x-2">Pages <i class="pi pi-file-o"></i></h2>
                    <InputText v-model="filters.pagesFilter.value" placeholder="Search Pages" />
                </div>
                <Column field="page_url" header="URL" :sortable="true">
                    <template #body="{ data }">
                        <a :href="data.page_url" target="_blank" rel="noopener noreferrer">
                            <Button key="view" icon="pi pi-external-link" class="button-info" />
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
                        <Dropdown
                            v-model="data.page_status"
                            @change="changePageStatusById(data.page_id, data.page_status)"
                            :options="['publish', 'draft', 'pending', 'auto-draft', 'future', 'private', 'inherit', 'trash']"
                        >
                            <template #optiongroup="slotProps">
                                <div class="d-flex ai-center">
                                    <div>{{ slotProps.option }}</div>
                                    <div>{{ slotProps.option.label }}</div>
                                </div>
                            </template>
                        </Dropdown>
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
import Dropdown from "primevue/dropdown";
import Button from "primevue/button";
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
            try {
                // Show the progress bar
                var progressBar = document.querySelector(".p-progressbar");
                progressBar.style.display = "block";

                const response = await axios.get(`${this.$env.VITE_API_URL}/website/${this.id}`);
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

                this.$toast.open({
                    message: "Website details fetched successfully.", // Show the success message
                    type: "success",
                    position: "top-right",
                });
            } catch (error) {
                console.log(error);

                this.$toast.open({
                    message: "Failed to fetch website details.", // Show the error message
                    type: "error",
                    position: "top-right",
                });
            } finally {
                // Hide the progress bar
                var progressBar = document.querySelector(".p-progressbar");
                progressBar.style.display = "none";
            }
        },
        togglePluginByName(name, path) {
            const VITE_API_URL = `${this.$env.VITE_API_URL}/website/${this.id}/toggle-plugin`;
            axios.post(VITE_API_URL, { name: name, path: path })
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
        },
        changePostStatusById(id, status) {
            const VITE_API_URL = `${this.$env.VITE_API_URL}/website/${this.id}/change-post-status`;
            axios.post(VITE_API_URL, { id: id, status: status })
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
        },
        changePageStatusById(id, status) {
            const VITE_API_URL = `${this.$env.VITE_API_URL}/website/${this.id}/change-page-status`;
            axios.post(VITE_API_URL, { id: id, status: status })
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
    components: {
  
    }
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
    margin-bottom: 1rem;
}

img {
    width: 100px;
    height: 100px;
    object-fit: cover;
    border-radius: 0.5rem;
}
</style>