import Vue from 'vue';
import Router from 'vue-router';
import axios from 'axios';
import Chartkick from 'chartkick';
import VueChartkick from 'vue-chartkick';
import Chart from 'chart.js';
import Vuetify from 'vuetify'


import routes from './config/routes';

Vue.use(Router);
Vue.use(VueChartkick, { Chartkick });
Vue.use(Vuetify)

const router = new Router({
  routes,
  mode: 'hash',
  saveScrollPosition: true
});

/* eslint-disable no-new */
const app = new Vue({
  router,
  //render: h => h(App)
});

app.$mount('#app')
