import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomePage2},
		{path: '/session', component: LoginView},
		{path: '/userProfile/:userID', component: ProfileView},
		{path: '/some/:id/link', component: HomeView},
	]
})

export default router
