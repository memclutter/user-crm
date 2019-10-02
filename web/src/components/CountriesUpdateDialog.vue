<template>
  <base-dialog
    v-model="dialog"
    :title="`Update country '${item.name}'`"
    color="primary"
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
      dialog: 'updateDialog',
      item: 'updateItem'
    })
  }

}
</script>
