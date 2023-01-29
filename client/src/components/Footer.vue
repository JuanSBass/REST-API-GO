<template>
  <v-footer
    dark
    padless
  >
    <v-card
      class="flex"
      flat
      tile
    >
      <v-card-title class="teal">
        <strong class="subheading">Thank's for visit my first Vue2 app!</strong>

        <v-spacer></v-spacer>

        <v-btn
          v-for="icon in icons"
          :key="icon"
          :href="icon.ref"
          class="mx-4"
          dark
          icon
          target="_blank"
        >
          <v-icon size="24px">
            {{ icon.icon }}
          </v-icon>
        </v-btn>
      </v-card-title>

      <v-card-text class="py-2 white--text text-center">
        © Copyright{{ new Date().getFullYear() }} — <strong> Página realizada por Juan Santillan.</strong>
      </v-card-text>
    </v-card>
  </v-footer>
</template>

<script>
  export default {

    name: "Footer",

    data: () => ({
      icons: [
        {icon: 'mdi-linkedin', ref: "https://www.linkedin.com/in/juan-santillan-velazco-6775a8243/"},
        {icon: "mdi-github", ref: "https://github.com/JuanSBass"},
        {icon: 'mdi-instagram', ref: "https://www.instagram.com/manzarekluke/"},
      ],
    }),
  }
</script>