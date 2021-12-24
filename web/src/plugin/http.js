import axios from 'axios'
import Vue from 'vue'

const Url = 'http://localhost:7070/api/v1/'

axios.defaults.baseURL = Url

axios.interceptors.request.use(function (config) {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})
Vue.prototype.$http = axios

export { Url }
