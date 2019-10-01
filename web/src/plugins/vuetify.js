import Vue from 'vue';
import Vuetify from 'vuetify/lib';

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
      options: {
        customProperties: true,
      },
    themes: {
      light: {
        primary: "#3f51b5",
        secondary: "#ffc107",
        accent: "#607d8b",
        error: "#ff5722",
        warning: "#ff9800",
        info: "#2196f3",
        success: "#4caf50"
      },
    },
  },
  icons: {
    iconfont: 'fa',
  },
});
