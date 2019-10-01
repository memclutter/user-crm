<template>
  <v-data-table
    class="elevation-5"
    :headers="headers"
    :loading="loading"
    :items="items"
    :items-per-page="limit"
    :server-items-length="totalCount"
    hide-default-footer
    disable-sort
  >
    <template v-slot:item.code="{ item }">
      <flag :iso="item.code" :squared="false"/>
    </template>
    <template v-slot:item.actions="{ item }">
      <v-btn icon color="success">
        <v-icon>fa-pen</v-icon>
      </v-btn>
      <v-btn icon color="error">
        <v-icon>fa-trash</v-icon>
      </v-btn>
    </template>
  </v-data-table>
</template>

<script>
import { mapState } from "vuex";

export default {

  data: () => ({
    headers: [
      {value: 'code', text: 'Code', align: 'left', sortable: true},
      {value: 'name', text: 'Name', align: 'left', sortable: false},
      {value: 'actions', text: 'Actions', align: 'right', sortable: false}
    ]
  }),

  computed: {
    ...mapState('countries', {
      loading: state => state.loading,
      items: state => state.items,
      limit: state => state.limit,
      totalCount: state => state.totalCount,
    })
  }

}
</script>
