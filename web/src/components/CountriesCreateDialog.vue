<template>
  <v-dialog v-model="dialog">
    <v-card>
      <v-toolbar dark color="primary">
        <v-toolbar-title>Create countries</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon @click="dialog = false">
          <v-icon>fa-times</v-icon>
        </v-btn>
      </v-toolbar>
      <v-card-text>
        <!-- TODO form for create country -->
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script>
import { mapFields } from "vuex-map-fields";

export default {

  computed: {
    ...mapFields('countries', {
      dialog: 'createDialog'
    })
  }

}
</script>
