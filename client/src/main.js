import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'

// //? Importamos nuestros componentes
// import Footer from "./components/Footer.vue"

// //? Creamos nuestros componentes
// Vue.component("Footer", Footer);

Vue.config.productionTip = false

new Vue({
  vuetify,
  render: h => h(App)
}).$mount('#app')
