<template>
  <v-dialog :value="value" @input="$emit('input', $event)" :width="width" :max-width="maxWidth">
    <v-card>
      <v-toolbar dark :color="color" flat>
        <v-toolbar-title>{{title}}</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon @click="$emit('input', false)">
          <v-icon>fa-times</v-icon>
        </v-btn>
      </v-toolbar>
      <v-card-text class="pt-8">
        <slot></slot>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script>
export default {

  props: {
    color: String,
    maxWidth: [String, Number],
    title: String,
    value: Boolean,
    width: [String, Number]
  }

}
</script>
