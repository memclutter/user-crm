<template>
  <base-dialog
    v-model="dialog"
    title="Create user"
    color="primary"
    width="500"
  >

    <!-- TODO users create form -->

  </base-dialog>
</template>

<script>
import { mapFields } from "vuex-map-fields";

import BaseDialog from "@/components/BaseDialog";

export default {

  components: {
    BaseDialog
  },

  computed: {
    ...mapFields('users', {
      dialog: 'createDialog'
    })
  }

}
</script>
