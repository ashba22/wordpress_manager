import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import '@fortawesome/fontawesome-free/css/all.css'
import 'tailwindcss/tailwind.css'
import 'primevue/resources/themes/aura-dark-green/theme.css'
import 'vue-toast-notification/dist/theme-default.css'

import ToastPlugin from 'vue-toast-notification'
import PrimeVue from 'primevue/config';
import Menubar from 'primevue/menubar';
import Badge from 'primevue/badge';
import InputText from 'primevue/inputtext';
import Avatar from 'primevue/avatar'; 
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import ColumnGroup from 'primevue/columngroup'
import Button from 'primevue/button'
import Card from 'primevue/card'
import InputNumber from 'primevue/inputnumber'
import Dialog from 'primevue/dialog'
import Ripple from 'primevue/ripple'

const app = createApp(App);
app.component('Menubar', Menubar);
app.component('Badge', Badge);
app.component('InputText', InputText);
app.component('Avatar', Avatar);
app.component('DataTable', DataTable);
app.component('Column', Column);
app.component('ColumnGroup', ColumnGroup);
app.component('Button', Button);
app.component('Card', Card);
app.component('InputNumber', InputNumber);
app.component('Dialog', Dialog);




app.use(PrimeVue, {
    unstyled: false,
    inputStyle: 'filled',
    ripple: true


});
app.use(router)
app.use(ToastPlugin)
app.mount('#app')
