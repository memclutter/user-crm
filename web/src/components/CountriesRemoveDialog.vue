<template>
  <base-dialog
    v-model="dialog"
    :title="`Remove country '${name}'`"
    color="error"
    width="500"
  >

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
    ...mapFields('countries', {
      dialog: 'removeDialog'
    }),
    ...mapState('countries', {
      name: state => state.form.fields.name
    })
  }

}
</script>
