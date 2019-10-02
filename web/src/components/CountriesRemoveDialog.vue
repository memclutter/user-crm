<template>
  <base-dialog
    v-model="dialog"
    :title="`Remove country '${name}'`"
    color="error"
    width="500"
  >
    <span class="font-weight-black">
      Are you sure remove this record?
    </span>

    <v-row>
      <v-col cols="auto">
        <v-btn depressed color="error" @click="remove">
          Yes
        </v-btn>
      </v-col>
      <v-col cols="auto">
        <v-btn depressed @click="dialog = false">
          No
        </v-btn>
      </v-col>
    </v-row>
  </base-dialog>
</template>

<script>
import { mapActions, mapState } from "vuex";
import { mapFields } from "vuex-map-fields";

import BaseDialog from "@/components/BaseDialog";

export default {

  components: {
    BaseDialog
  },

  computed: {
    ...mapFields('countries', {
      dialog: 'removeDialog'
    }),
    ...mapState('countries', {
      name: state => state.form.fields.name
    })
  },

  methods: {
    ...mapActions('countries', ['remove'])
  }

}
</script>
