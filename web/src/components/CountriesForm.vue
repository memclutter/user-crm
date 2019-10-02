<template>
  <v-form @submit.prevent="submit">
    <v-text-field outlined label="Country ISO2 code" v-model="code" :error-messages="errors.code"/>
    <v-text-field outlined label="Country name" v-model="name" :error-messages="errors.name"/>

    <v-row>
      <v-col cols="auto">
        <v-btn depressed color="primary" type="submit" :loading="loading">
          Save
        </v-btn>
      </v-col>
      <v-col cols="auto">
        <v-btn depressed>
          Cancel
        </v-btn>
      </v-col>
    </v-row>
  </v-form>
</template>

<script>
import { mapFields } from "vuex-map-fields";
import { mapActions, mapState } from "vuex";

export default {

  props: {
    update: Boolean
  },

  computed: {
    ...mapFields('countries', {
      code: 'form.fields.code',
      name: 'form.fields.name'
    }),
    ...mapState('countries', {
      loading: state => state.form.loading,
      errors: state => state.form.errors
    })
  },

  methods: {
    async submit() {
      await this.$store.dispatch('countries/submit', this.update);
    }
  }

}
</script>
