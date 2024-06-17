import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import HomeView from '../views/HomePage2.vue'
import FollowerView from '../views/FollowerView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/session', component: LoginView},
		{path: '/userProfile/:userID', component: HomeView},
		{path: '/userProfile/:userID/following', component: FollowerView}
	]
})

export default router
