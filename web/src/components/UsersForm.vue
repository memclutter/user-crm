<template>
  <v-form @submit.prevent="submit">
    <v-text-field outlined label="Username" v-model="username" :error-messages="errors.username"/>
    <v-text-field outlined label="Email" v-model="email" :error-messages="errors.email"/>
    <v-radio-group row label="Gender" v-model="gender" :error-messages="errors.gender" >
      <v-radio value="male" label="Male" />
      <v-radio value="female" label="Female" />
    </v-radio-group>
    <birthday-field outlined label="Brithday" v-model="birthday" :error-messages="errors.birthday" />

    <countries-field outlined label="Country" v-model="countryCode" :error-messages="errors.countryCode" />

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
import { mapState } from "vuex";
import BirthdayField from "@/components/BirthdayField";
import CountriesField from "@/components/CountriesField";

export default {
  components: {CountriesField, BirthdayField},
  props: {
    update: Boolean
  },

  computed: {
    ...mapFields('users', {
      username: 'form.fields.username',
      email: 'form.fields.email',
      gender: 'form.fields.gender',
      birthday: 'form.fields.birthday',
      countryCode: 'form.fields.countryCode',
    }),
    ...mapState('users', {
      loading: state => state.form.loading,
      errors: state => state.form.errors
    })
  },

  methods: {
    async submit() {
      await this.$store.dispatch('users/submit', this.update);
    }
  }

}
</script>
