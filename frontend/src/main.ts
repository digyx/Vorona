import './plugins/axios'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import marked from 'marked'

createApp(App).use(router).mount('#app')
Object.defineProperty(App.prototype, '$marked', {value: marked})
