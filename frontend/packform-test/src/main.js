import Vue from 'vue'
import VueRouter from 'vue-router'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import DateRangePicker from "@gravitano/vue-date-range-picker";

import App from './App.vue'
import Orders from './Orders.vue'
import Home from './Home.vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.config.productionTip = false

Vue.use(VueRouter);
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
Vue.use(DateRangePicker);

const routes = [
  { path: '/orders', component: Orders },
  { path: '/', component: Home },
];

const router = new VueRouter({
  mode: 'history',
  routes // short for `routes: routes`
});

new Vue({
  router,
  render: h => h(App),
}).$mount('#app');
