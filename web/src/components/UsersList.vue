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
      {value: 'id', text: '#', align: 'left', sortable: true},
      {value: 'avatarUrl', text: 'Avatar', align: 'left', sortable: false},
      {value: 'email', text: 'Email', align: 'left', sortable: false},
      {value: 'username', text: 'Username', align: 'left', sortable: false},
      {value: 'gender', text: 'Gender', align: 'left', sortable: false},
      {value: 'age', text: 'Age', align: 'left', sortable: false},
      {value: 'birthday', text: 'Birthday', align: 'left', sortable: false},
      {value: 'countryCode', text: 'Country', align: 'left', sortable: false},
      {value: 'actions', text: 'Actions', align: 'right', sortable: false}
    ]
  }),

  computed: {
    ...mapState('users', {
      loading: state => state.loading,
      items: state => state.items,
      limit: state => state.limit,
      totalCount: state => state.totalCount,
    })
  }

}
</script>
