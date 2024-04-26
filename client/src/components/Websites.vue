<template>
  <div class="container-fluid">
    <div class="flex flex-col">

      <div class="flex justify-between items-center p-4">
        <h2 class="text-2xl font-bold">Websites</h2>
        <button
          class="bg-blue-500 text-white px-4 py-2 rounded-sm transition-colors duration-300 hover:bg-blue-600"
          @click="toggleAddWebsiteModal"
        >
          Add Website
        </button>
      </div>
      <div class="card">
        <DataTable :value="websites" :paginator="true" :rows="10" :rowsPerPageOptions="[5, 10, 20]">
            <Column field="name" header="Name" sortable></Column>
            <Column field="url" header="URL" sortable></Column>
            <Column field="active" header="Status" sortable></Column>
            <Column header="Action">
                <template #body="slotProps">
                    <router-link :to="`/website/${slotProps.data.id}`"> 
                      <Button label="MANAGE" class="bg-blue-500 text-white px-4 py-2 rounded-sm transition-colors duration-300 hover:bg-blue-600"></Button>
                    </router-link>
                    <Button label="EDIT" @click="toggleEditWebsiteModal(slotProps.data)" class="bg-yellow-500 text-white px-4 py-2 rounded-sm transition-colors duration-300 hover:bg-yellow-600"></Button>
                    <Button label="Delete"  @click="handleDeleteWebsite(slotProps.data)" class="bg-red-500 text-white px-4 py-2 rounded-sm transition-colors duration-300 hover:bg-red-600"></Button>
                  </template>
            </Column>
        </DataTable>
      </div>

    </div>


      <!-- Add Website Modal -->
      <div
        ref="addWebsiteModal"
        class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
        :class="{ 'hidden': !activeAddWebsiteModal }"
      >
        <div class="bg-gray-800 p-4 rounded shadow-lg">
        <h5 class="text-xl font-bold mb-4 text-white">Add Website</h5>
        <button
          type="button"
          class="absolute top-2 right-2 text-gray-500 hover:text-gray-700"
          @click="toggleAddWebsiteModal"
        >
          <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          >
          <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M6 18L18 6M6 6l12 12"
          />
          </svg>
        </button>
        <form @submit.prevent="handleAddSubmit">
          <div class="mb-4">
          <label for="addWebsiteName" class="block font-bold mb-2 text-white">Name</label>
          <input
        type="text"
        class="border border-gray-300 rounded px-4 py-2 w-full"
        id="addWebsiteName"
        v-model="addWebsiteForm.name"
        required
          />
          </div>
          <div class="mb-4">
          <label for="addWebsiteUrl" class="block font-bold mb-2 text-white">URL</label>
          <input
        type="text"
        class="border border-gray-300 rounded px-4 py-2 w-full"
        id="addWebsiteUrl"
        v-model="addWebsiteForm.url"
        required
          />
          </div>
          <div class="mb-4">
          <input
        type="checkbox"
        class="form-checkbox"
        id="addWebsiteActive"
        v-model="addWebsiteForm.active"
          />
          <label class="ml-2 text-white" for="addWebsiteActive">Active</label>
          </div>
          <div class="flex justify-end">
          <button type="button" class="text-gray-500 hover:text-gray-700 mr-2" @click="toggleAddWebsiteModal">Cancel</button>
          <button type="submit" class="bg-blue-500 text-white px-4 py-2 rounded-sm">Submit</button>
          </div>
        </form>
        </div>
      </div>

      <!-- Edit Website Modal -->
      <div
        ref="editWebsiteModal"
        class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
        :class="{ 'hidden': !activeEditWebsiteModal }"
      >
        <div class="bg-gray-800 p-4 rounded shadow-lg">
        <h5 class="text-xl font-bold mb-4 text-white">Edit Website</h5>
        <button
          type="button"
          class="absolute top-2 right-2 text-gray-500 hover:text-gray-700"
          @click="toggleEditWebsiteModal"
        >
          <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          >
          <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M6 18L18 6M6 6l12 12"
          />
          </svg>
        </button>
        <form @submit.prevent="handleEditSubmit">
          <div class="mb-4">
          <label for="editWebsiteName" class="block font-bold mb-2 text-white">Name</label>
          <input
        type="text"
        class="border border-gray-300 rounded px-4 py-2 w-full"
        id="editWebsiteName"
        v-model="editWebsiteForm.name"
        required
          />
          </div>
          <div class="mb-4">
          <label for="editWebsiteUrl" class="block font-bold mb-2 text-white">URL</label>
          <input
        type="text"
        class="border border-gray-300 rounded px-4 py-2 w-full"
        id="editWebsiteUrl"
        v-model="editWebsiteForm.url"
        required
          />
          </div>
          <div class="mb-4">
          <input
        type="checkbox"
        class="form-checkbox"
        id="editWebsiteActive"
        v-model="editWebsiteForm.active"
          />
          <label class="ml-2 text-white" for="editWebsiteActive">Active</label>
          </div>
          <div class="flex justify-end">
          <button type="button" class="text-gray-500 hover:text-gray-700 mr-2" @click="toggleEditWebsiteModal">Cancel</button>
          <button type="submit" class="bg-blue-500 text-white px-4 py-2 rounded-sm">Submit</button>
          </div>
        </form>
        </div>
      </div>
      </div>


</template>


<script>
import axios from "axios";


export default {
  data() {
    return {
      activeAddWebsiteModal: false,
      activeEditWebsiteModal: false,
      addWebsiteForm: {
        name: "",
        url: "",
        active: false,
      },
      websites: [],
      editWebsiteForm: {
        id: "",
        name: "",
        url: "",
        active: false,
      },
      sortBy: "name", // default sorting by name
      reverse: false,
    };
  },
  components: {
    // Alert,
  },
  methods: {
    // Fetch websites with sorting
    fetchWebsites(sortBy, reverse) {
      const apiUrl = `http://localhost:5001/websites?sort_by=${sortBy}&reverse=${reverse}`;
      axios
        .get(apiUrl)
        .then((res) => {
          this.websites = res.data.websites;
        })
        .catch((error) => {
          console.error(error);
        });
    },
    // Handle sorting
    handleSort(sortBy) {
      if (sortBy === this.sortBy) {
        this.reverse = !this.reverse;
      } else {
        this.sortBy = sortBy;
        this.reverse = false;
      }
      this.fetchWebsites(this.sortBy, this.reverse);
    },
    addWebsite(payload) {
      const path = "http://localhost:5001/websites";
      axios
        .post(path, payload)
        .then(() => {
          this.fetchWebsites(this.sortBy, this.reverse);
          this.message = "Website added!";
          this.$toast.open({
            message: this.message,
            type: "success",
            position: "top-right",
            
          });

          this.toggleAddWebsiteModal();
        })
        .catch((error) => {
          console.log(error);
          this.fetchWebsites(this.sortBy, this.reverse);
        });
    },
    handleAddSubmit() {
      const payload = {
        name: this.addWebsiteForm.name,
        url: this.addWebsiteForm.url,
        active: this.addWebsiteForm.active,
      };
      this.addWebsite(payload);
      this.initForm();
    },
    handleDeleteWebsite(website) {
      this.removeWebsite(website.id);
    },
    handleEditSubmit() {
      const payload = {
        name: this.editWebsiteForm.name,
        url: this.editWebsiteForm.url,
        active: this.editWebsiteForm.active,
      };
      this.updateWebsite(payload, this.editWebsiteForm.id);
    },
    initForm() {
      this.addWebsiteForm.name = "";
      this.addWebsiteForm.url = "";
      this.addWebsiteForm.active = false;
      this.editWebsiteForm.id = "";
      this.editWebsiteForm.name = "";
      this.editWebsiteForm.url = "";
      this.editWebsiteForm.active = false;
    },
    removeWebsite(websiteID) {
      const path = `http://localhost:5001/website/${websiteID}`;
      axios
        .delete(path)
        .then(() => {
          this.fetchWebsites(this.sortBy, this.reverse);
          this.message = "Website deleted!";
          this.$toast.open({
            message: this.message,
            type: "info",
            position: "top-right",
          });

        })
        .catch((error) => {
          console.error(error);
          this.fetchWebsites(this.sortBy, this.reverse);
        });
    },
    toggleAddWebsiteModal() {
      this.activeAddWebsiteModal = !this.activeAddWebsiteModal;
    },
    toggleEditWebsiteModal(website) {
      this.activeEditWebsiteModal = !this.activeEditWebsiteModal;
      if (website) {
        this.editWebsiteForm = website;
      }
    },
    updateWebsite(payload, websiteID) {
      const path = `http://localhost:5001/website/${websiteID}`;
      axios
        .put(path, payload)
        .then(() => {
          this.fetchWebsites(this.sortBy, this.reverse);
          this.message = "Website updated!";
          this.$toast.open({
            message: this.message,
            type: "success",
            position: "top-right",
          });
          this.toggleEditWebsiteModal();
        })
        .catch((error) => {
          console.error(error);
          this.fetchWebsites(this.sortBy, this.reverse);
        });
    },
  },
  created() {
    this.fetchWebsites(this.sortBy, this.reverse);
  },
  //// when esc cicked clode modal
  mounted() {
    window.addEventListener("keydown", (e) => {
      if (e.key === "Escape") {
        this.activeAddWebsiteModal = false;
        this.activeEditWebsiteModal = false;
      }
    });
  },
  
};
</script>

<style scoped>
.card {
  background-color: #1a202c;
  border-radius: 0.5rem;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  margin: 1rem;
  padding: 1rem;
}
.button-success {
  background-color: #2b6cb0;
  border-radius: 0.25rem;
  color: white;
  padding: 0.5rem 1rem;
  text-decoration: none;
}
</style>
