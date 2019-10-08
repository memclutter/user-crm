import { getField, updateField} from 'vuex-map-fields'

import axios from '@/plugins/axios'

export default {
  namespaced: true,

  state: {
    loading: false,
    page: 1,
    sort: '',
    offset: 0,
    limit: 10,
    items: [],
    totalCount: 0,
    createDialog: false,
    updateDialog: false,
    removeDialog: false,
    filter: {
      search: null,
      ageRange: [10, 80],
      gender: null,
      countryCode: null
    },
    form: {
      loading: false,
      pk: null,
      fields: {
        username: null,
        email: null,
        avatarUrl: null,
        gender: null,
        birthday: null,
        countryCode: null
      },
      errors: {}
    }
  },

  getters: {
    getField,
    page: state => state.offset / state.limit + 1,
    pageCount: state => Math.floor(state.totalCount / state.limit) + (state.totalCount % state.limit > 0 ? 1 : 0),
  },

  mutations: {
    updateField,
    clearFilter: (state) => state.filter = {
      search: null,
      ageRange: [10, 80],
      gender: null,
      countryCode: null
    },
    setLoading: (state, loading) => state.loading = loading,
    setSort: (state, sort) => state.sort = sort,
    setOffset: (state, offset) => state.offset = offset,
    setData: (state, {items, totalCount}) => {
      state.items = items;
      state.totalCount = totalCount;
    },
    openCreateDialog: (state) => state.createDialog = true,
    closeCreateDialog: (state) => state.createDialog = false,
    openUpdateDialog: (state, item) => {
      state.updateDialog = true;
      state.form.pk = item.id;
      state.form.fields = item;
    },
    closeUpdateDialog: (state) => {
      state.updateDialog = false;
      state.form.pk = null;
      state.form.fields = {
        code: null,
        name: null
      };
      state.form.errors = {};
    },
    openRemoveDialog: (state, item) => {
      state.removeDialog = true;
      state.form.pk = item.id;
      state.form.fields = item;
    },
    closeRemoveDialog: (state) => {
      state.items = state.items.filter(item => item.code !== state.form.pk);
      state.removeDialog = false;
      state.form.pk = null;
      state.form.fields = {
        code: null,
        name: null
      };
    },
    setFormLoading: (state, loading) => state.form.loading = loading,
    setFormErrors: (state, errors) => state.form.errors = errors
  },

  actions: {
    async clearFilter(context) {
      context.commit('setOffset', 0);
      context.commit('clearFilter');
      await context.dispatch('load');
    },

    async applyFilter(context) {
      context.commit('setOffset', 0);
      await context.dispatch('load');
    },

    async load(context) {
      context.commit('setLoading', true);

      const params = {
        offset: context.state.offset,
        limit: context.state.limit,
        sort: context.state.sort,
        ...context.state.filter,
      };

      try {
        const {status, data} = await axios.get('/users', {params});

        if (status === 200) {
          context.commit('setData', data);
        }
      } finally {
        context.commit('setLoading', false);
      }
    },

    async changeSort(context, newSort) {
      context.commit('setSort', newSort);

      await context.dispatch('load');
    },

    async changePage(context, newPage) {
      context.commit('setOffset', (newPage - 1) * context.state.limit);

      await context.dispatch('load');
    },

    async remove(context) {
      context.commit('setFormLoading', true);

      try {
        const {status} = await axios.delete(`/users/${context.state.form.pk}`);

        if (status === 204) {
          context.commit('closeRemoveDialog');
        } else {
          // TODO handle other statuses
        }
      } finally {
        context.commit('setFormLoading', false);
      }
    },

    async submit(context, update) {
      context.commit('setFormLoading', true);

      try {
        if (update) {
          const {status, data} = await axios.put(`/users/${context.state.form.pk}`, context.state.form.fields);

          if (status === 200) {
            context.commit('closeUpdateDialog');
          } else {
            context.commit('setFormErrors', data);
          }
        } else {
          const {status, data} = await axios.post('/users', context.state.form.fields);

          if (status === 201) {
            context.commit('closeCreateDialog');
          } else {
            context.commit('setFormErrors', data);
          }
        }
      } finally {
        context.commit('setFormLoading', false);
      }
    }
  }
}