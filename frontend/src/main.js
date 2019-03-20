import Vue from "vue";
import App from "./App.vue";


Vue.config.productionTip = false;
Vue.config.devtools = true;

import Bridge from "./wailsbridge";
import EventBus from "./assets/js/eventbus";

Bridge.Start(() => {
  new Vue({
    render: h => h(App)
  }).$mount("#app");
  EventBus.$emit('wails:ready');
});
