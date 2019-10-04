<template>
  <base-dialog
    v-model="dialog"
    :title="`Remove user '${username}'`"
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
    ...mapFields('users', {
      dialog: 'removeDialog'
    }),
    ...mapState('users', {
      username: state => state.form.fields.username
    })
  },

  methods: {
    ...mapActions('users', ['remove'])
  }

}
</script>
