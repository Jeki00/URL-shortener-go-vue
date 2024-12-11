import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import ShowsItem from './components/ShowsItem.vue'

const app = createApp(App)

app.component('shows-item', ShowsItem)


app.mount('#app')