<template>
  <v-data-table
    class="elevation-5"
    :headers="headers"
    :loading="loading"
    :items="items"
    :items-per-page="limit"
    :server-items-length="totalCount"
    hide-default-footer
    hide-default-header
    disable-sort
  >
    <template v-slot:header="{props: {headers}}">
      <table-sort-header :sort="sort" :headers="headers" @updateSort="changeSort" />
    </template>
    <template v-slot:item.avatarUrl="{ item }">
      <v-avatar color="secondary" :size="32">
        <span class="white--text">
          {{item.username[0].toUpperCase()}}
        </span>
      </v-avatar>
    </template>
    <template v-slot:item.gender="{ item }">
      <v-icon :color="item.gender === 'male' ? 'secondary' : 'primary'">
        {{item.gender === 'male' ? 'fa-mars' : 'fa-venus'}}
      </v-icon>
    </template>
    <template v-slot:item.countryCode="{ item }">
      <flag :iso="item.countryCode" :squared="false"/>
    </template>
    <template v-slot:item.birthday="{ item }">
      {{ new Date() | moment('diff', item.birthday, 'years') }} (birthday {{ item.birthday | moment('DD/MM/YYYY') }})
    </template>
    <template v-slot:item.actions="{ item }">
      <v-btn icon color="success" @click="openUpdateDialog(item)">
        <v-icon>fa-pen</v-icon>
      </v-btn>
      <v-btn icon color="error" @click="openRemoveDialog(item)">
        <v-icon>fa-trash</v-icon>
      </v-btn>
    </template>
  </v-data-table>
</template>

<script>
import { mapActions, mapMutations, mapState } from "vuex";

import TableSortHeader from "@/components/TableSortHeader";

export default {

  components: {
    TableSortHeader
  },

  data: () => ({
    headers: [
      {value: 'id', text: '#', align: 'left', sortable: true},
      {value: 'avatarUrl', text: 'Avatar', align: 'left', sortable: false},
      {value: 'email', text: 'Email', align: 'left', sortable: false},
      {value: 'username', text: 'Username', align: 'left', sortable: false},
      {value: 'gender', text: 'Gender', align: 'left', sortable: true},
      {value: 'birthday', text: 'Age/Birthday', align: 'left', sortable: true},
      {value: 'countryCode', text: 'Country', align: 'left', sortable: true},
      {value: 'actions', text: 'Actions', align: 'right', sortable: false}
    ]
  }),

  computed: {
    ...mapState('users', {
      loading: state => state.loading,
      items: state => state.items,
      sort: state => state.sort,
      limit: state => state.limit,
      totalCount: state => state.totalCount,
    })
  },

  methods: {
    ...mapActions('users', ['changeSort']),
    ...mapMutations('users', ['openUpdateDialog', 'openRemoveDialog'])
  }

}
</script>
