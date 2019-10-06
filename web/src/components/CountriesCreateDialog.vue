<template>
  <base-dialog
    v-model="dialog"
    title="Create countries"
    color="primary"
    width="500"
  >
    <countries-form v-if="dialog" />
  </base-dialog>
</template>

<script>
import { mapFields } from "vuex-map-fields";

import BaseDialog from "@/components/BaseDialog";
import CountriesForm from "@/components/CountriesForm";

export default {

  components: {
    BaseDialog,
    CountriesForm
  },

  computed: {
    ...mapFields('countries', {
      dialog: 'createDialog'
    })
  }

}
</script>
