<template>

  <div class="main-content">

  
    <div class="grid">

     
   
      
      <div class="content col-12 lg:col-8">
        <div class="d-flex jc-between ai-center p-2">
          <h2 class="text-2xl">Websites</h2>
          <Button label="Add Website" @click="showAddWebsiteModal=true" class="button-raised button-rounded button-info button-text" icon="pi pi-plus"></Button>
        </div>
        <div class="mt-2">
          <DataTable :value="websites" :paginator="true" :rows="2" :rowsPerPageOptions="[2, 10, 20]" class="p-datatable-gridlines p-datatable-responsive">
            <Column field="name" header="Name" sortable></Column>
            <Column field="url" header="URL" sortable></Column>
            <Column header="Action">
              <template #body="slotProps">
                <div class="flex gap-2">
                  <router-link :to="`/website/${slotProps.data.id}`">
                    <Button label="MANAGE" icon="pi pi-cog" severity="info" class="button-rounded button-text"></Button>
                  </router-link>
                  <Button label="EDIT" icon="pi pi-pencil" @click="editWebsite(slotProps.data)" severity="warning" class="button-rounded button-text"></Button>  
                  <Button label="DELETE" icon="pi pi-trash" @click="confirmDelete(slotProps.data)" severity="danger" class="button-rounded button-text"></Button>
                </div>
              </template>
            </Column>
            <Column>
              <template #header>
                <div class="d-flex ai-center">
                  <span class="mr-2">Status</span>
                  <Button icon="pi pi-refresh" class="p-button-rounded p-button-info " @click="fetchWebsites"></Button>
                </div>
          
     
              </template>
              <template #body="slotProps">
                <div class="d-flex ai-center space-x-2">
                  <span v-if="slotProps.data.is_online" class="text-green-500 font-bold">Online</span>
                  <span v-else class="text-red-500 font-bold">Offline</span>
                </div>
              </template>
            </Column>
          </DataTable>
        </div>
        <!-- Add Website Dialog -->
        <Dialog header="Add Website" :visible.sync="showAddWebsiteModal" modal style="width: 50vw" class="">
          <form @submit.prevent="handleAddSubmit">
            <div class="field grid">
              <label for="name" class="col-fixed" style="width: 100px">Name</label>
              <div class="col">
                <InputText id="name" v-model="addWebsiteForm.name" required />
              </div>
            </div>
            <div class="field grid">
              <label for="url" class="col-fixed" style="width: 100px">URL</label>
              <div class="col">
                <InputText id="url" v-model="addWebsiteForm.url" required />
              </div>
            </div>
            <div class="field grid">
              <label for="api_key" class="col-fixed" style="width: 100px">API Key</label>
              <div class="col">
                <InputText id="api_key" v-model="addWebsiteForm.api_key" required />
              </div>
            </div>
            <div class="field grid">
              <label for="active" class="col-fixed" style="width: 100px">Active</label>
              <div class="col">
                <Checkbox id="active" v-model="addWebsiteForm.active" binary="true" class="mt-0" />
              </div>
            </div>
            <div class="d-flex jc-between mt-4">
              <Button label="Cancel" icon="pi pi-times" class="button-text" @click="showAddWebsiteModal=false"></Button>
              <Button type="submit" label="Submit" icon="pi pi-check" class="button-raised button-rounded"></Button>
            </div>
          </form>
        </Dialog>

        <!-- Edit Website Dialog -->
        <Dialog header="Edit Website" :visible.sync="showEditWebsiteModal" modal style="width: 50vw" class="">
          <form @submit.prevent="handleEditSubmit" class="d-flex flex-column gap-1">
            <div class="field grid">
              <label for="name" class="col-fixed" style="width: 100px">Name</label>
              <div class="col">
                <InputText id="name" v-model="editWebsiteForm.name" required />
              </div>
            </div>
            <div class="field grid">
              <label for="url" class="col-fixed" style="width: 100px">URL</label>
              <div class="col">
                <InputText id="url" v-model="editWebsiteForm.url" required />
              </div>
            </div>
            <div class="field grid">
              <label for="api_key" class="col-fixed" style="width: 100px">API Key</label>
              <div class="col">
                <InputText id="api_key" v-model="editWebsiteForm.api_key" required />
              </div>
            </div>
            <div class="field grid">
              <label for="active" class="col-fixed" style="width: 100px">Active</label>
              <div class="col">
                <Checkbox id="active" v-model="editWebsiteForm.active" binary="true" class="mt-0" />
              </div>
            </div>
            <div class="d-flex jc-between mt-4">
              <Button label="Cancel" icon="pi pi-times" class="button-text" @click="showEditWebsiteModal=false"></Button>
              <Button type="submit" label="Submit" icon="pi pi-check" class="button-raised button-rounded"></Button>
            </div>
          </form>
        </Dialog>
      </div>

      <Card class="mb-1 col-12 lg:col-4">
        <template #header>
          <p class="text-xl font-bold text-center p-3 border-bottom-1">Websites Stats and Updates</p>
        </template>
        <template #content>
          <div class="grid">
            <div class="col-12 md:col-4 text-center mb-4">
              <div class="flex flex-col align-items-center">
                <h3 class="font-bold">Websites</h3>
                <span class="text-2xl font-bold ml-2 text-blue-500">{{ websites.length }}</span>
              </div>
            </div>
            <div class="col-12 md:col-4 text-center mb-4">
              <div class="flex flex-col align-items-center">
                <h3 class="text-lg font-bold">Online</h3>
                <span class="text-2xl font-bold ml-2 text-green-500">{{ websites.filter(website => website.is_online).length }}</span>
              </div>
            </div>
            <div class="col-12 md:col-4 text-center mb-4">
              <div class="flex flex-col align-items-center">
                <h3 class="text-lg font-bold">Offline</h3>
                <span class="text-2xl font-bold ml-2 text-red-500">{{ websites.filter(website => !website.is_online).length }}</span>
              </div>
            </div>
            <div class="col-12 md:col-4 text-center mb-4">
              <div class="flex flex-col align-items-center">
                <h3 class="text-lg font-bold">WP Updates</h3>
                <span class="text-2xl font-bold ml-2">{{ websites_updates.reduce((total, website) => total + (website.updates.core !== '' ? 1 : 0), 0) }}</span>
              </div>
            </div>
            <div class="col-12 md:col-4 text-center mb-4">
              <div class="flex flex-col align-items-center">
                <h3 class="text-lg font-bold">Plugins Updates</h3>
                <span class="text-2xl font-bold ml-2">{{ websites_updates.reduce((total, website) => total + website.updates.plugins_update_count, 0) }}</span>
              </div>
            </div>
            <div class="col-12 md:col-4 text-center mb-4">
              <div class="flex flex-col align-items-center">
                <h3 class="text-lg font-bold">Template Updates</h3>
                <span class="text-2xl font-bold ml-2">{{ websites_updates.reduce((total, website) => total + website.updates.themes_update_count, 0) }}</span>
              </div>
            </div>
          </div>
        </template>
        <template #footer>
          <div class="flex justify-center p-2">
            <Button label="Refresh Stats" icon="pi pi-refresh" class="p-button-rounded p-button-info" @click="fetchWebsitesUpdates" />
          </div>
        </template>
      </Card>
      
    </div>
  </div>

</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      showAddWebsiteModal: false,
      showEditWebsiteModal: false,
      websites: [],
      websites_updates: [],
      addWebsiteForm: { name: '', url: '', api_key: '', active: false },
      editWebsiteForm: { id: '', name: '', url: '', api_key: '', active: false },
      
    };
  },
  components: {},
  methods: {
    async fetchWebsites() {
      console.log('fetching websites');
      console.log(this.$env.VITE_API_URL);
      try {
        this.showLoadingBar();
        const response = await axios.get(`${this.$env.VITE_API_URL}/websites`);
        this.websites = response.data;
        this.hideLoadingBar();
        this.$toast.open({
          message: "Websites Loaded",
          type: "info",
          position: "top-right",
        });
      } catch (error) {
        console.error(error);
        this.hideLoadingBar();
        this.$toast.open({
          message: "Error Websites",
          type: "error",
          position: "top-right",
        });
      }
    },
    async fetchWebsitesUpdates() {
      try {
        this.showLoadingBar();
        const response = await axios.get(`${this.$env.VITE_API_URL}/websites/updates`);
        this.websites_updates = response.data;
        this.hideLoadingBar();
        this.$toast.open({
          message: "Updates informations loaded",
          type: "info",
          position: "top-right",
        });
      } catch (error) {
        console.error(error);
        this.hideLoadingBar();
        this.$toast.open({
          message: "Error Websites Updates",
          type: "error",
          position: "top-right",
        });
      }
    },
    async handleAddSubmit() {
      try {
        await axios.post(`${this.$env.VITE_API_URL}/websites`, this.addWebsiteForm);
        this.showAddWebsiteModal = false;
        this.fetchWebsites();
      } catch (error) {
        console.error(error);
      }
    },
    async handleEditSubmit() {
      try {
        await axios.put(`${this.$env.VITE_API_URL}/website/${this.editWebsiteForm.id}`, this.editWebsiteForm);
        this.showEditWebsiteModal = false;
        this.fetchWebsites();
      } catch (error) {
        console.error(error);
      }
    },
    editWebsite(website) {
      this.editWebsiteForm = { ...website };
      this.showEditWebsiteModal = true;
    },
    confirmDelete(website) {
      if (confirm(`Are you sure you want to delete ${website.name}?`)) {
        axios.delete(`${this.$env.VITE_API_URL}/website/${website.id}`)
          .then(() => this.fetchWebsites())
          .catch(error => console.error(error));
      }
    },
    showLoadingBar() {
      const progressBar = document.querySelector(".p-progressbar");
      progressBar.style.display = "block";
    },
    hideLoadingBar() {
      const progressBar = document.querySelector(".p-progressbar");
      progressBar.style.display = "none";
    },
  },
  created() {
    this.fetchWebsites();
  },
  mounted() {
    this.fetchWebsitesUpdates();
  }
};
</script>
<style scoped>

.main-content {
  padding: 1rem;
  margin: 0 auto;

}


</style>