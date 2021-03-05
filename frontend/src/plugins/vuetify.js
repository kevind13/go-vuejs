import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
      options: {
        customProperties: true,
      },
    themes: {
      light: {
        primary: '#263238',
        secondary: '#455A64',
        accent: '#82B1FF',
        error: '#FF5252',
        info: '#000000',
        success: '#878282',
        warning: '#FFC107'
      },
    },
  },
  icons: {
    iconfont: 'md',
  },
});
