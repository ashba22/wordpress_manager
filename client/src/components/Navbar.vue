<template>
  <Menubar :model="items">
    <template #start>
        <img src="/logo.svg" alt="logo" class="logo" />
    </template>
    <template #item="{ item, props, hasSubmenu, root }">
        <a v-ripple class="flex align-items-center" v-bind="props.action" :href="item.route">
            <span :class="item.icon" />
            <span class="ml-2">{{ item.label }}</span>
            <Badge v-if="item.badge" :class="{ 'ml-auto': !root, 'ml-2': root }" :value="item.badge" />
            <span v-if="item.shortcut" class="ml-auto border-1 surface-border border-round surface-100 text-xs p-1">{{ item.shortcut }}</span>
            <i v-if="hasSubmenu" :class="['pi pi-angle-down', { 'pi-angle-down ml-2': root, 'pi-angle-right ml-auto': !root }]"></i>
        </a>
    </template>
    <template #end>
        <div class="flex align-items-center gap-2">
            <InputText placeholder="Search" type="text" class="w-8rem sm:w-auto" />
        </div>
    </template>
</Menubar>

</template>
<script setup>
import { ref } from 'vue';

const items = ref([
  {
    label: 'Home',
    icon: 'pi pi-home',
    route: '/'
  },
]);
</script>
