<template>
  <v-hover>
    <template v-slot:default="{ hover }">
      <v-container >
      <v-card
        class="mx-auto"
        max-width="344"
      >

      <v-img :src="`${$props.foodImage}`" ></v-img>
      
      <v-card-text>
        <h2 class="text-h6 primary--text">
          ${{ $props.foodPrice }}
        </h2>
        {{ foodName }}
      </v-card-text>
      
      <v-fade-transition>
        <v-overlay
        v-if="hover"
        absolute
        color="#036358"
        >
        <Modal :food-image="foodImage" :food-name="foodName" />
          </v-overlay>
        </v-fade-transition>

      </v-card></v-container>
    </template>
  </v-hover>
</template>

<script>
import Modal from './Modal.vue';

export default {
  
    name: "Products",

    components: {
      Modal
    },
    props: {
    foodName: {
      type: String,
      default: ""
    },
    foodPrice: {
      type: Number,
      default: 0
    },
    foodImage: {
      type: String,
      default: "https://media.giphy.com/media/3oEjI6SIIHBdRxXI40/giphy.gif"
    }
  },
    data: () => ({
      overlay: false,
      products: null,
    }),
  }
</script>