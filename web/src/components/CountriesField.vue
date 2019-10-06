<template>
  <v-autocomplete
    :outlined="outlined"
    :solo="solo"
    :placeholder="placeholder"
    :error-messages="errorMessages"
    :label="label"
    :value="value"
    :search-input.sync="search"
    item-value="code"
    item-text="name"
    :items="items"
    :loading="loading"
    :cache-items="false"
    @input="$emit('input', $event)"
  />
</template>

<script>
export default {

  props: {
    solo: Boolean,
    placeholder: String,
    outlined: Boolean,
    label: String,
    errorMessages: [Array, String],
    value: String
  },

  data: () => ({
    loading: false,
    search: '',
    items: []
  }),

  watch: {
    async search(search) {
      if (search) {
        await this.load({
          limit: 20,
          search: search
        });
      }
    }
  },

  async mounted() {
    if (this.value) {
      await this.load({
        limit: 1,
        code: this.value
      })
    }
  },

  methods: {
    async load(params) {
      this.loading = true;

      try {
        const {status, data} = await this.$axios.get('/countries', {params});

        if (status === 200) {
          this.items = data.items;
        } else {
          this.items = [];
        }
      } finally {
        this.loading = false;
      }
    }
  }

}
</script>
