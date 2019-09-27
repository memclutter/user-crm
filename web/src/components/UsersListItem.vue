<template>
  <v-list-item @click="$emit('click', $event)">
    <v-list-item-avatar color="secondary">
        <span class="white--text">
          {{username[0].toUpperCase()}}
        </span>
    </v-list-item-avatar>
    <v-list-item-content>
      <v-list-item-title>
        {{username}}
      </v-list-item-title>
      <v-list-item-title>
        {{email}}
      </v-list-item-title>
    </v-list-item-content>
    <v-spacer></v-spacer>
    <v-list-item-action>
      <v-list-item-icon>
        <flag :iso="countryCode" :squared="false"/>
      </v-list-item-icon>
    </v-list-item-action>
    <v-list-item-action>
      <v-list-item-icon>
        <v-icon :color="gender === 'male' ? 'secondary' : 'primary'">
          {{gender === 'male' ? 'fa-mars' : 'fa-venus'}}
        </v-icon>
      </v-list-item-icon>
    </v-list-item-action>
  </v-list-item>
</template>

<script>
export default {

  props: {
    id: Number,
    username: String,
    email: String,
    countryCode: String,
    gender: String
  }

}
</script>
