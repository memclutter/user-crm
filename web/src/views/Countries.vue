<template>
  <v-list three-line>
    <v-list-item v-for="item in items" :key="item.id" @click="">
      <v-list-item-action>
        <v-list-item-icon >
          <flag :iso="item.code" :squared="false"/>
        </v-list-item-icon>
      </v-list-item-action>
      <v-list-item-content>
        <v-list-item-title>
          {{item.name}}
        </v-list-item-title>
      </v-list-item-content>
    </v-list-item>
  </v-list>
</template>

<script>
import $http from '@/http'

export default {

  data: () => ({
    loading: false,
    offset: 0,
    limit: 25,
    items: [],
    totalCount: 0
  }),

  async mounted() {
    await this.load();
  },

  methods: {
    async load() {
      this.loading = true;

      try {
        const params = {limit: this.limit, offset: this.offset};
        const {data: {totalCount, items}} = await $http.get('/countries', {params});

        this.totalCount = totalCount;
        this.items = items;
      } finally {
        this.loading = false;
      }
    }
  }
}
</script>