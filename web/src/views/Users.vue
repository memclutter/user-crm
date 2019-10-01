<template>
  <div>
    <v-list three-line>
      <users-list-item v-for="item in items" :key="item.id" v-bind="item" />
    </v-list>
  </div>
</template>

<script>
import UsersListItem from "@/components/UsersListItem";

export default {

  components: {
    UsersListItem
  },

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
        const {data: {totalCount, items}} = await this.$axios.get('/users', {params});

        this.totalCount = totalCount;
        this.items = items;
      } finally {
        this.loading = false;
      }
    }
  }
}
</script>