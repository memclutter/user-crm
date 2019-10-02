import { getField, updateField} from 'vuex-map-fields'

import axios from '@/plugins/axios'

export default {
  namespaced: true,

  state: {
    loading: false,
    page: 1,
    offset: 0,
    limit: 10,
    items: [],
    totalCount: 0,
    createDialog: false,
    updateDialog: false,
    updateItem: {
      code: null,
      name: null
    },
    removeDialog: false,
    removeItem: {
      code: null,
      name: null
    }
  },

  getters: {
    getField,
    page: state => state.offset / state.limit + 1,
    pageCount: state => Math.floor(state.totalCount / state.limit) + (state.totalCount % state.limit > 0 ? 1 : 0),
  },

  mutations: {
    updateField,
    setLoading: (state, loading) => state.loading = loading,
    setOffset: (state, offset) => state.offset = offset,
    setData: (state, {items, totalCount}) => {
      state.items = items;
      state.totalCount = totalCount;
    },
    openCreateDialog: (state) => state.createDialog = true,
    closeCreateDialog: (state) => state.createDialog = false,
    openUpdateDialog: (state, item) => {
      state.updateDialog = true;
      state.updateItem = item;
    },
    closeUpdateDialog: (state) => {
      state.updateDialog = false;
      state.updateItem = {
        code: null,
        name: null
      };
    },
    openRemoveDialog: (state, item) => {
      state.removeDialog = true;
      state.removeItem = item;
    },
    closeRemoveDialog: (state) => {
      state.removeDialog = false;
      state.removeItem = {
        code: null,
        name: null
      };
    }
  },

  actions: {
    async load(context) {
      context.commit('setLoading', true);

      const params = {
        offset: context.state.offset,
        limit: context.state.limit
      };

      try {
        const {status, data} = await axios.get('/countries', {params});

        if (status === 200) {
          context.commit('setData', data);
        }
      } finally {
        context.commit('setLoading', false);
      }
    },

    async changePage(context, newPage) {
      context.commit('setOffset', (newPage - 1) * context.state.limit);

      await context.dispatch('load');
    },
  }
}