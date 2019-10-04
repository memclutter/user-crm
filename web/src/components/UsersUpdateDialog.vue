<template>
  <base-dialog
    v-model="dialog"
    :title="`Update user '${username}'`"
    color="primary"
    width="500"
  >
    <!-- TODO users form -->
  </base-dialog>
</template>

<script>
import { mapState } from "vuex";
import { mapFields } from "vuex-map-fields";

import BaseDialog from "@/components/BaseDialog";

export default {

  components: {
    BaseDialog
  },

  computed: {
    ...mapFields('users', {
      dialog: 'updateDialog'
    }),

    ...mapState('users', {
      username: state => state.form.fields.username
    })
  }

}
</script>
