import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import HomeView from '../views/HomePage2.vue'
import FollowerView from '../views/FollowerView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/session', component: LoginView},
		{path: '/userProfile/:userID', component: ProfileView},
		{path: '/userProfile/:userID/stream/', component: HomeView},
		{path: '/userProfile/:userID/following', component: FollowerView}
	]
})

export default router
