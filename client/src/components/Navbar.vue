
<template>
  <div class="card">
      <Menubar :model="items" :start="start" :item="item" :end="end" class="shadow-lg">
        <template #start>
          <div class="flex items-center">
            <img src="/logo.svg" alt="logo" class="w-16 h-16" />
            <span class="ml-2 text-color-white font-bold"><b>WPManager</b></span>
          </div>
        </template>
        <template #item="{ item, props, hasSubmenu }">
          <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
            <a v-ripple :href="href" v-bind="props.action" @click="navigate" class="flex items-center">
              <span :class="item.icon" />
              <span class="ml-2">{{ item.label }}</span>
            </a>
          </router-link>
          <a v-else v-ripple :href="item.url" :target="item.target" v-bind="props.action" class="flex items-center">
            <span :class="item.icon" />
            <span class="ml-2">{{ item.label }}</span>
            <span v-if="hasSubmenu" class="pi pi-fw pi-angle-down ml-2" />
          </a>
        </template>

      </Menubar>
  </div>
</template>

<script setup>
import { ref } from "vue";

const items = ref([
  {
    label: 'Home',
    icon: 'pi pi-home',
    route: '/'
  },
]);
</script>

<style scoped>
p-menubar {
border: 1px solid #e0e0e0;
}
</style>
