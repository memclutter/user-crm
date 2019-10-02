<template>
  <v-layout column>

    <v-flex shrink>
      <countries-list />
    </v-flex>

    <v-flex shrink align-self-center mt-8>
      <v-pagination :value="page" color="primary" :total-visible="10" :length="pageCount" @input="changePage"/>
    </v-flex>

    <countries-remove-dialog />
    <countries-update-dialog />
    <countries-create-dialog />

    <v-btn fab fixed bottom right color="success" @click="openCreateDialog">
      <v-icon>fa-plus</v-icon>
    </v-btn>
  </v-layout>
</template>

<script>
import { mapActions, mapGetters, mapMutations } from "vuex";

import CountriesList from "@/components/CountriesList";
import CountriesCreateDialog from "@/components/CountriesCreateDialog";
import CountriesUpdateDialog from "@/components/CountriesUpdateDialog";
import CountriesRemoveDialog from "@/components/CountriesRemoveDialog";

export default {

  components: {
    CountriesRemoveDialog,
    CountriesUpdateDialog,
    CountriesCreateDialog,
    CountriesList
  },

  computed: {
    ...mapGetters('countries', ['page', 'pageCount'])
  },

  async mounted() {
    await this.load();
  },

  methods: {
    ...mapActions('countries', ['load', 'changePage']),
    ...mapMutations('countries', ['openCreateDialog'])
  }
}
</script>