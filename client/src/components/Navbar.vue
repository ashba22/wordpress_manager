
<template>
  <div class="card"> 
      <Menubar :model="items" class="bg-gray-900 text-color-white" :start="start" :item="item" :end="end">
        <template #start>
          <div class="flex items-center">
            <img src="/logo.svg" alt="logo" class="w-16 h-16 rounded-full shadow-lg bg-gray-800 text-color-white" />
            <span class="ml-2 text-color-white">WPManager</span>
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
  {
    label: 'Services',
    icon: 'pi pi-cog',
    route: '/services'
  },
  {
    label: 'Contact',
    icon: 'pi pi-envelope',
    route: '/contact'
  },
  {
    label: 'Dashboard',
    icon: 'pi pi-chart-bar',
    route: '/dashboard'
  }
]);
</script>
