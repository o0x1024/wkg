import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Antd from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
// import VueParticles from 'vue-particles';

// const compression = require('compression')

const app = createApp(App);
// app.use((VueParticles))
// app.use(compression())
app.use(store);
app.use(router);
app.use(Antd);
app.mount("#app");
