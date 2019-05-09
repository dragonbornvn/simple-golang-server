import App from './vue_component/hello.js';
import CategoryView from './vue_component/categoryview.js';
import HomeView from './vue_component/homeview.js';
import TopicView from './vue_component/topic.js';

var router = new VueRouter({
	routes: [
		{ path: '/', name: 'Home', component: HomeView },
		{ path: '/category', name: 'Category', component: CategoryView },
		{ path: '/topic', name: 'Topic', component: TopicView },
		//{ path: '*', component: NotFound }
    ],
    mode: 'history',
})

Vue.use(VueRouter);
Vue.use(axios);

new Vue({
	router,
  render: h => h(App),
}).$mount(`#app`);