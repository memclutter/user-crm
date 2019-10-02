<template>
  <base-dialog
    v-model="dialog"
    :title="`Remove country '${item.name}'`"
    color="error"
    width="500"
  >

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
    ...mapFields('countries', {
      dialog: 'removeDialog',
      item: 'removeItem'
    })
  }

}
</script>
