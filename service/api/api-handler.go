package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Call the function that handle the User Research
	rt.router.GET("/userProfile/:userID", rt.getUserProfile)

	//Call the function that handle the Set Username
	rt.router.POST("/user/:userID/username", rt.setMyUsername)

	//Call the function that handle the followUser
	rt.router.POST("/user/:userID/following", rt.followUser)

	//Call the function that handle the unFollowUser
	rt.router.DELETE("/userProfile/:userID/following/user/:followerID", rt.unFollowUser)

	//Call the function that handle the getMyStream
	rt.router.GET("/userProfile/:userID/stream", rt.getMyStream)

	//Call the function that handle the banUser
	rt.router.POST("/userProfile/:userID/banList", rt.banUser)

	//Call the function that handle the unbanUser
	rt.router.DELETE("/userProfile/:userID/banList/user/:bannedID", rt.unbanUser)

	//Call the function that handle the uploadPhoto
	rt.router.POST("/userProfile/:userID/publishedPhotos", rt.uploadPhoto)

	//Call the function that handle the deletePhoto
	rt.router.DELETE("/userProfile/:userID/publishedPhotos/photo/:photoID", rt.deletePhoto)

	//Call the function that handle the likePhoto
	rt.router.POST("/userProfile/:userID/stream/photo/:photoID/like", rt.likePhoto)

	//Call the function that handle the unlikePhoto
	rt.router.DELETE("/userProfile/:userID/stream/photo/:photoID/like", rt.unlikePhoto)

	//Call the function that handle the commentPhoto
	rt.router.POST("userProfile/:userID/stream/photo/:photoID/comments", rt.commentPhoto)

	//Call the function that handle the uncommentPhoto
	rt.router.DELETE("userProfile/:userID/stream/photo/:photoID/comments/:commentID", rt.uncommentPhoto)

	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
