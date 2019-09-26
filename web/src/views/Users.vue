<template>
  <v-list three-line>
    <v-list-item v-for="item in items" :key="item.id" @click="">
      <v-list-item-avatar color="secondary">
        <span class="white--text">
          {{item.username[0].toUpperCase()}}
        </span>
      </v-list-item-avatar>
      <v-list-item-content>
        <v-list-item-title>
          {{item.username}}
        </v-list-item-title>
        <v-list-item-title>
          {{item.email}}
        </v-list-item-title>
      </v-list-item-content>
      <v-spacer></v-spacer>
      <v-list-item-action>
        <v-list-item-icon>
          <flag :iso="item.countryCode" :squared="false"/>
        </v-list-item-icon>
      </v-list-item-action>
      <v-list-item-action>
        <v-list-item-icon>
          <v-icon :color="item.gender === 'male' ? 'secondary' : 'primary'">
            {{item.gender === 'male' ? 'fa-mars' : 'fa-venus'}}
          </v-icon>
        </v-list-item-icon>
      </v-list-item-action>
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
        const params = {limit: this.limit, offset: this.offset}
        const {data: {totalCount, items}} = await $http.get('/users', {params});

        this.totalCount = totalCount;
        this.items = items;
      } finally {
        this.loading = false;
      }
    }
  }
}
</script>