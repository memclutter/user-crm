<template>
  <v-dialog v-model="dialog" width="290">
    <template v-slot:activator="{on}">
      <v-text-field
        :outlined="outlined"
        :color="color"
        readonly
        :prepend-icon="prependIcon"
        :prepend-inner-icon="prependInnerIcon"
        :append-icon="appendIcon"
        :append-outer-icon="appendOuterIcon"
        :label="label"
        :error-messages="errorMessages"
        :placeholder="placeholder"
        :value="value"
        v-on="on"
      />
    </template>

    <v-date-picker
      ref="picker"
      :min="min"
      :max="max"
      :first-day-of-week="1"
      :value="value"
      @input="dialog = false; $emit('input', $event)"
    />
  </v-dialog>
</template>

<script>
export default {

  props: {
    value: [String, Date],
    label: String,
    color: String,
    outlined: Boolean,
    placeholder: String,
    prependIcon: String,
    prependInnerIcon: String,
    appendIcon: String,
    appendOuterIcon: String,
    errorMessages: [Array, String],
    max: String,
    min: String
  },

  data: () => ({
    dialog: false
  }),

  watch: {
    dialog(dialog) {
      if (dialog) {
        setTimeout(() => this.$refs.picker.activePicker = 'YEAR')
      }
    }
  }


}
</script>
