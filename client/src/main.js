import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config';
import '@fortawesome/fontawesome-free/css/all.css'
import 'vue-toast-notification/dist/theme-default.css'

import 'primevue/resources/themes/soho-dark/theme.css'
import 'primevue/resources/primevue.css';
import '/node_modules/primeflex/primeflex.css';
import 'primeicons/primeicons.css';





import ToastPlugin from 'vue-toast-notification'

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
import ProgressBar from 'primevue/progressbar'
/// checkbox
import Checkbox from 'primevue/checkbox';
import RadioButton from 'primevue/radiobutton';
import InputSwitch from 'primevue/inputswitch';
import Chips from 'primevue/chips';
import Dropdown from 'primevue/dropdown';

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
app.component('ProgressBar', ProgressBar);
app.component('Checkbox', Checkbox);
app.component('RadioButton', RadioButton);
app.component('InputSwitch', InputSwitch);
app.component('Chips', Chips);
app.component('Dropdown', Dropdown);


app.use(PrimeVue, {
    unstyled: false,
    inputStyle: 'filled',
    ripple: true,

});




/// print the env variables
app.use(router)
app.use(ToastPlugin)

app.directive('ripple', Ripple);
//// import variables from the env file
app.config.globalProperties.$env = import.meta.env
/// print to the console
console.log(app.config.globalProperties.$env.VITE_API_URL);




app.mount('#app')


