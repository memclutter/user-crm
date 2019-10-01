<template>
  <v-layout column>

    <v-flex shrink>
      <countries-list />
    </v-flex>

    <v-flex shrink align-self-center mt-8>
      <v-pagination :value="page" color="primary" :total-visible="10" :length="pageCount" @input="changePage"/>
    </v-flex>

  </v-layout>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

import CountriesList from "@/components/CountriesList";

export default {

  components: {
    CountriesList
  },

  computed: {
    ...mapGetters('countries', ['page', 'pageCount'])
  },

  async mounted() {
    await this.load();
  },

  methods: {
    ...mapActions('countries', ['load', 'changePage'])
  }
}
</script>