<template>
  <base-dialog
    v-model="dialog"
    title="Create user"
    color="primary"
    width="500"
  >
    <users-form />
  </base-dialog>
</template>

<script>
import { mapFields } from "vuex-map-fields";

import BaseDialog from "@/components/BaseDialog";
import UsersForm from "@/components/UsersForm";

export default {

  components: {
    UsersForm,
    BaseDialog
  },

  computed: {
    ...mapFields('users', {
      dialog: 'createDialog'
    })
  }

}
</script>
