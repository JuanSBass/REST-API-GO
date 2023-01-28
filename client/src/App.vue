<template>
  <v-card
    class="mx-auto overflow-hidden justify-center"
    width="1200"
  >
    <!-- <v-system-bar color="deep-purple darken-3"></v-system-bar> -->

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
          <v-list-item>
            <v-list-item-title>Foo</v-list-item-title>
          </v-list-item>

        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-card-text>
      The navigation drawer will appear from the bottom on smaller size screens.
    </v-card-text>
    <Products />
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
  },

  components: {
    Footer,
    Products
  },

  data: () => ({
    
    icons: ["home", "shoping", "email"],
    drawer: false
  }),

  methods: {
    getProducts: async () => {
      try {
        let response = await (await axios.get("http://localhost:3000/products")).data
      console.log(response);
      } catch (error) {
        console.log(error);
      }
      
    },
    getCategories: async () => {
      try {
        let response = await (await axios.get("http://localhost:3000/categories")).data
      console.log(response); 
      } catch (error) {
        console.log(error);
      }
    
    },
    getProductsPerCategory: async (categoId) => {
      try {
        let response = await (await axios.get(`http://localhost:3000/products/${categoId}`)).data
      console.log(response);
      } catch (error) {
        console.log(error);
      }
      
    }, 
  }
};
</script>
