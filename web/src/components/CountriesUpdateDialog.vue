<template>
  <base-dialog
    v-model="dialog"
    :title="`Update country '${name}'`"
    color="primary"
    width="500"
  >
    <countries-form update v-if="dialog"  />
  </base-dialog>
</template>

<script>
import { mapState } from "vuex";
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
      dialog: 'updateDialog'
    }),

    ...mapState('countries', {
      name: state => state.form.fields.name
    })
  }

}
</script>
