<template>
  <thead class="table-header v-data-table-header">
  <tr>
    <template v-for="(header,index) in headers">
      <template v-if="header.sortable">
        <th :key="index" :class="renderClasses(header)" @click="updateSort(header)">
          <span>{{ header.text }}</span>
          <v-icon right v-if="header.sortable" size="18" class="v-data-table-header__icon">fa-long-arrow-alt-down</v-icon>
        </th>
      </template>
      <template v-else>
        <th :key="index" :class="renderClasses(header)">
          <span>{{ header.text }}</span>
        </th>
      </template>
    </template>
  </tr>
  </thead>
</template>

<script>
import { sortArrToStr, sortStrToArr } from "@/helpers";

export default {
  props: {
    headers: Array,
    sort: String
  },

  data: (vm) => ({
    sortArr: sortStrToArr(vm.sort)
  }),

  methods: {

    updateSort(header) {
      const index = this.sortArr.findIndex(el => el.name === header.value)

      if (index === -1) {
        this.sortArr.push({name: header.value, desc: false})
      } else {
        const desc = this.sortArr[index].desc;
        if (desc) {
          this.sortArr.splice(index, 1)
        } else {
          this.sortArr[index].desc = !desc
        }
      }

      this.$emit('updateSort', sortArrToStr(this.sortArr));
    },

    renderClasses(header) {
      let classes = {
        'sortable': header.sortable,
        [`text-${header.align}`]: true,
        'column': true
      };

      const sort = this.sortArr.find(el => el.name === header.value);

      if (sort) {
        classes = {
          ...classes,
          active: true,
          asc: !sort.desc,
          desc: sort.desc
        }
      }

      return classes
    }
  }
}
</script>
