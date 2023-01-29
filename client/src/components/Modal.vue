
<template>
  <v-row justify="center">
    <v-btn
      class="white--text"
      color="teal"
      @click="overlay = !overlay"
    >
      Show Overlay
    </v-btn>

    <v-overlay
      :z-index="zIndex"
      :value="overlay"
    >

    <div class="d-flex align-center mb-6" 
      style="display: flex; flex-direction: column; align-items: center; background-color: #414141;">
      
      <v-img
      style="display:flex; flex-direction: row-reverse; border-radius: 1rem; margin: 1rem 0"
      height="500"
      width="800"
      transition
      :src="`${foodImage}`"
      ><v-btn
        class="mx-2"
        fab
        color="indigo"
      >
      <v-icon dark>
        mdi-plus
      </v-icon>
    </v-btn></v-img>
<h1> {{ $props.foodName }} </h1>
        <v-btn
          class="white--text "
          color="teal"
          @click="overlay = false"
        >
          <v-icon>{{ icon }}</v-icon>Back
        </v-btn>

      

      </div>
    </v-overlay>
  </v-row>
</template>

<script>
  export default {
    name: "Modal",
    data: () => ({
      icon: "mdi-arrow-left-bold-box-outline",
      overlay: false,
      zIndex: 0,
    }),
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
  }
</script>
