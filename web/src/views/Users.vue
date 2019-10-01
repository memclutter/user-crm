<template>
  <v-layout column>

    <v-flex shrink>
      <users-list/>
    </v-flex>

    <v-flex shrink align-self-center mt-8>
      <v-pagination :value="page" color="primary" :total-visible="10" :length="pageCount" @input="changePage"/>
    </v-flex>

  </v-layout>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

import UsersList from "@/components/UsersList";

export default {

  components: {
    UsersList
  },

  computed: {
    ...mapGetters('users', ['page', 'pageCount'])
  },

  async mounted() {
    await this.load();
  },

  methods: {
    ...mapActions('users', ['load', 'changePage'])
  }
}
</script>