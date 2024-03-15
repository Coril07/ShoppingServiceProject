import { createRouter, createWebHistory } from "vue-router";
import Index from "./views/Index.vue";
import login from "./views/login.vue";
import sign_up from "./views/sign_up.vue";
import log_out from "./views/log_out.vue";
import admin from "./views/admin.vue";
import product from "./views/product.vue";
import cart from "./views/cart.vue";
import order from "./views/order.vue";
import personalfile from "./views/personalfile.vue";
const routes = [
  {
    path: "/",
    component: Index,
  },
  {
    path: "/login",
    component: login,
  },
  {
    path: "/sign_up",
    component: sign_up,
  },
  {
    path: "/log_out",
    component: log_out,
  },
  {
    path: "/admin",
    component: admin,
  },
  {
    path: "/product",
    component: product,
  },
  {
    path: "/cart",
    component: cart,
  },
  {
    path: "/order",
    component: order,
  },
  {
    path: "/personalfile",
    component: personalfile,
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
