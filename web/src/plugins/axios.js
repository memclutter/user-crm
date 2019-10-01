import Vue from 'vue';
import Axios from "axios";

const $axios = Axios.create({
  baseURL: "/api/",
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  }
});

Vue.prototype.$axios = $axios;

export default $axios;