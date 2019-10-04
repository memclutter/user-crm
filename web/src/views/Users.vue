<template>
  <v-layout column>

    <v-flex shrink>
      <users-list/>
    </v-flex>

    <v-flex shrink align-self-center mt-8>
      <v-pagination :value="page" color="primary" :total-visible="10" :length="pageCount" @input="changePage"/>
    </v-flex>

    <users-create-dialog />
    <users-update-dialog />
    <users-remove-dialog />

    <v-btn fab fixed bottom right color="success" @click="openCreateDialog">
      <v-icon>fa-plus</v-icon>
    </v-btn>
  </v-layout>
</template>

<script>
import { mapActions, mapGetters, mapMutations } from "vuex";

import UsersList from "@/components/UsersList";
import UsersCreateDialog from "@/components/UsersCreateDialog";
import UsersUpdateDialog from "@/components/UsersUpdateDialog";
import UsersRemoveDialog from "@/components/UsersRemoveDialog";

export default {

  components: {
    UsersRemoveDialog,
    UsersUpdateDialog,
    UsersCreateDialog,
    UsersList
  },

  computed: {
    ...mapGetters('users', ['page', 'pageCount'])
  },

  async mounted() {
    await this.load();
  },

  methods: {
    ...mapActions('users', ['load', 'changePage']),
    ...mapMutations('users', ['openCreateDialog'])
  }
}
</script>