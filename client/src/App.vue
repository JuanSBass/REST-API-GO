<template>
  <v-card
    class="mx-auto overflow-hidden"
    width="1200"
    style="margin: 0 auto;"
  >

    <v-app-bar
      color="deep-purple accent-4"
      dark
      prominentclass="flex text-center"
      
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Food Categories</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-avatar>
      <img
        src="https://cdn.vuetifyjs.com/images/john.jpg"
        alt="John"
      >
    </v-avatar>
    </v-app-bar>

    <v-navigation-drawer
      v-model="drawer"
      absolute
      bottom
      temporary
    >
      <v-list
        nav
        dense
      >
        <v-list-item-group
          v-model="group"
          active-class="deep-purple--text text--accent-4"
        >
          <v-list-item
          
          v-for="category in categories"
          :key="category.ID"
          @click="getProductsPerCategory(category.ID, category.categoName)"
          >
            <v-list-item-title>
            
            {{ category.categoName }}</v-list-item-title>
          </v-list-item>
          <v-list-item @click="getProducts()"><v-list-item-title>All food</v-list-item-title></v-list-item>
          
          

        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-card-text>
      {{ renderCatego }}
      
    </v-card-text>
    <div style="display: flex; flex-wrap: wrap; align-items: center;">
      <v-flex v-for="product in products" :key="product.ID" >
        <Products  :food-image="product.img" :food-price="product.price" :food-name="product.name"/>
      </v-flex>
    </div>
    
    <Footer/>
  </v-card>


</template>

<script>
import Footer from "./components/Footer.vue";
import Products from "./components/Products.vue";
import axios from "axios";

export default {
  name: 'App',
  mounted(){
    this.getCategories()
    this.getProducts()
  },

  components: {
    Footer,
    Products
  },

  data: () => ({
    drawer: false,
    products: null,
    categories: null,
    renderCatego: "Enjoy the menu!",
    icons: "mdi-",
    dialog: false,
    
  }),

  methods: {
    getProducts (){
      axios.get("http://localhost:3000/products")
      .then(res => res.data)
      .then(data => this.products = data)
      this.renderCatego = "Enjoy the menu!"
    },

    getCategories(){
      axios.get("http://localhost:3000/categories")
      .then(res => res.data)
      .then(data => this.categories = data)

    },

    getProductsPerCategory (categoId, categoName){
      axios.get(`http://localhost:3000/products/${categoId}`)
      .then(res => res.data)
      .then(data => this.products = data)
      this.renderCatego = `${categoName} food.`
    }, 
  }
};

</script>
