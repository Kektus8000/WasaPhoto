import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import HomeView from '../views/HomeView.vue'
import FollowerView from '../views/FollowerView.vue'
import FollowedView from '../views/FollowedView.vue'
import ProfileView from '../views/ProfileView.vue'
import BannedView from '../views/BannedView.vue'
import PhotoView from '../views/PhotoView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/session', component: HomeView},
		{path: '/userProfile/:userID', component: ProfileView},
		{path: '/userProfile/:userID/following', component: FollowerView},
		{path: '/userProfile/:userID/banList', component: BannedView},
		{path: '/userProfile/:userID/followeds', component: FollowedView},
		{path: '/userProfile/:userID/publishedPhotos/:photoID', component: PhotoView}
	]
})

export default router
