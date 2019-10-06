<template>
  <v-form>
    <v-row align="center">
      <v-col cols="12" md="2">
        <v-text-field solo placeholder="Search" v-model="search"/>
      </v-col>
      <v-col cols="12" md="2">
        <countries-field solo placeholder="Country" v-model="countryCode"/>
      </v-col>
      <v-col cols="12" md="3">
        <v-range-slider label="Age" :min="10" :max="80" v-model="ageRange"/>
      </v-col>
      <v-col cols="12" md="3">
        <v-select
          solo
          placeholder="Gender"
          :items="[{value: null, text: 'All'}, {value: 'male', text: 'Male'}, {value: 'female', text: 'Female'}]"
          v-model="gender"
        />
      </v-col>
      <v-spacer></v-spacer>
      <v-col cols="auto" align-self="start">
        <v-btn icon color="primary" x-large @click="applyFilter">
          <v-icon>fa-filter</v-icon>
        </v-btn>
      </v-col>
      <v-col cols="auto" align-self="start" @click="clearFilter">
        <v-btn icon color="warning" x-large>
          <v-icon>fa-times</v-icon>
        </v-btn>
      </v-col>
    </v-row>
  </v-form>
</template>

<script>
import { mapFields } from "vuex-map-fields";
import { mapActions } from "vuex";

import CountriesField from "@/components/CountriesField";

export default {

  components: {
    CountriesField
  },

  computed: {
    ...mapFields('users', {
      search: 'filter.search',
      countryCode: 'filter.countryCode',
      ageRange: 'filter.ageRange',
      gender: 'filter.gender'
    })
  },

  methods: {
    ...mapActions('users', ['applyFilter', 'clearFilter'])
  }

}
</script>