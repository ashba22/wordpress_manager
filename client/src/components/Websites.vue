<template>
  <div class="container mx-auto">
    <div class="flex flex-col">
      <h1 class="text-3xl font-bold mb-4">WordPress Manager</h1>
      <hr class="my-4" />
      <button
        type="button"
        class="bg-green-500 text-white px-4 py-2 rounded-sm mb-4 shadow-md hover:bg-green-600 transition-colors duration-300 w-32"
        @click="toggleAddWebsiteModal">
        Add Website
      </button>

      <table class="table-auto w-full bg-gray-200 p-8">
        <thead>
          <tr class="bg-gray-100">
            <th scope="col" @click="handleSort('name')" class="cursor-pointer text-left">
              Name
              <i v-if="sortBy === 'name'" :class="reverse ? 'fas fa-sort-up' : 'fas fa-sort-down'"></i>
            </th>
            <th scope="col" @click="handleSort('url')" class="cursor-pointer text-left">
              URL
              <i v-if="sortBy === 'url'" :class="reverse ? 'fas fa-sort-up' : 'fas fa-sort-down'"></i>
            </th>
            <th scope="col" @click="handleSort('active')" class="cursor-pointer text-left">
              Status
              <i v-if="sortBy === 'active'" :class="reverse ? 'fas fa-sort-up' : 'fas fa-sort-down'"></i>
            </th>
            <th class="text-left">Action</th>
          </tr>
        </thead>
        <tbody class="bg-white p-4">
          <tr v-for="(website, index) in websites" :key="index">
            <td>{{ website.name }}</td>
            <td>{{ website.url }}</td>
            <td>
              <span v-if="website.active" class="text-green-500">Active</span>
              <span v-else class="text-red-500">Inactive</span>
            </td>
            <td>
              <div class="flex items-center">
             
                <router-link
                :to="`/website/${website.id}`"
                class="bg-blue-500 text-white px-4 py-2 rounded-sm mr-2 transition-colors duration-300 hover:bg-blue-600"
                >
                View
                </router-link>
              <button
                class="bg-yellow-500 text-white px-4 py-2 rounded-sm mr-2 transition-colors duration-300 hover:bg-yellow-600"
                @click="toggleEditWebsiteModal(website)"
              >
                Update
              </button>
              <button
                class="bg-red-500 text-white px-4 py-2 rounded-sm transition-colors duration-300 hover:bg-red-600"
                @click="handleDeleteWebsite(website)"
              >
                Delete
              </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

      <!-- Add Website Modal -->
      <div
        ref="addWebsiteModal"
        class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
        :class="{ 'hidden': !activeAddWebsiteModal }"
      >
        <div class="bg-white p-4 rounded shadow-lg">
        <h5 class="text-xl font-bold mb-4">Add Website</h5>
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
          <label for="addWebsiteName" class="block font-bold mb-2">Name</label>
          <input
            type="text"
            class="border border-gray-300 rounded px-4 py-2 w-full"
            id="addWebsiteName"
            v-model="addWebsiteForm.name"
            required
          />
          </div>
          <div class="mb-4">
          <label for="addWebsiteUrl" class="block font-bold mb-2">URL</label>
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
          <label class="ml-2" for="addWebsiteActive">Active</label>
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
        <div class="bg-white p-4 rounded shadow-lg">
        <h5 class="text-xl font-bold mb-4">Edit Website</h5>
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
          <label for="editWebsiteName" class="block font-bold mb-2">Name</label>
          <input
            type="text"
            class="border border-gray-300 rounded px-4 py-2 w-full"
            id="editWebsiteName"
            v-model="editWebsiteForm.name"
            required
          />
          </div>
          <div class="mb-4">
          <label for="editWebsiteUrl" class="block font-bold mb-2">URL</label>
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
          <label class="ml-2" for="editWebsiteActive">Active</label>
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
