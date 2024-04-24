import { createRouter, createWebHistory } from 'vue-router'
import Websites from '../components/Websites.vue'
import Website from '../components/WebsitePage.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Websites },
    { path: '/website/:id', component: Website , props: true},
  ],
})

export default router